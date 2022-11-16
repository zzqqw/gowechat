package work

import (
	"errors"
	"gowechat/client"
	"time"
)

const contactClientName = "contact"

type accessTokenResp struct {
	client.BaseResp
	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}
type accessTokenReq struct {
	CorpID     string `url:"corpid"`
	CorpSecret string `url:"corpsecret"`
}
type withAccessToken struct {
	AccessToken string `url:"access_token"`
}

type WorkClient struct {
	CorpID        string
	ContactSecret string
	Client        *client.Client
}

func NewWorkClient(CorpID, ContactSecret string) *WorkClient {
	c := WorkClient{
		CorpID:        CorpID,
		ContactSecret: ContactSecret,
		Client:        client.NewClient(workBaseUrl, CorpID, ContactSecret),
	}
	c.Client.SetGetTokenFunc(c.getToken)
	c.Client.SetUrlQuery(withAccessToken{AccessToken: c.Client.GetToken()})
	return &c
}

func (c *WorkClient) GetClient() *client.Client {
	return c.Client
}
func (c *WorkClient) getToken() (client.TokenInfo, error) {
	var object = accessTokenResp{}
	var req = accessTokenReq{CorpID: c.CorpID, CorpSecret: c.ContactSecret}
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
