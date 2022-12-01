package work

import (
	"bytes"
	"errors"
	"gowechat/client"
)

type Media struct {
	work *WechatWork
}

func NewMedia(work *WechatWork) *Media {
	return &Media{work}
}

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeFile  = "file"
)

type MediaUploadReq struct {
	Media client.Media
	Type  string `json:"type"`
}
type MediaUploadResp struct {
	client.BaseResp
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// Upload 上传临时素材
//https://developer.work.weixin.qq.com/document/path/90253
func (m *Media) Upload(fileName, mediaType string, buf []byte) (resp MediaUploadResp, err error) {
	if mediaType != MediaTypeImage || mediaType != MediaTypeVoice || mediaType != MediaTypeVideo || mediaType != MediaTypeFile {
		return resp, errors.New("mediaType error must is oney [" + MediaTypeImage + "," + MediaTypeVoice + "," + "," + MediaTypeVideo + "," + MediaTypeFile + "]")
	}
	req := MediaUploadReq{
		Media: client.Media{
			FileName: fileName,
			Reader:   bytes.NewReader(buf),
		},
		Type: mediaType,
	}
	err = m.work.GetClient(ClientNameAgent).SetUrlQuery(req).UploadAssign("/cgi-bin/media/upload", "media", req.Media, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type MediaId struct {
	MediaID string `json:"media_id"`
}

// MediaUploadImgResp 永久图片素材上传响应
type MediaUploadImgResp struct {
	client.BaseResp
	URL string `json:"url"`
}

// UploadImg  上传图片
//https://developer.work.weixin.qq.com/document/path/90256
func (m *Media) UploadImg(fileName string, buf []byte) (resp MediaUploadImgResp, err error) {
	req := client.Media{
		FileName: fileName,
		Reader:   bytes.NewReader(buf),
	}
	err = m.work.GetClient(ClientNameAgent).UploadAssign("/cgi-bin/media/uploadimg", "media", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Get 获取临时素材
//https://developer.work.weixin.qq.com/document/path/90254
func (m *Media) get(mediaId string) (resp client.BaseResp, err error) {
	err = m.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/media/get", MediaId{MediaID: mediaId}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetJsSdk  获取高清语音素材
//https://developer.work.weixin.qq.com/document/path/90255
func (m *Media) getJsSdk(mediaId string) (resp client.BaseResp, err error) {
	err = m.work.GetClient(ClientNameAgent).GetAssign("/cgi-bin/media/jssdk", MediaId{MediaID: mediaId}, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type MediaUploadByUrlReq struct {
	Scene    int    `json:"scene"`
	Type     string `json:"type"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Md5      string `json:"md5"`
}

// UploadByUrl 异步上传临时素材
//https://developer.work.weixin.qq.com/document/path/96219
func (m *Media) UploadByUrl(req MediaUploadByUrlReq) (resp JobIdResp, err error) {
	err = m.work.GetClient(ClientNameAgent).PostJsonAssign("/cgi-bin/media/upload_by_url", req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
