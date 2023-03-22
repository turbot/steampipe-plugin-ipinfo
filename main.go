package main

import (
	"github.com/turbot/steampipe-plugin-ipinfo/ipinfo"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: ipinfo.Plugin})
}
