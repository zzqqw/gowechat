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
func (c *WorkClient) getToken() (client.TokenInfo, error) {
	var object = AccessTokenResp{}
	var req = AccessTokenReq{CorpID: c.Client.WxId, CorpSecret: c.Client.WxSecret}
	err := c.Client.HttpGetAssign("/cgi-bin/gettoken", req, &object)
	if err != nil {
		return client.TokenInfo{}, err
	}
	if object.ErrCode == 40013 || object.ErrCode == 42001 || object.ErrCode == 640014 {
		return client.TokenInfo{}, errors.New(object.ErrMsg)
	}
	var tokenInfo = client.TokenInfo{Token: object.AccessToken, ExpiresIn: time.Duration(object.ExpiresInSecs)}
	return tokenInfo, nil
}
