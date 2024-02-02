package response

type UserInfo struct {
	// https://open.dingtalk.com/document/orgapp/dingtalk-retrieve-user-information
	*BasicResponse

	NickName  string `json:"nick"`
	Avatar    string `json:"avatarUrl"`
	Mobile    string `json:"mobile"`
	OpenId    string `json:"openId"`
	UnionId   string `json:"unionId"`
	Email     string `json:"email"`
	StateCode string `json:"stateCode"`
	UserId    string
}

type GetUserIdResp struct {
	// https://open.dingtalk.com/document/isvapp/query-a-user-by-the-union-id
	*BasicResponse

	Result struct {
		ContactType int    `json:"contact_type"`
		Userid      string `json:"userid"`
	} `json:"result"`
}
