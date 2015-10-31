package dns-infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform'
)

//Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:	schema.Typestring,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USER", nil), 
				Description: "username for infoblox",
			},
			"password": &schema.Schema{
				Type:	schema.Typestring,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil), 
				Description: "password for infoblox",
			},
			"url": &schema.Schema{
				Type:	schema.Typestring,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_URL", nil), 
				Description: "url for infoblox api",
			},
		},
		
		ResourcesMap: map[string]*schema.Resource{
			"dnsinfoblox_record": resourceDNSInfobloxRecord(),
		},

		ConfigureFunc: provideConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User: d.Get("user").(string),
		Password: d.Get("password").(string),
		Url: d.get("url").(string),
	}

	return config.Client()
}
