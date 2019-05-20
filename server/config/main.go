package config

import (
	"encoding/json"
	"fmt"

	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/Brightscout/mattermost-plugin-survey/server/model"
)

const (
	CommandPrefix             = PluginName
	URLMappingKeyPrefix       = "url_"
	ServerExeToWebappRootPath = "/../webapp"

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"

	OverrideUsername = "Riff Bot"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
	BotUsername string `json:"BotUsername"`
	Survey string `json:"Survey"`

	// Derived Attributes
	BotUserID string
	ParsedSurvey model.Survey
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	// any post-processing on configurations goes here

	user, err := Mattermost.GetUserByUsername(c.BotUsername)
	if err != nil {
		return errors.Wrap(err, "failed to get bot user")
	}
	c.BotUserID = user.Id

	var parsedSurvey model.Survey
	if err := json.Unmarshal([]byte(c.Survey), &parsedSurvey); err != nil {
		Mattermost.LogError("Unable to parse json for the survey. Error: " + err.Error() + ". Please make sure it is a valid JSON of the provided format.")
		return err
	}
	c.ParsedSurvey = parsedSurvey

	fmt.Println(parsedSurvey)

	return nil
}

func (c *Configuration) IsValid() error {
	if c.BotUsername == "" {
		return errors.New("Bot username cannot be empty")
	}

	return nil
}
