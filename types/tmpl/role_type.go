package tmpl

import (
	"mark3/types/common"
)

type RoleConfigResp struct {
	ViewPermission []common.BaseConfig `json:"view_permission_config"`
	EditPermission []common.BaseConfig `json:"edit_permission_config"`
}

type RoleListReq struct {
	WsId   int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Key    string `form:"key" json:"key"`
	common.BasePageReq
}

type RoleCreateReq struct {
	WsId           int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId         int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required"`
	Desc           string `form:"desc" json:"desc" binding:"required"`
	ViewPermission string `form:"view_permission" json:"view_permission" binding:"required"`
	EditPermission string `form:"edit_permission" json:"edit_permission" binding:"required"`
}

type RoleUpdateReq struct {
	Id             int    `form:"id" json:"id" binding:"required"`
	WsId           int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId         int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required"`
	Desc           string `form:"desc" json:"desc" binding:"required"`
	ViewPermission string `form:"view_permission" json:"view_permission" binding:"required"`
	EditPermission string `form:"edit_permission" json:"edit_permission" binding:"required"`
}

type RoleDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
