package main

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/controller"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost == nil {
		config.Mattermost = p.API
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
		config.Mattermost.LogError("Processing " + r.URL.String() + ". Error: " + err.Error())
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
