package work

import (
	"gowechat/client"
)

type User struct {
	WechatWork
}

// UserIds get请求userid的数据
type UserIds struct {
	UserId string `url:"userid" json:"userid"`
}

// UserDepartmentId get请求department_id的数据
type UserDepartmentId struct {
	DepartmentId string `url:"department_id"`
}

func NewUser(work WechatWork) *User {
	u := &User{work}
	return u
}

type UserCreateReq struct {
	UserDetail
}

// Create 创建成员
// https://developer.work.weixin.qq.com/document/path/90195
func (u User) Create(req UserCreateReq) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type UserGetResp struct {
	client.BaseResp
	UserDetail
}

// Get 读取成员
// Get https://developer.work.weixin.qq.com/document/path/90196
func (u User) Get(userId string) (UserGetResp, error) {
	var resp UserGetResp
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/get", UserIds{userId}, &resp)
	if err != nil {
		return UserGetResp{}, err
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
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/create", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Delete 删除成员
// Delete https://developer.work.weixin.qq.com/document/path/90198
func (u User) Delete(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/delete", UserIds{userId}, &resp)
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
	err := u.getClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/user/batchdelete", UserBatchDeleteReq{UseridList: UseridList}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Simplelist 获取部门成员
// Simplelist  https://developer.work.weixin.qq.com/document/path/90200
func (u User) Simplelist(DepartmentId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/simplelist", UserDepartmentId{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type UserListResp struct {
	client.BaseResp
	UserList []UserDetail `json:"userlist"`
}

// List 获取部门成员详情
// List https://developer.work.weixin.qq.com/document/path/90201
func (u User) List(DepartmentId string) (UserListResp, error) {
	var resp UserListResp
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/simplelist", UserDepartmentId{DepartmentId: DepartmentId}, &resp)
	if err != nil {
		return UserListResp{}, err
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
	err := u.getClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/user/convert_to_openid", UserIds{userId}, &resp)
	if err != nil {
		return UserConverToOpenidResp{}, err
	}
	return resp, nil
}

// Authsucc 二次验证
// Authsucc  https://developer.work.weixin.qq.com/document/path/90203
func (u User) Authsucc(userId string) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(ContactClientName).HttpGetAssign("/cgi-bin/user/authsucc", UserIds{userId}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type userGetUserIdReq struct {
	Mobile string `json:"mobile"`
}
type UserGetUserIdResp struct {
	client.BaseResp
	UserIds
}

// GetUserId Getuserid 手机号获取userid
// GetUserId  https://developer.work.weixin.qq.com/document/path/96267
func (u User) GetUserId(mobile string) (UserGetUserIdResp, error) {
	var resp UserGetUserIdResp
	err := u.getClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/user/getuserid", userGetUserIdReq{Mobile: mobile}, &resp)
	if err != nil {
		return UserGetUserIdResp{}, err
	}
	return resp, nil
}

type GetUseridByEmailReq struct {
	Email     string `json:"email"`
	EmailType int    `json:"email_type"`
}
type GetUseridByEmailResp struct {
	client.BaseResp
	UserIds
}

// GetUseridByEmail  邮箱获取userid
// GetUseridByEmail  https://developer.work.weixin.qq.com/document/path/95895
func (u User) GetUseridByEmail(email string, emailTypes ...int) (GetUseridByEmailResp, error) {
	var resp GetUseridByEmailResp
	emailType := 1
	if len(emailTypes) == 1 {
		emailType = emailTypes[0]
	}
	err := u.getClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/user/get_userid_by_email", GetUseridByEmailReq{Email: email, EmailType: emailType}, &resp)
	if err != nil {
		return GetUseridByEmailResp{}, err
	}
	return resp, nil
}
