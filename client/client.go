package client

import (
	"context"
	"gowechat/util/requests"
	"sync"
)

type Client struct {
	BaseUrl  string
	Req      requests.Requests
	WechatId string
	Secret   string
	Token    *Token
}

func NewClient(BaseUrl, wechatId, secret string) Client {
	c := Client{
		BaseUrl:  BaseUrl,
		Req:      requests.NewRequests(),
		WechatId: wechatId,
		Secret:   secret,
		Token:    &Token{Mutex: &sync.RWMutex{}},
	}
	return c
}

// SetGetTokenFunc 设置获取token函数
func (c Client) SetGetTokenFunc(f func() (TokenInfo, error)) {
	c.Token.setGetTokenFunc(f)
}

// TokenRefresher 刷新token
func (c Client) TokenRefresher() {
	go c.Token.tokenRefresher(context.Background())
}

// WithTokenGet 携带token进行get请求
func (c Client) WithTokenGet(path string, req interface{}, object interface{}) error {
	return c.Req.SetBaseURL(c.BaseUrl).SetPath(path).
		SetGetReq(withAccessToken{AccessToken: c.Token.getToken()}).
		SetGetReq(req).
		GetForObject(object)
}

// WithTokenPostJson 携带token进行post请求
func (c Client) WithTokenPostJson(path string, req interface{}, object interface{}) error {
	return c.Req.SetBaseURL(c.BaseUrl).SetPath(path).
		SetGetReq(withAccessToken{AccessToken: c.Token.getToken()}).
		PostJsonForObject(req, object)
}
