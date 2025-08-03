package ws

import (
	model "mark3/repository/db/model/ws"
)

type WsInfoReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type WsVo struct {
	model.WsModel
	CreatorName string     `json:"creator_name"`
	Admin       []MemberVo `json:"admin"`
}

type WsCreateReq struct {
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc" binding:"required"`
}

type WsUpdateReq struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc" binding:"required"`
}

type WsDeleteReq struct {
	Id int `form:"id" json:"id" binding:"required"`
}
