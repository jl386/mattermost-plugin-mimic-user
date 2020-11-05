package main

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v5/plugin"

	"github.com/jl386/mattermost-plugin-mimic-user/server/config"
	"github.com/jl386/mattermost-plugin-mimic-user/server/controller"
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
	// If OnActivate has not been run yet.
	if config.Mattermost == nil {
		return nil
	}
	var configuration config.Configuration

	if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
		config.Mattermost.LogError("Error in LoadPluginConfiguration.", "Error", err.Error())
		return err
	}

	if err := configuration.ProcessConfiguration(); err != nil {
		config.Mattermost.LogError("Error in ProcessConfiguration.", "Error", err.Error())
		return err
	}

	if err := configuration.IsValid(); err != nil {
		config.Mattermost.LogError("Error in Validating Configuration.", "Error", err.Error())
		return err
	}

	config.SetConfig(&configuration)
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("New request:", "Host", r.Host, "RequestURI", r.RequestURI, "Method", r.Method)

	conf := config.GetConfig()
	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured.", "Error", err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	controller.InitAPI().ServeHTTP(w, r)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
