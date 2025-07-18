package tmpl

import (
	"mark3/types/common"
)

type ButtonScreenConfigResp struct {
	RequiredConfig []common.BaseConfig `json:"required_config"`
}

type ButtonScreenSurplusReq struct {
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}

type ButtonScreenListReq struct {
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}

type ButtonScreenVo struct {
	Id        int    `json:"id"`
	FieldKey  string `json:"field_key"`
	FieldName string `json:"field_name"`
	Required  string `json:"required"`
}

type ButtonScreenCreateReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int    `form:"button_id" json:"button_id" binding:"required"`
	FieldKey string `form:"field_key" json:"field_key" binding:"required"`
}

type ButtonScreenSetRequiredReq struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int    `form:"button_id" json:"button_id" binding:"required"`
	Required string `form:"required" json:"required" binding:"required"`
}

type ButtonScreenMoveReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId     int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId     int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId   int    `form:"button_id" json:"button_id" binding:"required"`
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type ButtonScreenDeleteReq struct {
	Id       int `form:"id" json:"id" binding:"required"`
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}
