package work

import "gowechat/client"

type Tag struct {
	work *WechatWork
}

func NewTag(work *WechatWork) *Tag {
	t := Tag{work}
	return &t
}

type TagId struct {
	TagId int `json:"tagid"`
}

type TagCreateResp struct {
	client.BaseResp
	TagId
}

// Create 创建标签
// https://developer.work.weixin.qq.com/document/path/90210
func (t *Tag) Create(req TagDetail) (TagCreateResp, error) {
	var resp TagCreateResp
	err := t.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/tag/create", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Update 更新标签名字
// https://developer.work.weixin.qq.com/document/path/90211
func (t *Tag) Update(req TagDetail) (client.BaseResp, error) {
	var resp client.BaseResp
	err := t.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/tag/update", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Delete  更新标签名字
// https://developer.work.weixin.qq.com/document/path/90212
func (t *Tag) Delete(req TagId) (client.BaseResp, error) {
	var resp client.BaseResp
	err := t.work.GetClient(ContactClientName).HttpGetAssign("/cgi-bin/tag/delete", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type TagGetResp struct {
	client.BaseResp
	UserList  []TagDetail `json:"userlist"`
	PartyList []int64     `json:"partylist"`
}

// get  获取标签成员
// https://developer.work.weixin.qq.com/document/path/90213
func (t *Tag) get(req TagId) (TagGetResp, error) {
	var resp TagGetResp
	err := t.work.GetClient(ContactClientName).HttpGetAssign("/cgi-bin/tag/get", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type TagUsersAddOrDel struct {
	TagId     int64    `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int    `json:"partylist"`
}

// AddTagUsers  增加标签成员
// https://developer.work.weixin.qq.com/document/path/90214
func (t *Tag) AddTagUsers(req TagUsersAddOrDel) (client.BaseResp, error) {
	var resp client.BaseResp
	err := t.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/tag/addtagusers", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DelTagUsers  删除标签成员
// https://developer.work.weixin.qq.com/document/path/90215
func (t *Tag) DelTagUsers(req TagUsersAddOrDel) (client.BaseResp, error) {
	var resp client.BaseResp
	err := t.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/tag/deltagusers", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type TagListResp struct {
	client.BaseResp
	TagList []TagDetail `json:"taglist"`
}

// List 获取标签列表
// https://developer.work.weixin.qq.com/document/path/90216
func (t *Tag) List() (TagListResp, error) {
	var resp TagListResp
	err := t.work.GetClient(ContactClientName).HttpGetAssign("/cgi-bin/tag/list", nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
