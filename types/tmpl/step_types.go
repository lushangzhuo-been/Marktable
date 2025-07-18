package tmpl

type StepInfoReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}

type StepCreateReq struct {
	WsId          int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StartStatusId int `form:"start_status_id" json:"start_status_id" binding:"required"`
	EndStatusId   int `form:"end_status_id" json:"end_status_id" binding:"required"`
}

type StepDeleteReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	Id     int `form:"id" json:"id" binding:"required"`
}
