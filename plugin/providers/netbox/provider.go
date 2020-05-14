package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var descriptions map[string]string

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"virtual_machine": resourceVM(),
		},
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["netbox api token"],
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["netbox main url"],
			},
		},
		ConfigureFunc: configureProvider,
	}
}
func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:    d.Get("token").(string),
		Url: d.Get("url").(string),
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}