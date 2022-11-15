package gowechat

import (
	"gowechat/constant"
	"gowechat/work"
)

func NewWechat() *Wechat {
	return &Wechat{}
}

type Wechat struct {
}

func (w Wechat) Work(config constant.WorkConfig) *work.WechatWork {
	return work.NewWechatWork(config)
}
