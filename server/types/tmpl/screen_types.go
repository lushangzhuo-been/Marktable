package tmpl

import (
	"mark3/types/common"
)

type ScreenConfigResp struct {
	ModuleConfig   []common.BaseConfig `json:"module_config"`
	RequiredConfig []common.BaseConfig `json:"required_config"`
}

type ScreenSurplusReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Module string `form:"module" json:"module" binding:"required"`
}

type ScreenListReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Module string `form:"module" json:"module" binding:"required"`
}

type ScreenVo struct {
	Id        int    `json:"id"`
	Module    string `json:"module"`
	FieldKey  string `json:"field_key"`
	FieldName string `json:"field_name"`
	CanModify string `json:"can_modify"`
	Required  string `json:"required"`
}

type ScreenCreateReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	FieldKey string `form:"field_key" json:"field_key" binding:"required"`
	Module   string `form:"module" json:"module" binding:"required"`
}

type ScreenSetReq struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Module   string `form:"module" json:"module" binding:"required"`
	Required string `form:"required" json:"required" binding:"required"`
}

type ScreenMoveReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId     int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Module     string `form:"module" json:"module" binding:"required"`
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type ScreenDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
