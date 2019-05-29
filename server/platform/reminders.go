package platform

import (
	"time"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	serverModel "github.com/mattermost/mattermost-server/model"
)

// SendSurveyReminders sends reminder posts to the user to fill the survey.
func SendSurveyReminders(postID, channelID, userID, meetingID string) {
	conf := config.GetConfig()
	reminderPost := &serverModel.Post{
		UserId:    conf.BotUserID,
		ChannelId: channelID,
		Message:   conf.ReminderText,
		ParentId:  postID,
		RootId:    postID,
		Props: serverModel.StringInterface{
			"from_webhook":      "true",
			"override_username": config.OverrideUsername,
		},
	}

	for i := 0; i < conf.MaxReminderCountInt; i++ {
		time.Sleep(conf.ReminderIntervalDuration)

		userMeetingMetadata := GetUserMeetingMetadata(userID, meetingID)
		userHasResponded := userMeetingMetadata != nil && userMeetingMetadata.RespondedAt != 0
		if userHasResponded {
			return
		}

		if _, err := config.Mattermost.CreatePost(reminderPost); err != nil {
			config.Mattermost.LogError("Failed to create reminder post.", "PostID", postID, "ChannelID", channelID, "Error", err.Error())
		}
	}
}
