package netbox

import (
	"fmt"
	netboxClient "github.com/netbox-community/go-netbox/netbox/client"
)

type Config struct {
	URL   string
	Token string
}

type Client struct {
	NetboxClient netboxClient.NetBox
}

func (c *Config) Validate()  error {
	if c.Token == "" {
		return fmt.Errorf("token must be specified")
	}
	if c.URL == "" {
		return fmt.Errorf("url must be specified")
	}
	return nil
}

