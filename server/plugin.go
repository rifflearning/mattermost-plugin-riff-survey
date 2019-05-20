package main

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/controller"
	"github.com/Brightscout/mattermost-plugin-survey/server/platform"
	"github.com/Brightscout/mattermost-plugin-survey/server/store/kvstore"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API
	config.Store = kvstore.NewStore()

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	// If the plugin is not activated
	if config.Mattermost == nil {
		return nil
	}

	var configuration config.Configuration

	if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
		config.Mattermost.LogError("Error in LoadPluginConfiguration: " + err.Error())
		return err
	}

	if err := configuration.ProcessConfiguration(); err != nil {
		config.Mattermost.LogError("Error in ProcessConfiguration: " + err.Error())
		return err
	}

	if err := configuration.IsValid(); err != nil {
		config.Mattermost.LogError("Error in Validating Configuration: " + err.Error())
		return err
	}

	config.SetConfig(&configuration)

	if err := initSurvey(); err != nil {
		config.Mattermost.LogError("Error in Initialising Survey: " + err.Error())
		return err
	}

	return nil
}

func initSurvey() error {
	survey := config.GetConfig().ParsedSurvey

	// TODO: for v2, determine ID for survey
	survey.ID = config.HardcodedSurveyID

	if err := platform.SaveSurvey(survey); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if err := conf.IsValid(); err != nil {
		config.Mattermost.LogError("This plugin is not configured: " + err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	endpoint := controller.GetEndpoint(r)
	if endpoint == nil {
		return
	}

	if endpoint.RequiresAuth && !controller.Authenticated(w, r) {
		config.Mattermost.LogError(fmt.Sprintf("Endpoint: %s '%s' requires Authentication.", endpoint.Method, endpoint.Path))
		http.Error(w, "This endpoint requires authentication.", http.StatusForbidden)
		return
	}

	if err := endpoint.Execute(w, r); err != nil {
		config.Mattermost.LogError(fmt.Sprintf("Processing: %s '%s'.", r.Method, r.URL.String()), "Error", err.Error())
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
