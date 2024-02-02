package response

import "time"

// AccessToken 应用access token
// 文档: https://open.dingtalk.com/document/orgapp/obtain-user-token
type AccessToken struct {
	*BasicResponse // 此接口未返回错误代码信息，仅仅能检查HTTP状态码

	ExpireIn    int    `json:"expireIn"`
	AccessToken string `json:"accessToken"`

	createTime time.Time
}

// UserAccessToken
// 文档: https://open.dingtalk.com/document/orgapp/obtain-user-token
type UserAccessToken struct {
	*BasicResponse // 此接口未返回错误代码信息，仅仅能检查HTTP状态码

	ExpireIn     int    `json:"expireIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	CorpId       string `json:"corpId"`
}

func (a AccessToken) GetToken() string {
	return a.AccessToken
}

func (a AccessToken) GetExpireIn() time.Duration {
	return time.Duration(a.ExpireIn) * time.Second
}

func (a AccessToken) GetExpireTime() time.Time {
	return a.createTime.Add(a.GetExpireIn())
}
