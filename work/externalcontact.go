package work

import "gowechat/client"

type ExternalContact struct {
	work *WechatWork
}

func NewExternalContact(work *WechatWork) *ExternalContact {
	return &ExternalContact{work}
}

type ExternalContactGetFollowUserListReq struct {
	client.BaseResp
	FollowUser []string `json:"follow_user"`
}

// GetFollowUserList 获取配置了客户联系功能的成员列表
// https://developer.work.weixin.qq.com/document/path/92576
func (u *ExternalContact) GetFollowUserList() (resp ExternalContactGetFollowUserListReq, err error) {
	err = u.work.GetClient(ClientNameCustomer).GetAssign("/cgi-bin/externalcontact/get_follow_user_list", nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ExternalContactListReq struct {
	client.BaseResp
	ExternalUserid []string `json:"external_userid"`
}

// List 获取客户列表
// https://developer.work.weixin.qq.com/document/path/92264
func (u *ExternalContact) List(userId string) (resp ExternalContactListReq, err error) {
	err = u.work.GetClient(ClientNameCustomer).GetAssign("/cgi-bin/externalcontact/list", UserId{
		UserId: userId,
	}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Get 获取客户详情
// https://developer.work.weixin.qq.com/document/path/92265
func (u *ExternalContact) Get(userId string) (resp ExternalContactListReq, err error) {
	err = u.work.GetClient(ClientNameCustomer).GetAssign("/cgi-bin/externalcontact/get", UserId{
		UserId: userId,
	}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
