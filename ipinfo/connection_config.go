package ipinfo

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type ipinfoConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &ipinfoConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ipinfoConfig {
	if connection == nil || connection.GetConfig() == nil {
		return ipinfoConfig{}
	}
	config, _ := connection.GetConfig().(ipinfoConfig)
	return config
}
