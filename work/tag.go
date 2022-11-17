package work

type Tag struct {
	work *WechatWork
}

func NewTag(work *WechatWork) *Tag {
	t := Tag{work}
	return &t
}
