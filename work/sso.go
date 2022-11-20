package work

import (
	"fmt"
	"gowechat/client"
)

type Sso struct {
	work *WechatWork
}

func NewSso(work *WechatWork) *Sso {
	return &Sso{work}
}

type QrConnectResp struct {
	LocationURL string
	RedirectUri string
	AppId       string
	AgentId     int64
	State       string
}

// QrConnect 构造网页授权链接
//https://developer.work.weixin.qq.com/document/path/91019
func (o *Sso) QrConnect(redirectUri string, states ...string) (resp QrConnectResp) {
	var state string
	if len(states) == 1 {
		state = states[0]
	} else {
		state = "STATE"
	}
	resp.RedirectUri = redirectUri
	resp.AppId = o.work.Cfg.CorpID
	resp.AgentId = o.work.Cfg.AgentID
	resp.State = state

	resp.LocationURL = fmt.Sprintf(
		"%s/wwopen/sso/qrConnect?appid=%s&agentid=%s&redirect_uri=%s&state=%s",
		workBaseUrl, resp.AppId, resp.AgentId, resp.RedirectUri, resp.State,
	)
	return resp
}

type SsoGetUserInfoResp struct {
	client.BaseResp
	UserId
	OpenId string `json:"openid"`
}

// GetUserInfo 获取访问用户身份
//https://developer.work.weixin.qq.com/document/path/91437
func (o *Sso) GetUserInfo(code string) (resp SsoGetUserInfoResp, err error) {
	err = o.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/auth/getuserinfo", CodeReq{Code: code}, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
