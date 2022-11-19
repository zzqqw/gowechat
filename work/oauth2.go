package work

type Oauth2 struct {
	work *WechatWork
}

func NewOauth2(work *WechatWork) *Oauth2 {
	return &Oauth2{work}
}

func (o Oauth2) Authorize() {

}
