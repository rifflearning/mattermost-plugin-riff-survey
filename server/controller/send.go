package controller

import (
	"net/http"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
)

var sendSurvey = &Endpoint{
	Path:         "/send",
	Method:       http.MethodGet,
	Execute:      executeSendSurvey,
	RequiresAuth: true,
}

func executeSendSurvey(w http.ResponseWriter, r *http.Request) error {
	userID := r.URL.Query().Get("user_id")
	meetingID := r.URL.Query().Get("meeting_id")

	// TODO: verify that user is in the meeting

	config.Mattermost.LogDebug("Send survey executed.", "userID", userID, "meetingID", meetingID)

	if err := platform.SendSurveyPost(userID, meetingID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
