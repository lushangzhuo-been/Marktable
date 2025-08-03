package ws

import "mark3/types/common"

type GetUserListReq struct {
	Key string `form:"key" json:"key"`
	common.BasePageReq
}
