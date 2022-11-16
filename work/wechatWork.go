package work

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gowechat/client"
	"gowechat/constant"
	"gowechat/util"
	"time"
)

const workBaseUrl = "https://qyapi.weixin.qq.com"

type WechatWork struct {
	clientName string
	clients    map[string]*client.Client
}

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

func NewWechatWork(cfg constant.WorkConfig) *WechatWork {
	wk := WechatWork{}
	clients := make(map[string]*client.Client)
	contact := client.NewClient(workBaseUrl, cfg.CorpID, cfg.ContactSecret)
	contact.Token.SetGetTokenFunc(func() (client.TokenInfo, error) {
		var object = accessTokenResp{}
		var req = accessTokenReq{CorpID: contact.WxId, CorpSecret: contact.WxSecret}
		err := contact.HttpGetAssign("/cgi-bin/gettoken", req, &object)

		if err != nil {
			return client.TokenInfo{}, err
		}
		if object.ErrCode == 40013 || object.ErrCode == 42001 || object.ErrCode == 640014 {
			logrus.Error(fmt.Sprintf("getToken errcode=%v,errmsg=%sv", object.ErrCode, object.ErrMsg))
			return client.TokenInfo{}, errors.New(object.ErrMsg)
		}
		var tokenInfo = client.TokenInfo{TokenStr: object.AccessToken, ExpiresIn: time.Duration(object.ExpiresInSecs)}
		logrus.Debug(fmt.Sprintf("contactClinet获取token，req=%v , resp=%v", util.JsonMarshalIndent(req), util.JsonMarshalIndent(tokenInfo)))
		return tokenInfo, nil
	})
	go contact.Token.TokenRefresher(context.Background())
	clients[contactClientName] = contact
	wk.clients = clients
	return &wk
}

func (c WechatWork) User() *User {
	return NewUser(c)
}

func (c WechatWork) GetClient(name string) *client.Client {
	ct := c.clients[name]
	ct.SetUrlQuery(withAccessToken{AccessToken: ct.Token.GetTokenStr()})
	return ct
}
