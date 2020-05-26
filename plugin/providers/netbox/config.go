package netbox

import (
	"errors"

	"github.com/netbox-community/go-netbox/netbox/client"
)

type Config struct {
	URL   string
	Token string
}

type ProviderNetboxClient struct {
	netboxClient  *client.NetBox
	configuration Config
}

func (c *Config) Validate() error {
	if c.Token == "" {
		return errors.New("token must be specified")
	}
	if c.URL == "" {
		return errors.New("url must be specified")
	}
	return nil
}

func (client *ProviderNetboxClient) CheckConnection() error {
	_, err := client.netboxClient.Dcim.DcimRacksList(nil, nil)
	if err != nil {
		return err
	}
	return nil
}
