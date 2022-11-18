package work

import "gowechat/client"

type Tag struct {
	work *WechatWork
}

func NewTag(work *WechatWork) *Tag {
	return &Tag{work}
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
func (t *Tag) Create(req TagDetail) (resp TagCreateResp, err error) {
	err = t.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/tag/create", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Update 更新标签名字
// https://developer.work.weixin.qq.com/document/path/90211
func (t *Tag) Update(req TagDetail) (resp client.BaseResp, err error) {
	err = t.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/tag/update", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Delete  更新标签名字
// https://developer.work.weixin.qq.com/document/path/90212
func (t *Tag) Delete(req TagId) (resp client.BaseResp, err error) {
	err = t.work.GetClient(ClientNameContact).GetAssign("/cgi-bin/tag/delete", req, &resp)
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
func (t *Tag) get(req TagId) (resp TagGetResp, err error) {
	err = t.work.GetClient(ClientNameContact).GetAssign("/cgi-bin/tag/get", req, &resp)
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
func (t *Tag) AddTagUsers(req TagUsersAddOrDel) (resp client.BaseResp, err error) {
	err = t.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/tag/addtagusers", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DelTagUsers  删除标签成员
// https://developer.work.weixin.qq.com/document/path/90215
func (t *Tag) DelTagUsers(req TagUsersAddOrDel) (resp client.BaseResp, err error) {
	err = t.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/tag/deltagusers", req, &resp)
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
func (t *Tag) List() (resp TagListResp, err error) {
	err = t.work.GetClient(ClientNameContact).GetAssign("/cgi-bin/tag/list", nil, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
