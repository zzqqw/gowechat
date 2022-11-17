package work

import (
	"github.com/sirupsen/logrus"
	"gowechat/client"
	"gowechat/constant"
	"sync"
)

type WechatWork struct {
	clientName string
	clients    map[string]*WorkClient
}

var (
	WechatWorkInstance *WechatWork
	once               = &sync.Once{}
)

func NewWechatWork(cfg constant.WorkConfig) *WechatWork {
	if WechatWorkInstance == nil {
		once.Do(func() {
			wk := WechatWork{}
			clients := make(map[string]*WorkClient)
			clients[ContactClientName] = NewWorkClient(cfg.CorpID, cfg.ContactSecret)
			clients[CustomerClientName] = NewWorkClient(cfg.CorpID, cfg.CustomerSecret)
			wk.clients = clients
			WechatWorkInstance = &wk
		})
	}
	return WechatWorkInstance
}

func (c WechatWork) User() *User {
	return NewUser(c)
}
func (c WechatWork) Department() *Department {
	return NewDepartment(c)
}
func (c WechatWork) getClient(clientName string) *client.Client {
	ct := c.clients[clientName]
	if ct == nil {
		logrus.Error(clientName + " Client not registered")
		return nil
	}
	ct.Client.SetUrlQuery(withAccessToken{AccessToken: ct.Client.GetToken()})
	return ct.GetClient()
}
