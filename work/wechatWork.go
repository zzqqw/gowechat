package work

import (
	"github.com/sirupsen/logrus"
	"gowechat/constant"
	"gowechat/util"
)

const workBaseUrl = "https://qyapi.weixin.qq.com"

type WechatWork struct {
	config constant.WorkConfig
}

func NewWechatWork(cfg constant.WorkConfig) (work *WechatWork) {
	work = &WechatWork{config: cfg}
	return work
}

func (c WechatWork) User() *User {
	logrus.Debug("初始化成员管理:", util.JsonMarshalIndent(c.config))
	return NewUser(c)
}
