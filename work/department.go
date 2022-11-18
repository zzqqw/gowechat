package work

import (
	"gowechat/client"
)

type Department struct {
	work *WechatWork
}

func NewDepartment(work *WechatWork) *Department {
	return &Department{work}
}

type DepartmentIds struct {
	Id int `json:"id"`
}
type DepartmentCreateResp struct {
	client.BaseResp
	DepartmentIds
}

// Create 创建部门
// https://developer.work.weixin.qq.com/document/path/90205
func (u *Department) Create(req DepartmentDetail) (resp DepartmentCreateResp, err error) {
	err = u.work.GetClient(ClientNameCustomer).PostJsonAssign("/cgi-bin/department/create", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Update 更新部门
// https://developer.work.weixin.qq.com/document/path/90206
func (u *Department) Update(req DepartmentDetail) (resp client.BaseResp, err error) {
	err = u.work.GetClient(ClientNameCustomer).PostJsonAssign("/cgi-bin/department/update", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Delete 删除部门
// https://developer.work.weixin.qq.com/document/path/90207
func (u *Department) Delete(id int) (resp client.BaseResp, err error) {
	err = u.work.GetClient(ClientNameCustomer).PostJsonAssign("/cgi-bin/department/delete", DepartmentIds{
		Id: id,
	}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type DepartmentListResp struct {
	client.BaseResp
	Department []DepartmentDetail `json:"department"`
}

// List 获取部门列表
// https://developer.work.weixin.qq.com/document/path/90208
func (u *Department) List(id ...int) (resp DepartmentListResp, err error) {
	c := u.work.GetClient(ClientNameCustomer)
	if len(id) == 1 {
		c.SetUrlQuery(DepartmentIds{Id: id[0]})
	}
	err = c.SetUrlQueryValEmptyContinue().GetAssign("/cgi-bin/department/list", nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type DepartmentSimpleListResp struct {
	client.BaseResp
	DepartmentId []DepartmentDetail `json:"department_id"`
}

// SimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (u *Department) SimpleList(id ...int) (resp DepartmentSimpleListResp, err error) {
	var req = DepartmentIds{}
	if len(id) == 1 {
		req.Id = id[0]
	}
	err = u.work.GetClient(ClientNameCustomer).GetAssign("/cgi-bin/department/simplelist", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type DepartmentGetResp struct {
	client.BaseResp
	DepartmentDetail
}

// Get 获取部门列表
// https://developer.work.weixin.qq.com/document/path/95351
func (u *Department) Get(id ...int) (resp DepartmentGetResp, err error) {
	var req DepartmentIds
	if len(id) == 1 {
		req.Id = id[0]
	}
	err = u.work.GetClient(ClientNameCustomer).GetAssign("/cgi-bin/department/get", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
