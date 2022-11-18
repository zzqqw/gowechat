package work

import "gowechat/client"

type Agent struct {
	work *WechatWork
}

func NewAgent(work *WechatWork) *Agent {
	return &Agent{work}
}

type AgentGetRep struct {
	client.BaseResp
	AgentDetail
	SquareLogoURL           string              `json:"square_logo_url"`
	AllowUserinfos          AgentGetUsers       `json:"allow_userinfos"`
	AllowPartys             AgentGetAllowPartys `json:"allow_partys"`
	AllowTags               AgentGetAllowTags   `json:"allow_tags"`
	Close                   int                 `json:"close"`
	CustomizedPublishStatus int                 `json:"customized_publish_status"`
}

type AgentGetUser struct {
	Userid string `json:"userid"`
}

type AgentGetUsers struct {
	User []AgentGetUser `json:"user"`
}

type AgentGetAllowPartys struct {
	Partyid []int `json:"partyid"`
}

type AgentGetAllowTags struct {
	Tagid []int `json:"tagid"`
}

// Get  获取指定的应用详情
// https://developer.work.weixin.qq.com/document/path/90227
func (b *Agent) Get() (resp AgentGetRep, err error) {
	err = b.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/agent/get", nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Set  设置应用
// https://developer.work.weixin.qq.com/document/path/90228
func (b *Agent) Set(req AgentGetRep) (resp client.BaseResp, err error) {
	err = b.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/agent/set", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
