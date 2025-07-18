package tmpl

import (
	model "mark3/repository/db/model/tmpl"
	types2 "mark3/types/common"
)

type NodeConfigResp struct {
	ModeConfig []types2.BaseConfig `json:"mode_config"`
}

type NodeVo struct {
	Node    model.NodeModel `json:"node"`
	Buttons []ButtonVo      `json:"buttons"`
}

type NodeListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type NodeInfoReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type NodeCreateReq struct {
	WsId        int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId      int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Mode        string `form:"mode" json:"mode" binding:"required"`
	StatusText  string `form:"status_text" json:"status_text" binding:"required"`
	StatusColor string `form:"status_color" json:"status_color"`
}

type NodeUpdateReq struct {
	Id          int    `form:"id" json:"id" binding:"required"`
	WsId        int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId      int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Mode        string `form:"mode" json:"mode" binding:"required"`
	StatusText  string `form:"status_text" json:"status_text" binding:"required"`
	StatusColor string `form:"status_color" json:"status_color"`
}

type NodeDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type NodeGetAllStatusesReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
