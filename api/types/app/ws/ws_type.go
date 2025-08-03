package ws

import (
	model "mark3/repository/db/model/ws"
	"mark3/types/app"
)

type WsInfoReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}

type WsInfoResp struct {
	Ws        model.WsModel  `json:"ws"`
	AdminList []app.MemberVo `json:"admin_list"`
	TmplCnt   int64          `json:"tmpl_cnt"`
	MemberCnt int64          `json:"member_cnt"`
	IsMember  string         `json:"is_member"`
	IsAdmin   string         `json:"is_admin"`
}
