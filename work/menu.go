package work

type Menu struct {
	work *WechatWork
}

func NewMenu(work *WechatWork) *Menu {
	return &Menu{work}
}
