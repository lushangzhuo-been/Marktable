package tmpl

import (
	commonTypes "mark3/types/common"
)

type CommonReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type GetTmplListReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
}

type GetTmplOtherListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type GetTmplInfoReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type GetConfigReq struct {
	CommonReq
}

type GetStatusListReq struct {
	CommonReq
}

type GetListModuleScreenConfigReq struct {
	CommonReq
}

type SetListModuleScreenConfigReq struct {
	CommonReq
	Screen string `form:"screen" json:"screen" binding:"required"`
}

type GetScreenReq struct {
	CommonReq
	Module  string `form:"module" json:"module" binding:"required"`
	Columns string `form:"columns" json:"columns"`
}

type GetListDataReq struct {
	CommonReq
	commonTypes.BasePageReq
	Filter       string `form:"filter" json:"filter"`
	FilterDown   string `form:"filter_down" json:"filter_down"`
	BoardId      int    `form:"board_id" json:"board_id"`
	Lor          string `form:"lor" json:"lor"`
	SortOrder    string `form:"sort_order" json:"sort_order"`
	Export       string `form:"export" json:"export"`
	ParentTmplId int    `form:"parent_tmpl_id" json:"parent_tmpl_id"`
	ObjId        string `form:"obj_id" json:"obj_id"`
}

type GetListDataSelectReq struct {
	CommonReq
	commonTypes.BasePageReq
	Key string `form:"key" json:"key"`
	Ex  string `form:"ex" json:"ex"`
}

type GetUserAuthReq struct {
	CommonReq
	Id       string `form:"id" json:"id" binding:"required"`
	AuthMode string `form:"auth_mode" json:"auth_mode" binding:"required"`
}

type GetDataReq struct {
	CommonReq
	Id string `form:"id" json:"id" binding:"required"`
}

type UpdateReq struct {
	CommonReq
	Id string `form:"id" json:"id" binding:"required"`
}

type GetStepListReq struct {
	CommonReq
	Id string `form:"id" json:"id" binding:"required"`
}

type GetStepScreenReq struct {
	CommonReq
	StepId int    `form:"step_id" json:"step_id" binding:"required"`
	Id     string `form:"id" json:"id" binding:"required"`
}

type SwitchStepReq struct {
	CommonReq
	StepId int    `form:"step_id" json:"step_id" binding:"required"`
	Id     string `form:"id" json:"id" binding:"required"`
}

type GetButtonScreenReq struct {
	CommonReq
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}

type ClickButtonActionReq struct {
	CommonReq
	NodeId   int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int    `form:"button_id" json:"button_id" binding:"required"`
	Id       string `form:"id" json:"id" binding:"required"`
}

type DeleteReq struct {
	CommonReq
	Ids string `form:"ids" json:"ids" binding:"required"`
}

type ScreenVo struct {
	FieldKey      string        `json:"field_key"`
	Name          string        `json:"name"`
	Desc          string        `json:"desc"`
	Mode          string        `json:"mode"`
	EnumValues    []interface{} `json:"enum_values"`
	Expr          string        `json:"expr"`
	CanModify     string        `json:"can_modify"`
	Required      string        `json:"required"`
	Level         string        `json:"level"`
	ReadOnlyRule  string        `json:"read_only_rule"`
	RelatedTmplId int           `json:"related_tmpl_id"`
}

type AddProgress struct {
	CommonReq
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type UpdateProgress struct {
	CommonReq
	Id      string `form:"id" json:"id" binding:"required"`
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type DeleteProgress struct {
	CommonReq
	Id      string `form:"id" json:"id" binding:"required"`
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
}

type GetProgressList struct {
	CommonReq
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
	commonTypes.BasePageReq
}

type GetLogList struct {
	CommonReq
	IssueId string `form:"issue_id" json:"issue_id" binding:"required"`
	commonTypes.BasePageReq
}

type GetMemberList struct {
	CommonReq
}

type CreateSubActionReq struct {
	CommonReq
	ParentTmplId int    `form:"parent_tmpl_id" json:"parent_tmpl_id" binding:"required"`
	ObjId        string `form:"obj_id" json:"obj_id" binding:"required"`
}

type DeleteSubActionReq struct {
	CommonReq
	ParentTmplId int    `form:"parent_tmpl_id" json:"parent_tmpl_id" binding:"required"`
	ObjId        string `form:"obj_id" json:"obj_id" binding:"required"`
	Ids          string `form:"ids" json:"ids" binding:"required"`
}

type GetSubListCountReq struct {
	CommonReq
	ObjId string `form:"obj_id" json:"obj_id" binding:"required"`
}
