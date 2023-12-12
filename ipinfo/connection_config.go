package ipinfo

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type ipinfoConfig struct {
	Token        *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &ipinfoConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ipinfoConfig {
	if connection == nil || connection.Config == nil {
		return ipinfoConfig{}
	}
	config, _ := connection.Config.(ipinfoConfig)
	return config
}
