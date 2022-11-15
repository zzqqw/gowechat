package work

import "gowechat/client"

type User struct {
	*ContactClinet
}

func NewUser(work WechatWork) *User {
	u := &User{NewContactClinet(work)}
	return u
}

type UserCreateReq struct {
	UserDetail
}

// Create 创建成员
// https://developer.work.weixin.qq.com/document/path/90195
func (u User) Create(req UserCreateReq) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.HttpPostJsonAssign("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type userGetResp struct {
	client.BaseResp
	UserDetail
}

// Get 读取成员
// Get https://developer.work.weixin.qq.com/document/path/90196
func (u User) Get(userId string) (userGetResp, error) {
	var resp userGetResp
	err := u.HttpGetAssign("/cgi-bin/user/get", UserIdReq{userId}, &resp)
	if err != nil {
		return userGetResp{}, err
	}
	return resp, nil
}

type UserUpdateReq struct {
	UserDetail
}

// Update 更新成员
// Update https://developer.work.weixin.qq.com/document/path/90197
func (u User) Update(req UserUpdateReq) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.HttpGetAssign("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Delete 删除成员
// Delete https://developer.work.weixin.qq.com/document/path/90198
func (u User) Delete(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.HttpGetAssign("/cgi-bin/user/delete", UserIdReq{userId}, &resp)
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
	err := u.HttpPostJsonAssign("/cgi-bin/user/batchdelete", UserBatchDeleteReq{UseridList: UseridList}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Simplelist 获取部门成员
// Simplelist  https://developer.work.weixin.qq.com/document/path/90200
func (u User) Simplelist(DepartmentId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.HttpGetAssign("/cgi-bin/user/simplelist", DepartmentIdReq{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type userListResp struct {
	client.BaseResp
	UserList []UserDetail `json:"userlist"`
}

// List 获取部门成员详情
// List https://developer.work.weixin.qq.com/document/path/90201
func (u User) List(DepartmentId string) (userListResp, error) {
	var resp userListResp
	err := u.HttpGetAssign("/cgi-bin/user/simplelist", DepartmentIdReq{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return userListResp{}, err
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
	err := u.HttpGetAssign("/cgi-bin/user/convert_to_openid", UserIdReq{userId}, &resp)
	if err != nil {
		return UserConverToOpenidResp{}, err
	}
	return resp, nil
}

// Authsucc 二次验证
// Authsucc  https://developer.work.weixin.qq.com/document/path/90203
func (u User) Authsucc(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.HttpGetAssign("/cgi-bin/user/authsucc", UserIdReq{userId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}
