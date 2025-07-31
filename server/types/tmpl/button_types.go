package tmpl

import model "mark3/repository/db/model/tmpl"

type ButtonVo struct {
	model.ButtonModel
	TargetNodeName string `json:"target_node_name"`
}

type ButtonInfoReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId int `form:"node_id" json:"node_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type ButtonCreateReq struct {
	WsId         int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId       int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId       int    `form:"node_id" json:"node_id" binding:"required"`
	Name         string `form:"name" json:"name" binding:"required"`
	TargetNodeId int    `form:"target_node_id" json:"target_node_id" binding:"required"`
}

type ButtonUpdateReq struct {
	Id           int    `form:"id" json:"id" binding:"required"`
	WsId         int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId       int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId       int    `form:"node_id" json:"node_id" binding:"required"`
	Name         string `form:"name" json:"name" binding:"required"`
	TargetNodeId int    `form:"target_node_id" json:"target_node_id" binding:"required"`
}

type ButtonDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId int `form:"node_id" json:"node_id" binding:"required"`
}
