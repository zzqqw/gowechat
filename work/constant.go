package work

// UserDetail 客户详情
type UserDetail struct {
	UserID           string                     `json:"userid"`
	OpenUserid       string                     `json:"open_userid"`
	Name             string                     `json:"name"`
	EnglishName      string                     `json:"english_name"`
	Alias            string                     `json:"alias"`
	Position         string                     `json:"position"`
	Status           int                        `json:"status"`
	IsLeader         int                        `json:"isleader"`
	TelePhone        string                     `json:"telephone"`
	Gender           string                     `json:"gender"`
	BizMail          string                     `json:"biz_mail"`
	Avatar           string                     `json:"avatar"`
	ThumbAvatar      string                     `json:"thumb_avatar"`
	QrCode           string                     `json:"qr_code"`
	Address          string                     `json:"address"`
	Enable           int                        `json:"enable"`
	Mobile           int                        `json:"mobile"`
	HideMobile       int                        `json:"hide_mobile"`
	Department       []int64                    `json:"department"`
	MainDepartment   int                        `json:"main_department"`
	Order            []uint32                   `json:"order"`
	IsLeaderInDept   []uint32                   `json:"is_leader_in_dept"`
	DirectLeader     []string                   `json:"direct_leader"`
	ExternalPosition string                     `json:"external_position"`
	Extattr          UserDetailExtattrs         `json:"extattr"`
	ExternalProfile  UserDetailExternalProfiles `json:"external_profile"`
}
type UserDetailExtattr struct {
	Type int                   `json:"type"`
	Name string                `json:"name"`
	Text UserDetailExtattrText `json:"text"`
	Web  UserDetailExtattrWeb  `json:"web"`
}
type UserDetailExtattrWeb struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
type UserDetailExtattrText struct {
	Value string `json:"value"`
}
type UserDetailExtattrs struct {
	Attrs []UserDetailExtattr `json:"attrs"`
}
type UserDetailExternalProfiles struct {
	ExternalCorpName string                    `json:"external_corp_name"`
	ExternalAttr     []UserDetailExtattr       `json:"external_attr"`
	WechatChannels   []UserDetailWechatChannel `json:"wechat_channels"`
}
type UserDetailWechatChannel struct {
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

// UserIdReq get请求userid的数据
type UserIdReq struct {
	UserId string `url:"userid"`
}

// UserIdReq get请求department_id的数据
type DepartmentIdReq struct {
	DepartmentId string `url:"department_id"`
}
