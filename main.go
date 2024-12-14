package main

import (
	"github.com/icpes/steampipe-plugin-azureaadright/azureaadright"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azureaadright.Plugin})
}
