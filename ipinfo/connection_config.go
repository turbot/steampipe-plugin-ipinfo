package ipinfo

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type ipinfoConfig struct {
	Token        *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
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
