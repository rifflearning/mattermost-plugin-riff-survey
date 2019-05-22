package controller

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/platform"
)

var sendSurvey = &Endpoint{
	Path:         "/send",
	Method:       http.MethodGet,
	Execute:      executeSendSurvey,
	RequiresAuth: true,
}

func executeSendSurvey(w http.ResponseWriter, r *http.Request) error {
	conf := config.GetConfig()
	userID := r.URL.Query().Get("user_id")
	meetingID := r.URL.Query().Get("meeting_id")

	// TODO: verify that user is in the meeting

	config.Mattermost.LogDebug("Send survey executed.", "userID", userID, "meetingID", meetingID)

	channel, appErr := config.Mattermost.GetDirectChannel(conf.BotUserID, userID)
	if appErr != nil {
		http.Error(w, "Unable to create DM Channel.", http.StatusInternalServerError)
		return errors.Wrap(appErr, "Unable to create DM Channel.")
	}

	// TODO: Get add meetingID to props instead of surveyID and get survey questions using meetingID
	surveyID := config.HardcodedSurveyID
	latestSurveyInfo := platform.GetLatestSurveyInfo(surveyID)
	if latestSurveyInfo == nil {
		http.Error(w, "Survey does not exist.", http.StatusInternalServerError)
		return errors.New("survey does not exist")
	}

	// TODO: Refactor this to platform.CreateSurveyPost
	post := &model.Post{
		UserId:    conf.BotUserID,
		ChannelId: channel.Id,
		Type:      "custom_survey",
		Message:   "Survey",
		Props: model.StringInterface{
			"from_webhook":      "true",
			"override_username": config.OverrideUsername,
			"meeting_id":        meetingID,
			"survey_id":         surveyID,
			"survey_version":    latestSurveyInfo.Version,
		},
	}

	if _, err := config.Mattermost.CreatePost(post); err != nil {
		return errors.Wrap(err, "failed to create survey post for the channel: "+channel.Id)
	}

	return nil
}
