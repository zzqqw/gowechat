package work

import (
	"gowechat/client"
)

type Department struct {
	WechatWork
}

func NewDepartment(work WechatWork) *Department {
	return &Department{work}
}

type DepartmentIds struct {
	Id int `url:"id"`
}
type DepartmentListResp struct {
	client.BaseResp
	Department []DepartmentDetail `json:"department"`
}

// List 获取部门列表
// https://developer.work.weixin.qq.com/document/path/90208
func (u Department) List(id ...int) (DepartmentListResp, error) {
	var resp DepartmentListResp
	var req = DepartmentIds{}
	if len(id) == 1 {
		req.Id = id[0]
	}
	err := u.getClient(CustomerClientName).HttpGetAssign("/cgi-bin/department/list", req, &resp)
	if err != nil {
		return DepartmentListResp{}, err
	}
	return resp, nil
}

type DepartmentSimpleListResp struct {
	client.BaseResp
	DepartmentId []DepartmentDetail `json:"department_id"`
}

// SimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (u Department) SimpleList(id ...int) (DepartmentSimpleListResp, error) {
	var resp DepartmentSimpleListResp
	var req = DepartmentIds{}
	if len(id) == 1 {
		req.Id = id[0]
	}
	err := u.getClient(CustomerClientName).HttpGetAssign("/cgi-bin/department/simplelist", req, &resp)
	if err != nil {
		return DepartmentSimpleListResp{}, err
	}
	return resp, nil
}

type DepartmentGetResp struct {
	client.BaseResp
	DepartmentDetail
}

//type DepartmentIds struct {
//	Id int `url:"id"`
//}

// Get 获取部门列表
// https://developer.work.weixin.qq.com/document/path/95351
func (u Department) Get(id ...int) (DepartmentGetResp, error) {
	var resp DepartmentGetResp
	var req DepartmentIds
	if len(id) == 1 {
		req.Id = id[0]
	}
	err := u.getClient(CustomerClientName).HttpGetAssign("/cgi-bin/department/get", req, &resp)
	if err != nil {
		return DepartmentGetResp{}, err
	}
	return resp, nil
}
