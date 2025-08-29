package dto

type WsCreateReq struct {
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `form:"desc" json:"desc" binding:"required"`
}
