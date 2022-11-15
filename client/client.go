package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"gowechat/util"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	BaseUrl  string
	UrlQuery []interface{}
	WxId     string
	WxSecret string
	Token    *Token
	Resty    *resty.Client
}

func NewClient(url, id, secret string) *Client {
	return &Client{
		BaseUrl:  url,
		WxId:     id,
		WxSecret: secret,
		Token:    &Token{Mutex: &sync.RWMutex{}},
		Resty:    resty.NewWithClient(&http.Client{}),
	}
}

// SetUrlQuery   添加get请求参数
func (c *Client) SetUrlQuery(urlQuery interface{}) *Client {
	c.UrlQuery = append(c.UrlQuery, urlQuery)
	return c
}

// HttpGetAssign Get 请求并渲染struct
func (c *Client) HttpGetAssign(path string, req interface{}, assign interface{}) error {
	reqs := append(c.UrlQuery, req)
	queryHost := c.composeReqUrl(path, reqs)
	resp, err := c.Resty.R().Get(queryHost)
	if err != nil {
		return err
	}
	bodyResp := resp.Body()
	err = json.Unmarshal(bodyResp, &assign)
	logrus.Debug(fmt.Sprintf("resty Get: %v, assign:%v", queryHost, util.JsonMarshalIndent(assign)))
	return err
}

// HttpPostJsonAssign  Post 请求并渲染struct
func (c *Client) HttpPostJsonAssign(path string, body interface{}, assign interface{}) error {
	queryHost := c.composeReqUrl(path, c.UrlQuery)
	resp, err := c.Resty.
		SetHeader("Content-Type", "application/json").
		R().
		SetBody(body).
		Post(queryHost)
	if err != nil {
		return err
	}
	bodyResp := resp.Body()
	err = json.Unmarshal(bodyResp, &assign)
	logrus.Debug(fmt.Sprintf("resty PostJson: %v, assign:%v", queryHost, util.JsonMarshalIndent(assign)))
	return err
}

func (c *Client) composeReqUrl(path string, req []interface{}) string {
	urls, err := url.Parse(c.BaseUrl)
	if err != nil {
		panic(err)
	}
	urls.Path = path
	var rawQuery string
	for i := range req {
		values, _ := query.Values(req[i])
		rawQuery += values.Encode() + "&"
	}
	urls.RawQuery = rawQuery
	return urls.String()
}
