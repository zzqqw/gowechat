package work

// UserDetail 客户详情
type UserDetail struct {
	UserID          string                     `json:"userid"`
	Name            string                     `json:"name"`
	Department      []int64                    `json:"department"`
	Position        string                     `json:"position"`
	Status          int                        `json:"status"`
	IsLeader        int                        `json:"isleader"`
	Extattr         UserDetailExtattrs         `json:"extattr"`
	TelePhone       string                     `json:"telephone"`
	Enable          int                        `json:"enable"`
	HideMobile      int                        `json:"hide_mobile"`
	Order           []uint32                   `json:"order"`
	MainDepartment  int                        `json:"main_department"`
	Alias           string                     `json:"alias"`
	IsLeaderInDept  []uint32                   `json:"is_leader_in_dept"`
	DirectLeader    []string                   `json:"direct_leader"`
	ExternalProfile UserDetailExternalProfiles `json:"external_profile"`
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
	ExternalCorpName string              `json:"external_corp_name"`
	ExternalAttr     []UserDetailExtattr `json:"external_attr"`
}
type UserDetailWechatChannel struct {
	Nickname string `json:"nickname"`
}

// UserIdReq get请求userid的数据
type UserIdReq struct {
	UserId string `url:"userid"`
}

// UserIdReq get请求department_id的数据
type DepartmentIdReq struct {
	DepartmentId string `url:"department_id"`
}
