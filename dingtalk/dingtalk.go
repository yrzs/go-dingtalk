package dingtalk

// DingConfig 钉钉实例配置
type DingConfig struct {
	AppId        string
	AgentId      string
	ClientId     string
	ClientSecret string
}

type Client struct {
	conf *DingConfig
}

// NewClient 创建钉钉实例
func NewClient(conf *DingConfig) (*Client, error) {
	return &Client{
		conf: conf,
	}, nil
}
