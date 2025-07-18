package tmpl

import (
	"mark3/types/common"
)

type StepScreenConfigResp struct {
	RequiredConfig []common.BaseConfig `json:"required_config"`
}

type StepScreenSurplusReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}

type StepScreenListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}

type StepScreenVo struct {
	Id        int    `json:"id"`
	FieldKey  string `json:"field_key"`
	FieldName string `json:"field_name"`
	Required  string `json:"required"`
}

type StepScreenCreateReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId   int    `form:"step_id" json:"step_id" binding:"required"`
	FieldKey string `form:"field_key" json:"field_key" binding:"required"`
}

type StepScreenSetRequiredReq struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId   int    `form:"step_id" json:"step_id" binding:"required"`
	Required string `form:"required" json:"required" binding:"required"`
}

type StepScreenMoveReq struct {
	WsId       int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId     int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId     int    `form:"step_id" json:"step_id" binding:"required"`
	Coordinate string `form:"coordinate" json:"coordinate" binding:"required"`
}

type StepScreenDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}
