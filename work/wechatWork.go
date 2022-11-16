package work

import (
	"github.com/sirupsen/logrus"
	"gowechat/client"
	"gowechat/constant"
)

const workBaseUrl = "https://qyapi.weixin.qq.com"

type WechatWork struct {
	clientName string
	clients    map[string]*WorkClient
}

func NewWechatWork(cfg constant.WorkConfig) *WechatWork {
	wk := WechatWork{}
	clients := make(map[string]*WorkClient)
	clients[contactClientName] = NewWorkClient(cfg.CorpID, cfg.ContactSecret)
	wk.clients = clients
	return &wk
}

func (c WechatWork) User() *User {
	return NewUser(c)
}

func (c WechatWork) GetClient(clientName string) *client.Client {
	ct := c.clients[clientName]
	if ct == nil {
		logrus.Error(clientName + " Client not registered")
	}
	ct.Client.SetUrlQuery(withAccessToken{AccessToken: ct.Client.GetToken()})
	return ct.GetClient()
}
