package netbox

import (
	"log"

	runtimeClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	netboxClient "github.com/netbox-community/go-netbox/netbox/client"
)

var descriptions map[string]string

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_virtual_machine": resourceVM(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_virtual_machine": dataSourceNetboxVM(),
		},
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["netbox api token"],
				DefaultFunc: schema.EnvDefaultFunc("TF_NETBOX_TOKEN", nil),
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["netbox main url"],
				DefaultFunc: schema.EnvDefaultFunc("TF_NETBOX_URL", nil),
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
	if err != nil {
		return nil, err
	}
	t := runtimeClient.New(config.URL, "/api", []string{"https"})
	t.DefaultAuthentication = runtimeClient.APIKeyAuth("Authorization", "header", "Token "+config.Token)
	c := netboxClient.New(t, strfmt.Default)
	cs := ProviderNetboxClient{
		netboxClient:  c,
		configuration: config,
	}
	connectionOK := cs.CheckConnection()
	if connectionOK != nil {
		log.Printf("[DEBUG] provider.go CheckConnection() FAILED")
	}

	return &cs, nil
}
