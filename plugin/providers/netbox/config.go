package netbox

import "fmt"

type Config struct {
	Token string
	Url   string
}

func (c *Config) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("token must be specified")
	}
	if c.Url == "" {
		return fmt.Errorf("url must be specified")
	}
	return nil
}
