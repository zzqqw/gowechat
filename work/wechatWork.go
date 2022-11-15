package work

import (
	"github.com/sirupsen/logrus"
	"gowechat/constant"
	"gowechat/util"
	"sync"
)

const workBaseUrl = "https://qyapi.weixin.qq.com"

type WechatWork struct {
	config constant.WorkConfig
}

var work *WechatWork
var once sync.Once

func NewWechatWork(cfg constant.WorkConfig) *WechatWork {
	once.Do(func() {
		work = &WechatWork{config: cfg}
	})
	return work
}

func (c WechatWork) User() *User {
	logrus.Debug("初始化成员管理:", util.JsonMarshalIndent(c.config))
	return NewUser(c)
}
