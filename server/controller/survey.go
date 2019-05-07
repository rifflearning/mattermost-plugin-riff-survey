package controller

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
)

var sendSurvey = &Endpoint{
	Path:         "/survey",
	Method:       http.MethodGet,
	Execute:      executeSendSurvey,
	RequiresAuth: true,
}

func executeSendSurvey(w http.ResponseWriter, r *http.Request) error {
	conf := config.GetConfig()
	// userID := r.Header.Get(config.HeaderMattermostUserID)
	channelID := r.URL.Query().Get("channel_id")

	post := &model.Post{
		UserId:    conf.BotUserID,
		ChannelId: channelID,
		Type:      model.POST_DEFAULT,
		Message:   "Sample survey post",
		Props: model.StringInterface{
			"from_webhook": "true",
			// "override_username": config.OverrideUsername,
			// "override_icon_url": config.OverrideIconURL,
		},
	}

	if _, err := config.Mattermost.CreatePost(post); err != nil {
		http.Error(w, "Error creating the survey post.", http.StatusInternalServerError)
		return errors.Wrap(err, "Error creating survey post for channel: "+channelID)
	}

	return nil
}
