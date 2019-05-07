package config

import (
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
)

const (
	CommandPrefix             = PluginName
	URLMappingKeyPrefix       = "url_"
	ServerExeToWebappRootPath = "/../webapp"

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
	BotUsername string `json:"BotUsername"`

	// Derived Attributes
	BotUserID string
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

	return nil
}

func (c *Configuration) IsValid() error {
	// Add config validations here.
	// Check for required fields, formats, etc.

	return nil
}
