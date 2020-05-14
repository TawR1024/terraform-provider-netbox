package main

import (
	"github.com/TawR1024/terrafom-provider-netbox/plugin/providers/netbox"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netbox.Provider,
	})
}
