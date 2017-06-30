package main

import (
	"github.com/bhoriuchi/terraform-provider-javascript/javascript"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: javascript.Provider,
	})
}
