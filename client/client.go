package client

import (
	"gowechat/util/requests"
	"sync"
)

type Client struct {
	BaseUrl    string
	withGetReq interface{}
	WxId       string
	WxSecret   string
	Requests   *requests.Requests
	Token      *Token
}

func NewClient(url, id, secret string) *Client {
	return &Client{
		BaseUrl:  url,
		WxId:     id,
		WxSecret: secret,
		Requests: requests.NewRequests(),
		Token:    &Token{Mutex: &sync.RWMutex{}},
	}
}

// SetWithGetReq  添加get请求参数
func (c *Client) SetWithGetReq(withGetReq interface{}) *Client {
	c.withGetReq = withGetReq
	return c
}

// WithTokenGet 携带token进行get请求
func (c *Client) WithTokenGet(path string, req interface{}, object interface{}) error {
	return c.Requests.SetBaseURL(c.BaseUrl).SetPath(path).SetGetReq(c.withGetReq).SetGetReq(req).GetForObject(object)
}

// WithTokenPostJson 携带token进行post请求
func (c *Client) WithTokenPostJson(path string, req interface{}, object interface{}) error {
	return c.Requests.SetBaseURL(c.BaseUrl).SetPath(path).SetGetReq(c.withGetReq).PostJsonForObject(req, object)
}
