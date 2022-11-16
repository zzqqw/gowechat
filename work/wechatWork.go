package work

import (
	"fmt"
	"gowechat/client"
	"gowechat/constant"
)

const workBaseUrl = "https://qyapi.weixin.qq.com"

type WechatWork struct {
	clientName string
	clients    map[string]*client.Client
}

func NewWechatWork(cfg constant.WorkConfig) *WechatWork {
	wk := WechatWork{}
	clients := make(map[string]*client.Client)
	clients[contactClientName] = NewWorkClient(cfg.CorpID, cfg.ContactSecret).GetClient()
	wk.clients = clients
	return &wk
}

func (c WechatWork) User() *User {
	return NewUser(c)
}

func (c WechatWork) GetClient(name string) *client.Client {
	ct := c.clients[name]
	fmt.Println(ct)
	return ct
}
