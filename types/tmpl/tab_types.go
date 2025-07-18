package tmpl

import "mark3/types/common"

const (
	TabTmplDetail string = "tmpl_detail"
	TabFile       string = "file"
	TabSubTmpl    string = "sub_tmpl"
)

var TabMap = map[string]string{
	TabTmplDetail: "流程详情",
	TabFile:       "附件",
	TabSubTmpl:    "子任务",
}

var TabConfig = []common.BaseConfig{
	{
		Value: "tmpl_detail",
		Label: "流程详情",
	},
	{
		Value: "file",
		Label: "附件",
	},
	{
		Value: "sub_tmpl",
		Label: "子任务",
	},
}

type TmplTabListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type TmplTabCreateReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Tab    string `form:"tab" json:"tab" binding:"required"`
}

type TmplTabDeleteReq struct {
	TmplTabCreateReq
}

type TmplTabConfigReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
