package constant

type WorkConfig struct {
	// CorpID 外部企业ID
	CorpID string `json:"corp_id" validate:"required,corp_id"`
	//企业主应用AgentID
	AgentID int64 `json:"agent_id" validate:"number,gt=0"`
	//企业主应用secret
	AgentSecret string `json:"agent_secret" validate:"required"`
	// ContactSecret 通讯录secret
	ContactSecret string `json:"contact_secret" validate:"required"`
	// CustomerSecret 客户联系secret
	CustomerSecret string `json:"customer_secret" validate:"required"`
	// CallbackToken 企业微信事件回调Token
	CallbackToken string `json:"callback_token" validate:"required"`
	// CallbackAesKey 企业微信事件回调AesKey
	CallbackAesKey string `json:"callback_aes_key" validate:"required"`
	// PriKeyPath 会话存档解密私钥
	PriKeyPath string `json:"pri_key_path"`
	// MsgArchBatchSize 会话存档拉取，每次拉取的条数
	MsgArchBatchSize int `json:"msg_arch_batch_size"`
	// MsgArchTimeout 会话存档拉取，超时时间
	MsgArchTimeout int `json:"msg_arch_timeout"`
	// MsgArchProxy  会话存档拉取代理地址
	MsgArchProxy string `json:"msg_arch_proxy"`
	// MsgArchProxyPasswd  会话存档拉取代理密码
	MsgArchProxyPasswd string `json:"msg_arch_proxy_passwd"`
}
