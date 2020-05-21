package netbox

import (
	"errors"
	netboxClient "github.com/netbox-community/go-netbox/netbox/client"
)

type Config struct {
	URL   string
	Token string
}

type Client struct {
	NetboxClient netboxClient.NetBox
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
