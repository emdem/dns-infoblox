package dns-infoblox

import (
	"fmt"
	"log"

	"github.com/emdem/go-infoblox"
)

type Config struct {
	User string
	Password string
	Url string
}

//Client returns a new client for accessing dnsinfoblox
func(c *Config) Client() (*dns-infoblox.Client, error) {
	client, err := dns-infoblox.NewClient(c.User, c.Password, c.Url)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] DNSInflbox Client configured for user: %s", client.User)
	return client, nil
}
