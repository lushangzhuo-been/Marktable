package app

import (
	"mark3/types/common"
)

type CommonGetWsUserListReq struct {
	WsId int    `form:"ws_id" json:"ws_id" binding:"required"`
	Key  string `form:"key" json:"key"`
	Ex   string `form:"ex" json:"ex"`
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
