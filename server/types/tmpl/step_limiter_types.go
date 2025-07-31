package tmpl

type StepLimiterListReq struct {
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}

type StepLimiterInfoReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}

type StepLimiterCreateReq struct {
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId        int    `form:"step_id" json:"step_id" binding:"required"`
	Mode          string `form:"mode" json:"mode" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type StepLimiterUpdateReq struct {
	Id            int    `form:"id" json:"id" binding:"required"`
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId        int    `form:"step_id" json:"step_id" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type StepLimiterDeleteReq struct {
	Id     int `form:"id" json:"id" binding:"required"`
	WsId   int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	StepId int `form:"step_id" json:"step_id" binding:"required"`
}
