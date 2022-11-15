package work

import "gowechat/client"

type User struct {
	*contactClinet
}

func NewUser(work WechatWork) *User {
	c := NewContactClinet(work)
	u := &User{c}
	return u
}

type UserIdReq struct {
	UserId string `url:"userid"`
}
type DepartmentIdReq struct {
	DepartmentId string `url:"department_id"`
}

type UserCreateReq struct {
	UserId           string             `json:"userid"`
	Name             string             `json:"name"`
	Alias            string             `json:"alias"`
	Mobile           string             `json:"mobile"`
	Department       []int64            `json:"department"`
	Order            []int64            `json:"order"`
	Position         string             `json:"position"`
	Gender           string             `json:"gender"`
	Email            string             `json:"email"`
	BizMail          string             `json:"biz_mail"`
	IsLeaderInDept   []int64            `json:"is_leader_in_dept"`
	DirectLeader     []string           `json:"direct_leader"`
	Enable           int64              `json:"enable"`
	AvatarMediaid    string             `json:"avatar_mediaid"`
	Telephone        string             `json:"telephone"`
	Address          string             `json:"address"`
	MainDepartment   string             `json:"main_department"`
	ToInvite         bool               `json:"to_invite"`
	ExternalPosition string             `json:"external_position"`
	Extattr          ExtattrsReq        `json:"extattr"`
	ExternalProfile  ExternalProfileReq `json:"external_profile"`
}
type ExternalProfileReq struct {
	ExternalCorpName string             `json:"external_corp_name"`
	WechatChannels   []WechatChannelReq `json:"wechat_channels"`
}
type WechatChannelReq struct {
	Nickname string `json:"nickname"`
}
type ExtattrReq struct {
	Type int            `json:"type"`
	Name string         `json:"name"`
	Text ExtattrTextReq `json:"text"`
	Web  ExtattrWebReq  `json:"web"`
}
type ExtattrWebReq struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
type ExtattrTextReq struct {
	Value string `json:"value"`
}
type ExtattrsReq struct {
	Attrs []extattrResp `json:"attrs"`
}

// Create 创建成员
// https://developer.work.weixin.qq.com/document/path/90195
func (u User) Create(req UserCreateReq) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenPostJson("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type extattrResp struct {
	Type int             `json:"type"`
	Name string          `json:"name"`
	Text extattrTextResp `json:"text"`
	Web  extattrWebResp  `json:"web"`
}
type extattrWebResp struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
type extattrTextResp struct {
	Value string `json:"value"`
}
type extattrsResp struct {
	Attrs []extattrResp `json:"attrs"`
}
type externalProfilesResp struct {
	ExternalCorpName string        `json:"external_corp_name"`
	ExternalAttr     []extattrResp `json:"external_attr"`
}
type userGetResp struct {
	client.BaseResp
	UserID          string               `json:"userid"`
	Name            string               `json:"name"`
	Department      []int64              `json:"department"`
	Position        string               `json:"position"`
	Status          int                  `json:"status"`
	IsLeader        int                  `json:"isleader"`
	Extattr         extattrsResp         `json:"extattr"`
	TelePhone       string               `json:"telephone"`
	Enable          int                  `json:"enable"`
	HideMobile      int                  `json:"hide_mobile"`
	Order           []uint32             `json:"order"`
	MainDepartment  int                  `json:"main_department"`
	Alias           string               `json:"alias"`
	IsLeaderInDept  []uint32             `json:"is_leader_in_dept"`
	DirectLeader    []string             `json:"direct_leader"`
	ExternalProfile externalProfilesResp `json:"external_profile"`
}

// Get 读取成员
// Get https://developer.work.weixin.qq.com/document/path/90196
func (u User) Get(userId string) (userGetResp, error) {
	var resp userGetResp
	err := u.WithTokenGet("/cgi-bin/user/get", UserIdReq{userId}, &resp)
	if err != nil {
		return userGetResp{}, err
	}
	return resp, nil
}

type UserUpdateReq struct {
	UserId           string             `json:"userid"`
	Name             string             `json:"name"`
	Alias            string             `json:"alias"`
	Mobile           string             `json:"mobile"`
	Department       []int64            `json:"department"`
	Order            []int64            `json:"order"`
	Position         string             `json:"position"`
	Gender           string             `json:"gender"`
	Email            string             `json:"email"`
	BizMail          string             `json:"biz_mail"`
	IsLeaderInDept   []int64            `json:"is_leader_in_dept"`
	DirectLeader     []string           `json:"direct_leader"`
	Enable           int64              `json:"enable"`
	AvatarMediaid    string             `json:"avatar_mediaid"`
	Telephone        string             `json:"telephone"`
	Address          string             `json:"address"`
	MainDepartment   int                `json:"main_department"`
	ToInvite         bool               `json:"to_invite"`
	ExternalPosition string             `json:"external_position"`
	Extattr          ExtattrsReq        `json:"extattr"`
	ExternalProfile  ExternalProfileReq `json:"external_profile"`
}

// Update 更新成员
// Update https://developer.work.weixin.qq.com/document/path/90197
func (u User) Update(req UserUpdateReq) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenGet("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Delete 删除成员
// Delete https://developer.work.weixin.qq.com/document/path/90198
func (u User) Delete(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenGet("/cgi-bin/user/delete", UserIdReq{userId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type UserBatchDeleteReq struct {
	UseridList []string `url:"useridlist"`
}

// BatchDelete 批量删除成员
// BatchDelete https://developer.work.weixin.qq.com/document/path/90199
func (u User) BatchDelete(UseridList []string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenPostJson("/cgi-bin/user/batchdelete", UserBatchDeleteReq{UseridList: UseridList}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Simplelist 获取部门成员
// Simplelist  https://developer.work.weixin.qq.com/document/path/90200
func (u User) Simplelist(DepartmentId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenGet("/cgi-bin/user/simplelist", DepartmentIdReq{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// List 获取部门成员详情
// List https://developer.work.weixin.qq.com/document/path/90201
func (u User) List(DepartmentId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenGet("/cgi-bin/user/simplelist", DepartmentIdReq{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type UserConverToOpenidResp struct {
	client.BaseResp
	Openid string `json:"openid"`
}

// ConvertToOpenid userid与openid互换
// ConvertToOpenid  https://developer.work.weixin.qq.com/document/path/90202
func (u User) ConvertToOpenid(userId string) (UserConverToOpenidResp, error) {
	var resp UserConverToOpenidResp
	err := u.WithTokenGet("/cgi-bin/user/convert_to_openid", UserIdReq{userId}, &resp)
	if err != nil {
		return UserConverToOpenidResp{}, err
	}
	return resp, nil
}

// Authsucc 二次验证
// Authsucc  https://developer.work.weixin.qq.com/document/path/90203
func (u User) Authsucc(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.WithTokenGet("/cgi-bin/user/authsucc", UserIdReq{userId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}
