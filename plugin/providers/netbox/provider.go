package netbox

import (
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	netboxClient "github.com/netbox-community/go-netbox/netbox/client"
)

var descriptions map[string]string

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_virtual_machine": resourceVM(),
		},
		DataSourcesMap: map[string]*schema.Resource{
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
		URL:   d.Get("url").(string),
		Token: d.Get("token").(string),
	}
	err := config.Validate()
	if err != nil{
		return nil,err
	}
	t := runtimeclient.New(config.URL, "/api", []string{"https"})
	t.DefaultAuthentication = runtimeclient.APIKeyAuth("Authorization", "header", "Token "+config.Token)
	c := netboxClient.New(t, strfmt.Default)
	return &c, nil
}