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
	err = m.work.GetClient(ClientNameAgent).SetUrlQuery(req).UploadAssign("/cgi-bin/media/upload", "media", req.Media, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// MediaUploadImgResp 永久图片素材上传响应
type MediaUploadImgResp struct {
	client.BaseResp
	URL string `json:"url"`
}

func (m *Media) UploadImg(fileName string, buf []byte) (resp MediaUploadImgResp, err error) {
	req := client.Media{
		FileName: fileName,
		Reader:   bytes.NewReader(buf),
	}
	err = m.work.GetClient(ClientNameAgent).UploadAssign("/cgi-bin/media/uploadimg", "media", req, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
