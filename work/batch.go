package work

import "gowechat/client"

type Batch struct {
	work *WechatWork
}

func NewBatch(work *WechatWork) *Batch {
	b := Batch{work}
	return &b
}

type SyncResp struct {
	client.BaseResp
	JobId string `json:"jobid"`
}
type SyncUser struct {
	MediaID  string       `json:"media_id"`
	ToInvite bool         `json:"to_invite"`
	Callback SyncCallback `json:"callback"`
}
type SyncCallback struct {
	URL            string `json:"url"`
	Token          string `json:"token"`
	Encodingaeskey string `json:"encodingaeskey"`
}

// SyncUser 增量更新成员
// https://developer.work.weixin.qq.com/document/path/90980
func (b *Batch) SyncUser(req SyncUser) (SyncResp, error) {
	var resp SyncResp
	err := b.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/batch/syncuser", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ReplaceUser 增量更新成员
// https://developer.work.weixin.qq.com/document/path/90981
func (b *Batch) ReplaceUser(req SyncUser) (SyncResp, error) {
	var resp SyncResp
	err := b.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/batch/replaceuser", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ReplaceParty struct {
	MediaID  string       `json:"media_id"`
	Callback SyncCallback `json:"callback"`
}

// ReplaceParty 全量覆盖部门
// https://developer.work.weixin.qq.com/document/path/90982
func (b *Batch) ReplaceParty(req ReplaceParty) (SyncResp, error) {
	var resp SyncResp
	err := b.work.GetClient(ContactClientName).HttpPostJsonAssign("/cgi-bin/tag/replaceparty", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
