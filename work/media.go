package work

type Media struct {
	work *WechatWork
}

func NewMedia(work *WechatWork) *Media {
	return &Media{work}
}
