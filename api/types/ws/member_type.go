package ws

import (
	"mark3/types/common"
)

type MemberConfigResp struct {
	StatusConfig []common.BaseConfig `json:"status_config"`
	RoleConfig   []common.BaseConfig `json:"role_config"`
}

type MemberListReq struct {
	WsId int    `form:"ws_id" json:"ws_id" binding:"required"`
	Key  string `form:"key" json:"key"`
	common.BasePageReq
}

type MemberVo struct {
	Id       int    `json:"id"`
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
	Status   string `json:"status"`
}

type MemberCreateReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	UseridList string `form:"userid_list" json:"userid" binding:"required"`
	Role       string `form:"role" json:"role" binding:"required"`
	Status     string `form:"status" json:"status" binding:"required"`
}

type MemberUpdateReq struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	Role   string `form:"role" json:"role" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
}

type MemberDeleteReq struct {
	Id   int `form:"id" json:"id" binding:"required"`
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}
