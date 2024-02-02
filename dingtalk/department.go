package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/yrzs/go-dingtalk/core"
	"github.com/yrzs/go-dingtalk/response"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
)

// ListDepartment 部门列表
func (c *Client) ListDepartment(ctx context.Context, parentDeptId string) (*response.DepartmentList, error) {
	accessToken, _ := c.getAccessToken(ctx)
	data := map[string]string{
		"dept_id":  parentDeptId,
		"language": "zh_CN",
	}

	endpoint := "https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=" + accessToken.AccessToken

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	var resp *response.DepartmentList
	if err := core.Request(req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}

type DepartmentNode struct {
	AutoAddUser     bool              `json:"auto_add_user"`
	CreateDeptGroup bool              `json:"create_dept_group"`
	DeptID          int               `json:"dept_id"`
	Name            string            `json:"name"`
	ParentID        int               `json:"parent_id"`
	Children        []*DepartmentNode `json:"children"` // 子部门列表
}

// ListDepartmentTree 获取部门树 会被qps限制
func (c *Client) ListDepartmentTree(ctx context.Context, parentDeptId string) (*DepartmentNode, error) {
	if parentDeptId == "" {
		return nil, nil
	}
	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	token := accessToken.AccessToken
	endpoint := "https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=" + token
	data := map[string]string{
		"dept_id":  parentDeptId,
		"language": "zh_CN",
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	var resp *response.DepartmentList
	if err := core.Request(req, &resp); err != nil {
		return nil, err
	}
	deptIdInt, err := strconv.Atoi(parentDeptId)
	if err != nil {
		return nil, err
	}
	currentNode := &DepartmentNode{
		DeptID:   deptIdInt,
		Children: make([]*DepartmentNode, 0),
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, dept := range resp.Result {
		dept := dept // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			childNode, err := c.ListDepartmentTree(ctx, strconv.Itoa(dept.DeptID))
			if err != nil {
				return err
			}
			childNode.DeptID = dept.DeptID
			childNode.Name = dept.Name
			childNode.ParentID = dept.ParentID
			childNode.AutoAddUser = dept.AutoAddUser
			childNode.CreateDeptGroup = dept.CreateDeptGroup
			// 需要同步访问currentNode.Children
			// 可以考虑使用互斥锁或其他同步机制
			currentNode.Children = append(currentNode.Children, childNode)
			return nil
		})
	}

	// 等待所有goroutine完成
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return currentNode, nil
}

// GetDepartment 获取部门详情
// https://open.dingtalk.com/document/isvapp/query-department-details0-v2
func (c *Client) GetDepartment(ctx context.Context, deptId string) (*response.Department, error) {
	accessToken, _ := c.getAccessToken(ctx)
	data := map[string]string{
		"dept_id": deptId,
	}

	endpoint := "https://oapi.dingtalk.com/topapi/v2/department/get?access_token=" + accessToken.AccessToken

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	var resp *response.Department
	if err := core.Request(req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}

// ListDepartmentByUserId 通过userId 获取用户的部门列表
func (c *Client) ListDepartmentByUserId(ctx context.Context, userId string) (*response.DepartmentIdsResp, error) {
	accessToken, _ := c.getAccessToken(ctx)
	data := map[string]string{
		"userid": userId,
	}
	endpoint := "https://oapi.dingtalk.com/topapi/v2/department/listparentbyuser?access_token=" + accessToken.AccessToken
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	var resp *response.DepartmentIdsResp
	if err := core.Request(req, &resp); err != nil {
		return resp, err
	}

	return resp, nil

}
