package config

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
	"go.uber.org/atomic"
)

const (
	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"
	PluginAPIBaseURL       = URLPluginBase + "/api/v1"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
}

func GetConfig() *Configuration {
	conf := config.Load()
	if conf == nil {
		return nil
	}
	return conf.(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

// ProcessConfiguration is used for post-processing on the configuration.
func (c *Configuration) ProcessConfiguration() error {
	return nil
}

// IsValid is used for config validations.
func (c *Configuration) IsValid() error {
	return nil
}
