package client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"gowechat/util"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Url                      string
	UrlQuery                 []interface{}
	UrlQueryValEmptyContinue bool
	WxId                     string
	WxSecret                 string
	Resty                    *resty.Client
	GetTokenFunc             func() (TokenInfo, error)
	TokenKey                 string
	Cache                    *cache.Cache
}

type TokenInfo struct {
	Token     string        `json:"token"`
	ExpiresIn time.Duration `json:"expires_in"`
}

func NewClient(url, id, secret string) *Client {
	c := &Client{
		Url:      url,
		WxId:     id,
		WxSecret: secret,
		Resty:    resty.NewWithClient(&http.Client{}),
		// Create a cache with a default expiration time of 5 minutes, and which
		// purges expired items every 10 minutes
		//创建一个默认过期时间为5分钟的缓存
		//每10分钟清除过期项目
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	}
	c.TokenKey = c.WxId + ":" + util.EncryptForMd5(c.WxSecret+c.Url)
	return c
}

func (c *Client) SetGetTokenFunc(f func() (TokenInfo, error)) {
	c.GetTokenFunc = f
}

// SetUrlQuery   添加get请求参数
func (c *Client) SetUrlQuery(urlQuery interface{}) *Client {
	c.UrlQuery = append(c.UrlQuery, urlQuery)
	return c
}

// SetUrlQueryValEmptyContinue UrlQueryValEmptyContinue 添加get请求参数
func (c *Client) SetUrlQueryValEmptyContinue() *Client {
	c.UrlQueryValEmptyContinue = true
	return c
}

func (c *Client) GetToken() string {
	tokenInterface, b := c.Cache.Get(c.TokenKey)
	if b {
		token := util.InterfaceToString(tokenInterface)
		logrus.Debug(fmt.Sprintf("Get Token for cache cahceKey:%v cacheValue:%v", c.TokenKey, token))
		return token
	}
	tokenInfo, err := c.GetTokenFunc()
	if err != nil {
		logrus.Error("getToken failed err=" + err.Error())
		return ""
	}
	c.Cache.Set(c.TokenKey, tokenInfo.Token, (tokenInfo.ExpiresIn-10)*time.Second)
	logrus.Debug(fmt.Sprintf("Get Token for Api cahceKey:%v cacheValue:%v", c.TokenKey, tokenInfo.Token))
	return tokenInfo.Token
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
	logrus.Debug(fmt.Sprintf("resty Get: %v, assign:%v", queryHost, util.InterfaceToString(assign)))
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
	logrus.Debug(fmt.Sprintf("resty PostJson: %v, assign:%v", queryHost, util.InterfaceToString(assign)))
	return err
}

// 解析url 并拼装get请求参数
func (c *Client) composeReqUrl(path string, req []interface{}) string {
	urls, err := url.Parse(c.Url)
	if err != nil {
		panic(err)
	}
	urls.Path = path
	var rawQuery string
	for _, queryStruct := range req {
		if queryStruct == nil {
			continue
		}
		s := structs.New(queryStruct)
		for f, valInterface := range s.Map() {
			val := util.InterfaceToString(valInterface)
			if c.UrlQueryValEmptyContinue && (val == "" || val == "0") {
				continue
			}
			rawQuery += s.Field(f).Tag("json") + "=" + url.QueryEscape(val) + "&"
		}
	}
	urls.RawQuery = rawQuery
	return urls.String()
}
