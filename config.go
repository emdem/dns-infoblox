package dnsinfoblox

import (
	"fmt"
	"log"

	infoblox "github.com/fanatic/go-infoblox"
)

type Config struct {
	User string
	Password string
	Url string
}

//Client returns a new client for accessing dnsinfoblox
func(c *Config) Client() (*dnsinfoblox.Client, error) {
	client, err := infoblox.NewClient(c.Url, c.User, c.Password, false, false)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] DNSInflbox Client configured for user: %s", client.User)
	return client, nil
}
