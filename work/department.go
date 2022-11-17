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
	Id int `url:"id" json:"id"`
}
type DepartmentCreateResp struct {
	client.BaseResp
	DepartmentIds
}

// Create 创建部门
// https://developer.work.weixin.qq.com/document/path/90205
func (u Department) Create(req DepartmentDetail) (DepartmentCreateResp, error) {
	var resp DepartmentCreateResp
	err := u.getClient(CustomerClientName).HttpPostJsonAssign("/cgi-bin/department/create", req, &resp)
	if err != nil {
		return DepartmentCreateResp{}, err
	}
	return resp, nil
}

// Update 更新部门
// https://developer.work.weixin.qq.com/document/path/90206
func (u Department) Update(req DepartmentDetail) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(CustomerClientName).HttpPostJsonAssign("/cgi-bin/department/update", req, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

// Delete 删除部门
// https://developer.work.weixin.qq.com/document/path/90207
func (u Department) Delete(id int) (client.BaseResp, error) {
	var resp client.BaseResp
	err := u.getClient(CustomerClientName).HttpPostJsonAssign("/cgi-bin/department/delete", DepartmentIds{
		Id: id,
	}, &resp)
	if err != nil {
		return client.BaseResp{}, err
	}
	return resp, nil
}

type DepartmentListResp struct {
	client.BaseResp
	Department []DepartmentDetail `json:"department"`
}

// List 获取部门列表
// https://developer.work.weixin.qq.com/document/path/90208
func (u Department) List(id ...int) (DepartmentListResp, error) {
	var resp DepartmentListResp
	var err error
	c := u.getClient(CustomerClientName)
	if len(id) == 1 {
		c.SetUrlQuery(DepartmentIds{Id: id[0]})
	}
	err = c.HttpGetAssign("/cgi-bin/department/list", nil, &resp)
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
