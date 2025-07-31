package tmpl

type TmplSubConfigCheckReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type TmplSubConfigListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
}

type TmplSubConfigCreateReq struct {
	WsId      int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId    int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	SubTmplId int `form:"sub_tmpl_id" json:"sub_tmpl_id" binding:"required"`
}

type TmplSubConfigDeleteReq struct {
	TmplSubConfigCreateReq
}
