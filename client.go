package palworld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	c *resty.Client
}

type Option func(*Client)

func WithTimeout(timeout time.Duration) func(*Client) {
	return func(c *Client) {
		c.c.SetTimeout(timeout)
	}
}

func New(baseURL string, password string, options ...Option) *Client {
	c := &Client{resty.New()}

	c.c.SetDisableWarn(true)
	c.c.SetBaseURL(baseURL)
	c.c.SetBasicAuth("admin", password)

	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) get(path string, res any) error {
	resp, err := c.c.R().
		SetHeader("Accept", "application/json").
		Get(path)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Body(), res); err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("palworld client error: %s", resp.Status())
	}

	return nil
}

func (c *Client) post(path string, body any) error {
	resp, err := c.c.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(path)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("palworld client error: %s", resp.Status())
	}

	return nil
}
