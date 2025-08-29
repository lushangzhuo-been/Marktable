package tmpl

import "mark3/types/common"

type AuthConfigReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type AuthConfigResp struct {
	RoleConfig []common.BaseConfig `json:"role_config"`
	AuthConfig []common.BaseConfig `json:"auth_config"`
}

type AuthUpdateReq struct {
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Mode          string `form:"mode" json:"mode" binding:"required"`
	WsRoles       string `form:"ws_roles" json:"ws_roles"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type AuthDetailReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Mode   string `form:"mode" json:"mode" binding:"required"`
}
