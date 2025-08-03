package tmpl

import (
	"mark3/types/common"
)

type FieldConfigResp struct {
	ModeConfig      []common.BaseConfig `json:"mode_config"`
	NumberComConfig []common.BaseConfig `json:"number_com_config"`
	TimeComConfig   []common.BaseConfig `json:"time_com_config"`
	PersonComConfig []common.BaseConfig `json:"person_com_config"`
}

type FieldListReq struct {
	WsId      int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId    int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	HasStatus string `form:"has_status" json:"has_status"`
}

type FieldInfoReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type FieldEnumValue struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type FieldNumberComExpr struct {
}

type FieldTimeComExpr struct {
}

type FieldPersonComExpr struct {
}

type FieldProgressComExpr struct {
}

type FieldCreateReq struct {
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	Desc          string `form:"desc" json:"desc"`
	Mode          string `form:"mode" json:"mode" binding:"required"`
	EnumValues    string `form:"enum_values" json:"enum_values"`
	Expr          string `form:"expr" json:"expr"`
	RelatedTmplId int    `form:"related_tmpl_id" json:"related_tmpl_id"`
}

type FieldUpdateReq struct {
	Id            int    `form:"id" json:"id" binding:"required"`
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	Desc          string `form:"desc" json:"desc"`
	EnumValues    string `form:"enum_values" json:"enum_values"`
	Expr          string `form:"expr" json:"expr"`
	RelatedTmplId int    `form:"related_tmpl_id" json:"related_tmpl_id"`
}

type FieldGetReadOnlyRuleReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type FieldUpdateReadOnlyRuleReq struct {
	Id            int    `form:"id" json:"id" binding:"required"`
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StatusList    string `form:"status_list" json:"status_list"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type FieldGetAllPersonComReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type FieldDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}
