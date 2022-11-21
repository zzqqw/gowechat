package work

import (
	"fmt"
	"gowechat/client"
	"gowechat/helper/encoding/url"
)

type Oauth2 struct {
	work *WechatWork
}

func NewOauth2(work *WechatWork) *Oauth2 {
	return &Oauth2{work}
}

type Auth2AuthorizeResp struct {
	LocationURL  string `json:"location_url"`
	RedirectUri  string `json:"redirect_uri"`
	AppId        string `json:"app_id"`
	ResponseType string `json:"response_type"`
	AgentId      int64  `json:"agent_id"`
	State        string `json:"state"`
	Scope        string `json:"scope"`
}

// Authorize 构造网页授权链接
//https://developer.work.weixin.qq.com/document/path/91022
func (o *Oauth2) Authorize(redirectUri string, states ...string) (resp Auth2AuthorizeResp) {
	var state string
	if len(states) == 1 {
		state = states[0]
	} else {
		state = "STATE"
	}
	resp.RedirectUri = redirectUri
	resp.AppId = o.work.Cfg.CorpID
	resp.ResponseType = "code"
	resp.Scope = "snsapi_base"
	resp.State = state
	resp.AgentId = o.work.Cfg.AgentID
	resp.LocationURL = fmt.Sprintf(
		"%s/connect/oauth2/authorize?appid=%v&redirect_uri=%v&response_type=%v&scope=%v&state=%v&agentid=%v#wechat_redirect",
		workBaseUrl, resp.AppId, url.Encode(redirectUri), resp.ResponseType, resp.Scope, state, resp.AgentId,
	)
	return resp
}

type Oauth2GetUserInfoResp struct {
	client.BaseResp
	UserId
	UserTicket string `json:"user_ticket"`
}

// GetUserInfo 获取访问用户身份
//https://developer.work.weixin.qq.com/document/path/91023
func (o *Oauth2) GetUserInfo(code string) (resp Oauth2GetUserInfoResp, err error) {
	err = o.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/auth/getuserinfo", CodeReq{Code: code}, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type Auth2GetUserDetailReq struct {
	UserTicket string `json:"user_ticket"`
}
type Oauth2GetUserDetailResp struct {
	client.BaseResp
	UserId
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
	QrCode  string `json:"qr_code"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	BizMail string `json:"biz_mail"`
	Address string `json:"address"`
}

// GetUserDetail 获取访问用户敏感信息
//https://developer.work.weixin.qq.com/document/path/95833
func (o *Oauth2) GetUserDetail(userTicket string) (resp Oauth2GetUserDetailResp, err error) {
	err = o.work.GetClient(ClientNameAgent).PostJsonAssign("/cgi-bin/auth/getuserdetail", Auth2GetUserDetailReq{
		UserTicket: userTicket,
	}, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
