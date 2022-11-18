package work

import (
	"github.com/sirupsen/logrus"
	"gowechat/client"
	"gowechat/constant"
	"sync"
)

type WechatWork struct {
	Cfg        constant.WorkConfig
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
			wk := WechatWork{Cfg: cfg}
			clients := make(map[string]*WorkClient)
			clients[ContactClientName] = NewWorkClient(cfg.CorpID, cfg.ContactSecret)
			clients[CustomerClientName] = NewWorkClient(cfg.CorpID, cfg.CustomerSecret)
			clients[AgentClientName] = NewWorkClient(cfg.CorpID, cfg.AgentSecret)
			wk.clients = clients
			WechatWorkInstance = &wk
		})
	}
	return WechatWorkInstance
}

func (c *WechatWork) User() *User {
	return NewUser(c)
}
func (c *WechatWork) Department() *Department {
	return NewDepartment(c)
}
func (c *WechatWork) Tag() *Tag {
	return NewTag(c)
}

func (c *WechatWork) GetClient(clientName string) *client.Client {
	ct := c.clients[clientName]
	if ct == nil {
		logrus.Error(clientName + " Client not registered")
		return nil
	}
	nameClient := ct.Client
	nameClient.SetUrlQuery(WithAccessToken{AccessToken: ct.Client.GetToken()})
	if clientName == AgentClientName {
		nameClient.SetUrlQuery(WithAgentId{AgentId: c.Cfg.AgentID})
	}
	return nameClient
}
