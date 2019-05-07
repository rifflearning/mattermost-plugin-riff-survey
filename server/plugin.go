package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"

	"github.com/Brightscout/mattermost-plugin-survey/server/command"
	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/controller"
	"github.com/Brightscout/mattermost-plugin-survey/server/util"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	if err := p.registerCommands(); err != nil {
		config.Mattermost.LogError(err.Error())
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

func (p *Plugin) registerCommands() error {
	for _, c := range command.Commands {
		if err := config.Mattermost.RegisterCommand(c.Command); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split, argErr := util.SplitArgs(args.Command)
	if argErr != nil {
		return util.CommandError(argErr.Error())
	}

	cmdName := split[0]
	var params []string

	if len(split) > 1 {
		params = split[1:]
	}

	commandConfig := command.Commands[cmdName]
	if commandConfig == nil {
		return nil, &model.AppError{Message: "Unknown command: [" + cmdName + "] encountered"}
	}

	context := p.prepareContext(args)
	if response, err := commandConfig.Validate(params, context); response != nil {
		return response, err
	}

	config.Mattermost.LogInfo("Executing command: " + cmdName + " with params: [" + strings.Join(params, ", ") + "]")
	return commandConfig.Execute(params, context)
}

func (p *Plugin) prepareContext(args *model.CommandArgs) command.Context {
	return command.Context{
		CommandArgs: args,
		Props:       make(map[string]interface{}),
	}
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
		config.Mattermost.LogError(fmt.Sprintf("This endpoint: %s %s requires Authentication.", endpoint.Method, endpoint.Path))
		http.Error(w, "This endpoint requires Authentication.", http.StatusForbidden)
		return
	}

	if err := endpoint.Execute(w, r); err != nil {
		config.Mattermost.LogError("Processing " + r.URL.String() + ". Error: " + err.Error())
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
