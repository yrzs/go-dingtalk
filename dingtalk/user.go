package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yrzs/go-dingtalk/core"
	"github.com/yrzs/go-dingtalk/response"
	"net/http"
)

// getUserInfo 根据用户个人token
// https://open.dingtalk.com/document/orgapp/dingtalk-retrieve-user-information?spm=ding_open_doc.document.0.0.4a1477a27izOMd
func (c *Client) getUserInfo(ctx context.Context, userToken string, unionId string) (*response.UserInfo, error) {
	var user *response.UserInfo

	endpoint := fmt.Sprintf("https://api.dingtalk.com/v1.0/contact/users/%s", unionId)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	req.Header.Set("x-acs-dingtalk-access-token", userToken)
	req.Header.Set("Content-Type", "application/json")

	if err := core.Request(req, &user); err != nil {
		return user, err
	}
	return user, nil
}

// GetUserinfoByCode 通过code获取用户信息
// https://open.dingtalk.com/document/orgapp/dingtalk-retrieve-user-information?spm=ding_open_doc.document.0.0.4a1477a27izOMd
func (c *Client) GetUserinfoByCode(ctx context.Context, code string) (*response.UserInfo, error) {
	userTokenResp, err := c.getUserAccessToken(ctx, code)
	if err != nil {
		return nil, err
	}
	token := userTokenResp.AccessToken
	userInfo, err := c.getUserInfo(ctx, token, "me")
	if err != nil {
		return nil, err
	}
	userIdResp, err := c.getUserIdByUnionId(ctx, userInfo.UnionId)
	if err == nil && userIdResp.Result.Userid != "" {
		userInfo.UserId = userIdResp.Result.Userid
	}
	return userInfo, nil
}

// GetUserIdByUnionId 根据 unionId 获取用户userId
// https://open.dingtalk.com/document/isvapp/query-a-user-by-the-union-id
func (c *Client) getUserIdByUnionId(ctx context.Context, unionId string) (*response.GetUserIdResp, error) {
	token, _ := c.getAccessToken(ctx)
	endpoint := fmt.Sprintf("https://oapi.dingtalk.com/topapi/user/getbyunionid?access_token=%s", token.AccessToken)
	data := map[string]string{
		"unionid": unionId,
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	var resp *response.GetUserIdResp
	if err := core.Request(req, &resp); err != nil {
		return resp, err
	}
	return resp, nil

}
