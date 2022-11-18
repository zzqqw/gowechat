package work

import (
	"errors"
	"gowechat/client"
	"time"
)

const ContactClientName = "contact"
const CustomerClientName = "customer"
const AgentClientName = "agent"
const workBaseUrl = "https://qyapi.weixin.qq.com"

type AccessTokenResp struct {
	client.BaseResp
	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}
type AccessTokenReq struct {
	CorpID     string `json:"corpid"`
	CorpSecret string `json:"corpsecret"`
}
type WithAccessToken struct {
	AccessToken string `json:"access_token"`
}
type WithAgentId struct {
	AgentId int64 `json:"agentid"`
}

type WorkClient struct {
	Client *client.Client
}

func NewWorkClient(wxId, wxSecret string) *WorkClient {
	c := WorkClient{
		Client: client.NewClient(workBaseUrl, wxId, wxSecret),
	}
	c.Client.SetGetTokenFunc(c.getToken)
	//go c.Client.GetToken()
	return &c
}
func (c *WorkClient) getToken() (token client.TokenInfo, err error) {
	var object = AccessTokenResp{}
	err = c.Client.HttpGetAssign("/cgi-bin/gettoken", AccessTokenReq{CorpID: c.Client.WxId, CorpSecret: c.Client.WxSecret}, &object)
	if err != nil {
		return token, err
	}
	if object.ErrCode == 40013 || object.ErrCode == 42001 || object.ErrCode == 640014 {
		return token, errors.New(object.ErrMsg)
	}
	token.Token = object.AccessToken
	token.ExpiresIn = time.Duration(object.ExpiresInSecs)
	return token, nil
}
