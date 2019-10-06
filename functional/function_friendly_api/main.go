package main

import (
	"errors"
	"fmt"
)

// Client type
type Client struct {
	channelSecret string
	channelToken  string
	endPointBase  string
	tierLevel     string
}

// ClientOptionHandlerFunc handles functional option
type ClientOptionHandlerFunc func(*Client) error

// NewClient an instance example object
func NewClient(e string, options ...ClientOptionHandlerFunc) (*Client, error) {
	c := &Client{
		endPointBase: e,
	}

	return c, nil
}

// Use use applied middleware
func (c *Client) Use(options ...ClientOptionHandlerFunc) error {
	for _, option := range options {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// SendBroadcast sends example method
func (c *Client) SendBroadcast() bool {
	return true
}

// ----

// Secret option
func Secret(s string) ClientOptionHandlerFunc {
	return func(c *Client) error {
		if len(s) == 0 {
			return errors.New("invalid secret")
		}
		c.channelSecret = s

		return nil
	}
}

// Token option
func Token(t string) ClientOptionHandlerFunc {
	return func(c *Client) error {
		if len(t) == 0 {
			return errors.New("invalid token")
		}
		c.channelToken = t

		return nil
	}
}

// Tier option
func Tier(t int) ClientOptionHandlerFunc {
	return func(c *Client) error {
		if t < 0 {
			return errors.New("invalid tier level")
		}

		switch t := t; {
		case t > 5:
			c.tierLevel = "high"
		default:
			c.tierLevel = "low"
		}

		return nil
	}
}

func main() {
	client, _ := NewClient("<this_is_endpoint>")
	client.Use(Secret("<this_is_secret>"))
	client.Use(Token("<this_is_token>"))
	client.Use(Tier(6))

	fmt.Printf("%+v", client)
}
