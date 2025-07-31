package tmpl

import "mark3/types/common"

type GetUserListReq struct {
	WsId int    `form:"ws_id" json:"ws_id"`
	Key  string `form:"key" json:"key"`
	common.BasePageReq
}
