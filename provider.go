package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"ess_cluster": resourceServer(),
		},
	}
}
