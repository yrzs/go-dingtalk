package response

type BasicResponse struct {
	Message string `json:"errmsg"`
	Code    int    `json:"errcode"`
}
