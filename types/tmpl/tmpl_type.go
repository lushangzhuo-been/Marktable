package tmpl

import (
	"mark3/types/common"
)

type TmplConfigResp struct {
	ModeConfig []common.BaseConfig `json:"mode_config"`
}

type TmplListReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}

type WsFileTmplListReq struct {
	UserId   int `form:"user_id" json:"user_id" binding:"required"`
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	WsFileId int `form:"ws_file_id" json:"ws_file_id" binding:"required"`
}

type TmplInfoReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
	Id   int `form:"id" json:"id" binding:"required"`
}

type TmplCreateReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Desc     string `form:"desc" json:"desc" binding:"required"`
	Mode     string `form:"mode" json:"mode" binding:"required"`
	WsFileId int    `form:"ws_file_id" json:"ws_file_id"`
}

type TmplUpdateReq struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	WsId int    `form:"ws_id" json:"ws_id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc" binding:"required"`
	Mode string `form:"mode" json:"mode" binding:"required"`
}

type TmplDeleteReq struct {
	Id   int `form:"id" json:"id" binding:"required"`
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}
