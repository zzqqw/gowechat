package work

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gowechat/client"
	"gowechat/util"
	"time"
)

type contactClinet struct {
	client.Client
}

type accessTokenResp struct {
	client.BaseResp
	AccessToken   string `json:"access_token"`
	ExpiresInSecs int64  `json:"expires_in"`
}
type accessTokenReq struct {
	CorpID     string `url:"corpid"`
	CorpSecret string `url:"corpsecret"`
}
type withAccessToken struct {
	AccessToken string `url:"access_token"`
}

func NewContactClinet(wk WechatWork) *contactClinet {
	c := client.NewClient(workBaseUrl, wk.config.CorpID, wk.config.ContactSecret)
	cc := contactClinet{*c}
	cc.Token.SetGetTokenFunc(cc.getToken)
	cc.SetWithGetReq(withAccessToken{AccessToken: cc.Token.GetTokenStr()})
	//c.Requests.Req = append(c.Requests.Req, withAccessToken{AccessToken: cc.Token.GetTokenStr()})
	return &cc
}

func (c *contactClinet) getToken() (client.TokenInfo, error) {
	var object = accessTokenResp{}
	var req = accessTokenReq{CorpID: c.WxId, CorpSecret: c.WxSecret}
	err := c.Requests.SetPath("/cgi-bin/gettoken").SetBaseURL(workBaseUrl).SetGetReq(req).GetForObject(&object)
	if err != nil {
		return client.TokenInfo{}, err
	}
	//40013 排查方法: 需确认CorpID是否填写正确，在 web管理端-设置 可查看
	//42001 排查方法: access_token有时效性，需要重新获取一次
	//640014 排查方法: 判断当前空间是否没有有效的管理员
	if object.ErrCode == 40013 || object.ErrCode == 42001 || object.ErrCode == 640014 {
		logrus.Error(fmt.Sprintf("getToken errcode=%v,errmsg=%sv", object.ErrCode, object.ErrMsg))
		return client.TokenInfo{}, errors.New(object.ErrMsg)
	}
	var tokenInfo = client.TokenInfo{TokenStr: object.AccessToken, ExpiresIn: time.Duration(object.ExpiresInSecs)}
	logrus.Debug(fmt.Sprintf("contactClinet获取token，req=%v , resp=%v", util.JsonMarshalIndent(req), util.JsonMarshalIndent(tokenInfo)))
	return tokenInfo, nil
}
