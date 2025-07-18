package rule

import (
	"go.mongodb.org/mongo-driver/bson"
	model "mark3/repository/db/model/rule"
	"mark3/repository/db/model/user"
	"mark3/types/common"
)

type Action struct {
	Id            int    `form:"id" json:"id"`
	RuleId        int    `form:"rule_id" json:"rule_id"`
	ActionType    string `form:"action_type" json:"action_type"`
	FieldKey      string `form:"field_key" json:"field_key"`
	FieldValue    string `form:"field_value" json:"field_value"`
	EmailTitle    string `form:"email_title" json:"email_title"`
	EmailContents string `form:"email_contents" json:"email_contents"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
	IsDeleted     string `form:"is_deleted" json:"is_deleted"`
}

type CreateReq struct {
	WsId        int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId      int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	RuleType    string `form:"rule_type" json:"rule_type" binding:"required"`
	IsEnable    string `form:"is_enable" json:"is_enable"`
	IsDeleted   string `form:"is_deleted" json:"is_deleted"`
	Description string `form:"description" json:"description"`
	Filter      string `form:"filter" json:"filter"`
	TriggerDay  int    `form:"trigger_day" json:"trigger_day"`
	TriggerTime string `form:"trigger_time" json:"trigger_time"`
	FieldKey    string `form:"field_key" json:"field_key"`
	OldValue    string `form:"old_value" json:"old_value"`
	NewValue    string `form:"new_value" json:"new_value"`
	ActionList  string `form:"action_list" json:"action_list"`
}

type UpdateReq struct {
	CreateReq
	Id int `form:"id" json:"id" binding:"required"`
}

type DeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type DetailReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type ActionListVo struct {
	model.RuleActionModel
	FieldValueUser      []user.UserModel `json:"field_value_user"`
	FieldValueDocuments []bson.M         `json:"field_value_documents"`
	UserListInfo        []user.UserModel `json:"user_list_info"`
}

type DetailVo struct {
	model.RuleModel
	OldValueUser      []user.UserModel `json:"old_value_user"`
	NewValueUser      []user.UserModel `json:"new_value_user"`
	OldValueDocuments []bson.M         `json:"old_value_documents"`
	NewValueDocuments []bson.M         `json:"new_value_documents"`
	Creator           []user.UserModel `json:"creator"`
	ActionList        []ActionListVo   `json:"action_list"`
	FormatFilter      string           `json:"format_filter"`
	TmplName          string           `json:"tmpl_name"`
}

type PageReq struct {
	WsId            int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId          int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	RuleTypeSelect  string `form:"rule_type_select" json:"rule_type_select"`
	IsEnableSelect  string `form:"is_enable_select" json:"is_enable_select"`
	CreatedBySelect string `form:"created_by_select" json:"created_by_select"`
	common.BasePageReq
}

type RulePageVo struct {
	model.RuleModel
	Creator []user.UserModel `json:"creator"`
}

type EnableReq struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	IsEnable string `form:"is_enable" json:"is_enable" binding:"required"`
}

type RuleLogReq struct {
	common.BasePageReq
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	RuleId int `form:"rule_id" json:"rule_id"`
}

type ActionLogReq struct {
	WsId            int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId          int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	RuleExecuteUuid string `form:"rule_execute_uuid" json:"rule_execute_uuid" binding:"required"`
}

type ActionLogListVo struct {
	model.RuleActionLogModel
	EmailUser           []user.UserModel `json:"email_user"`
	FieldValueUser      []user.UserModel `json:"field_value_user"`
	FieldValueDocuments []bson.M         `json:"field_value_documents"`
}

type ConfigReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type ConfigResp struct {
	RuleType          []common.BaseConfig `json:"rule_type"`
	RuleEnable        []common.BaseConfig `json:"rule_enable"`
	RuleDeleted       []common.BaseConfig `json:"rule_deleted"`
	RuleActionType    []common.BaseConfig `json:"rule_action_type"`
	RuleActionDeleted []common.BaseConfig `json:"rule_action_deleted"`
	ActionLogStatus   []common.BaseConfig `json:"action_log_status"`
}

type RuleActionLogResp struct {
	RuleDetail interface{} `json:"rule_detail"`
	List       interface{} `json:"list"`
	Cnt        int64       `json:"cnt"`
}
