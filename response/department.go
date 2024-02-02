package response

type DepartmentList struct {
	// https://open.dingtalk.com/document/orgapp/obtain-the-department-list-v2
	*BasicResponse

	Result []struct {
		AutoAddUser     bool   `json:"auto_add_user"`
		CreateDeptGroup bool   `json:"create_dept_group"`
		DeptID          int    `json:"dept_id"`
		Name            string `json:"name"`
		ParentID        int    `json:"parent_id"`
	} `json:"result"`
}

type Department struct {
	// https://open.dingtalk.com/document/isvapp/query-department-details0-v2
	*BasicResponse

	Result struct {
		Brief                 string   `json:"brief"`
		DeptPermits           []any    `json:"dept_permits"`
		OuterPermitUsers      []any    `json:"outer_permit_users"`
		EmpApplyJoinDept      bool     `json:"emp_apply_join_dept"`
		OrgDeptOwner          string   `json:"org_dept_owner"`
		OuterDept             bool     `json:"outer_dept"`
		AutoApproveApply      bool     `json:"auto_approve_apply"`
		DeptGroupChatID       string   `json:"dept_group_chat_id"`
		GroupContainSubDept   bool     `json:"group_contain_sub_dept"`
		AutoAddUser           bool     `json:"auto_add_user"`
		DeptManagerUseridList []string `json:"dept_manager_userid_list"`
		ParentID              int      `json:"parent_id"`
		HideDept              bool     `json:"hide_dept"`
		Name                  string   `json:"name"`
		OuterPermitDepts      []any    `json:"outer_permit_depts"`
		UserPermits           []any    `json:"user_permits"`
		DeptID                int      `json:"dept_id"`
		CreateDeptGroup       bool     `json:"create_dept_group"`
		Order                 int      `json:"order"`
	} `json:"result"`
}

type DepartmentIdsResp struct {
	*BasicResponse

	Result struct {
		ParentList []struct {
			ParentDeptIDList []int `json:"parent_dept_id_list"`
		} `json:"parent_list"`
	} `json:"result"`
}
