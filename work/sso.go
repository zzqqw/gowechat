package work

import (
	"fmt"
	"gowechat/client"
	"gowechat/helper/encoding/url"
)

type Sso struct {
	work *WechatWork
}

func NewSso(work *WechatWork) *Sso {
	return &Sso{work}
}

type QrConnectResp struct {
	LocationURL string `json:"location_url"`
	RedirectUri string `json:"redirect_uri"`
	AppId       string `json:"app_id"`
	AgentId     int64  `json:"agent_id"`
	State       string `json:"state"`
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
		"%v/wwopen/sso/qrConnect?appid=%v&agentid=%v&redirect_uri=%v&state=%v",
		workBaseUrl, resp.AppId, resp.AgentId, url.Encode(resp.RedirectUri), resp.State,
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
	err = o.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/auth/getuserinfo", CodeReq{Code: code}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
