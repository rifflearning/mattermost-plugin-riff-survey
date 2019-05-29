package reminders

import (
	"time"

	serverModel "github.com/mattermost/mattermost-server/model"
	serverUtils "github.com/mattermost/mattermost-server/utils"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
	"github.com/Brightscout/mattermost-plugin-survey/server/util"
)

var (
	addReminderChannel chan string
	doneChannel        chan bool

	ticker *time.Ticker
)

func InitReminders() {
	addReminderChannel = make(chan string)
	doneChannel = make(chan bool)
	ticker = time.NewTicker(config.ReminderTickerDuration)

	go HandleReminders()
}

func StopReminders() {
	doneChannel <- true
	ticker.Stop()
}

func HandleReminders() {
	for {
		select {
		case t := <-ticker.C:
			sendReminderNotifications(t)

		case postID := <-addReminderChannel:
			addReminder(postID)

		case <-doneChannel:
			return
		}
	}
}

func sendReminderNotifications(currentTime time.Time) {
	conf := config.GetConfig()
	remindersList, err := config.Store.GetRemindersList()
	if err != nil {
		config.Mattermost.LogError("Failed to get reminders list. Unable to send survey reminders.", "Error", err.Error())
		return
	}

	newRemindersList := make([]string, 0)
	reminderPost := &serverModel.Post{
		UserId:  conf.BotUserID,
		Message: conf.ReminderText,
		Props: serverModel.StringInterface{
			"from_webhook":      "true",
			"override_username": config.OverrideUsername,
		},
	}

	for _, surveyPostID := range remindersList {
		reminderMetadata, err := config.Store.GetReminderMetadata(surveyPostID)
		if err != nil {
			config.Mattermost.LogError("Failed to get reminder metadata. Unable to send the reminder.", "Error", err.Error(), "PostID", surveyPostID)
			continue
		}

		if reminderMetadata.TotalRemindersSent > conf.MaxReminderCountInt {
			config.Mattermost.LogInfo("Max reminders sent for post.", "SurveyPostID", surveyPostID, "ReminderMetadata", string(reminderMetadata.EncodeToByte()))
			continue
		}

		userMeetingMetadata := util.GetUserMeetingMetadata(reminderMetadata.UserID, reminderMetadata.MeetingID)
		userHasResponded := userMeetingMetadata != nil && userMeetingMetadata.RespondedAt != 0
		if userHasResponded {
			config.Mattermost.LogInfo("User has already responded to the survey. Reminder not sent.", "PostID", surveyPostID, "ReminderMetadata", string(reminderMetadata.EncodeToByte()))
			continue
		}

		shouldSendFirstReminder := reminderMetadata.TotalRemindersSent == 0 && currentTime.After(serverUtils.TimeFromMillis(reminderMetadata.SurveySentAt).Add(conf.ReminderIntervalDuration))
		shouldSendSubsequentReminder := currentTime.After(serverUtils.TimeFromMillis(reminderMetadata.PreviousReminderSentAt).Add(conf.ReminderIntervalDuration))
		shouldSendReminder := shouldSendFirstReminder || shouldSendSubsequentReminder

		if shouldSendReminder {
			reminderPost.ChannelId = reminderMetadata.ChannelID
			reminderPost.ParentId = surveyPostID
			reminderPost.RootId = surveyPostID

			post, err := config.Mattermost.CreatePost(reminderPost)
			if err != nil {
				config.Mattermost.LogError("Failed to create reminder post.", "SurveyPostID", surveyPostID, "Error", err.Error())
				continue
			}

			reminderMetadata.TotalRemindersSent = reminderMetadata.TotalRemindersSent + 1
			reminderMetadata.PreviousReminderSentAt = post.CreateAt
			if err := config.Store.SaveReminderMetadata(reminderMetadata); err != nil {
				config.Mattermost.LogError("Failed to save reminder metadata.", "Error", err.Error())
				continue
			}
		}

		if reminderMetadata.TotalRemindersSent < conf.MaxReminderCountInt {
			newRemindersList = append(newRemindersList, surveyPostID)
		}
	}

	if err := config.Store.SaveRemindersList(newRemindersList); err != nil {
		config.Mattermost.LogError("Failed to update reminders list.", "OldRemindersList", string(model.StringArrayToByte(remindersList)), "NewRemindersList", string(model.StringArrayToByte(newRemindersList)), "Error", err.Error())
		return
	}
}

func addReminder(postID string) {
	remindersList, err := config.Store.GetRemindersList()
	if err != nil {
		config.Mattermost.LogError("Failed to get reminders list. Unable to add reminder to reminders list.", "Error", err.Error(), "ReminderPostID", postID)
		return
	}

	remindersList = append(remindersList, postID)
	if err := config.Store.SaveRemindersList(remindersList); err != nil {
		config.Mattermost.LogError("Failed to save reminders list. Unable to add reminder to reminders list.", "Error", err.Error(), "ReminderPostID", postID)
		return
	}
}

func AddNew(postID, channelID, userID, meetingID string, surveySentAt int64) {
	reminderMetadata := &model.ReminderMetadata{
		MeetingID:              meetingID,
		UserID:                 userID,
		PostID:                 postID,
		ChannelID:              channelID,
		SurveySentAt:           surveySentAt,
		PreviousReminderSentAt: 0,
		TotalRemindersSent:     0,
	}

	reminderMetadata.PreSave()
	if err := config.Store.SaveReminderMetadata(reminderMetadata); err != nil {
		config.Mattermost.LogError("Failed to save reminder metadata.", "Error", err.Error())
		return
	}

	addReminderChannel <- postID
}
