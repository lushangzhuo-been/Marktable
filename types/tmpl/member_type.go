package tmpl

import (
	"mark3/types/common"
)

type MemberConfigResp struct {
	StatusConfig []common.BaseConfig `json:"status_config"`
}

type MemberListReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Key    string `form:"key" json:"key"`
	common.BasePageReq
}

type MemberVo struct {
	Id       int    `json:"id"`
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	Avatar   string `json:"avatar"`
	Status   string `json:"status"`
}

type MemberCreateReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId     int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	UserIdList string `form:"userid_list" json:"userid_list" binding:"required"`
	RoleId     int    `form:"role_id" json:"role_id" binding:"required"`
	Status     string `form:"status" json:"status" binding:"required"`
}

type MemberUpdateReq struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	RoleId int    `form:"role_id" json:"role_id" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
}

type MemberDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
