package ws_file

type WsFileInfoReq struct {
	WsId int `form:"ws_id" json:"ws_id" binding:"required"`
	Id   int `form:"id" json:"id" binding:"required"`
}
