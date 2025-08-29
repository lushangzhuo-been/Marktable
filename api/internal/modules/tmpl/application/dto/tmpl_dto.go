package dto

type TmplCreateReq struct {
	WsId     int    `form:"ws_id" json:"ws_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Desc     string `form:"desc" json:"desc" binding:"required"`
	Mode     string `form:"mode" json:"mode" binding:"required"`
	WsFileId int    `form:"ws_file_id" json:"ws_file_id"`
}
