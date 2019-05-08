package controller

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
)

var sendSurvey = &Endpoint{
	Path:         "/send",
	Method:       http.MethodGet,
	Execute:      executeSendSurvey,
	RequiresAuth: true,
}

func executeSendSurvey(w http.ResponseWriter, r *http.Request) error {
	conf := config.GetConfig()
	userID := r.Header.Get(config.HeaderMattermostUserID)

	channel, appErr := config.Mattermost.GetDirectChannel(conf.BotUserID, userID)
	if appErr != nil {
		http.Error(w, "Error creating DM Channel.", http.StatusInternalServerError)
		return errors.Wrap(appErr, "Error creating DM Channel.")
	}

	post := &model.Post{
		UserId:    conf.BotUserID,
		ChannelId: channel.Id,
		Type:      model.POST_DEFAULT,
		Message:   "Sample survey post",
		Props: model.StringInterface{
			"from_webhook":      "true",
			"override_username": config.OverrideUsername,
		},
	}

	if _, err := config.Mattermost.CreatePost(post); err != nil {
		http.Error(w, "Error creating the survey post.", http.StatusInternalServerError)
		return errors.Wrap(err, "Error creating survey post for channel: "+channel.Id)
	}

	return nil
}
