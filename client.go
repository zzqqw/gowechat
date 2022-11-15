package gowechat

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type client struct {
	req *resty.Client
}

func NewClient() (c *client) {
	c = &client{
		req: resty.NewWithClient(&http.Client{}),
	}
	return c
}

func (c client) clone() *resty.Client {
	return c.req
}

func (c client) Get(url string) (*resty.Response, error) {
	return c.clone().R().Get(url)
}
