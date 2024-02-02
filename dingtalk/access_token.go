package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/yrzs/go-dingtalk/core"
	"github.com/yrzs/go-dingtalk/response"
	"net/http"
)

// GetAccessToken 获取企业内部应用access token
// https://open.dingtalk.com/document/orgapp/obtain-the-access_token-of-an-internal-app?spm=ding_open_doc.document.0.0.184b3f20WYVl0p
func (c *Client) getAccessToken(ctx context.Context) (*response.AccessToken, error) {
	data := map[string]string{
		"appKey":    c.conf.ClientId,
		"appSecret": c.conf.ClientSecret,
	}

	jsonData, _ := json.Marshal(data)
	endpoint := "https://api.dingtalk.com/v1.0/oauth2/accessToken"
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	var token *response.AccessToken
	if err := core.Request(req, &token); err != nil {
		return token, err
	}

	return token, nil
}

// getUserAccessToken 获取用户access token
// https://open.dingtalk.com/document/orgapp/obtain-user-token?spm=ding_open_doc.document.0.0.4a1477a27izOMd
func (c *Client) getUserAccessToken(ctx context.Context, code string) (*response.UserAccessToken, error) {
	val := map[string]string{
		"clientId":     c.conf.ClientId,
		"clientSecret": c.conf.ClientSecret,
		"code":         code,
		"grantType":    "authorization_code",
	}
	jv, _ := json.Marshal(val)
	endpoint := "https://api.dingtalk.com/v1.0/oauth2/userAccessToken"
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jv))
	req.Header.Set("Content-Type", "application/json")

	var token *response.UserAccessToken
	if err := core.Request(req, &token); err != nil {
		return token, err
	}

	return token, nil
}
