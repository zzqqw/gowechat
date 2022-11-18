package work

import "gowechat/client"

type Batch struct {
	work *WechatWork
}

func NewBatch(work *WechatWork) *Batch {
	return &Batch{work}
}

type JobId struct {
	JobId string `json:"jobid"`
}

type SyncResp struct {
	client.BaseResp
	JobId
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
func (b *Batch) SyncUser(req SyncUser) (resp SyncResp, err error) {
	err = b.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/batch/syncuser", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ReplaceUser 增量更新成员
// https://developer.work.weixin.qq.com/document/path/90981
func (b *Batch) ReplaceUser(req SyncUser) (resp SyncResp, err error) {
	err = b.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/batch/replaceuser", req, &resp)
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
func (b *Batch) ReplaceParty(req ReplaceParty) (resp SyncResp, err error) {
	err = b.work.GetClient(ClientNameContact).PostJsonAssign("/cgi-bin/tag/replaceparty", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type SyncGetResult struct {
	client.BaseResp
	Status     int          `json:"status"`
	Type       string       `json:"type"`
	Total      int          `json:"total"`
	Percentage int          `json:"percentage"`
	Result     []SyncResult `json:"result"`
}
type SyncResult []struct {
	client.BaseResp
	Userid string `json:"userid"`
}

type SyncGetResultResp struct {
	client.BaseResp
	JobId
}

// GetResult 获取异步任务结果
// https://developer.work.weixin.qq.com/document/path/90983
func (b *Batch) GetResult(jobId string) (resp SyncGetResult, err error) {
	err = b.work.GetClient(ClientNameContact).
		GetAssign("/cgi-bin/tag/getresult", JobId{JobId: jobId}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
