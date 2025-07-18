package message

import commonTypes "mark3/types/common"

type GetMessageListReq struct {
	commonTypes.BasePageReq
	Mode string `form:"mode" json:"mode" binding:"required"`
}

type ReadAllMessageReq struct {
	Mode string `form:"mode" json:"mode" binding:"required"`
}

type ReadMessageReq struct {
	Id string `form:"id" json:"id" binding:"required"`
}

type GetMessageListResp struct {
	List      interface{} `json:"list"`
	Cnt       int64       `json:"cnt"`
	UnReadCnt int64       `json:"un_read_cnt"`
}
