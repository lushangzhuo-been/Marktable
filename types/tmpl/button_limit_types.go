package tmpl

type ButtonLimitListReq struct {
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}

type ButtonLimitInfoReq struct {
	Id       int `form:"id" json:"id" binding:"required"`
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}

type ButtonLimitCreateReq struct {
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId        int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId      int    `form:"button_id" json:"button_id" binding:"required"`
	Mode          string `form:"mode" json:"mode" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type ButtonLimitUpdateReq struct {
	Id            int    `form:"id" json:"id" binding:"required"`
	WsId          int    `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId        int    `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId        int    `form:"node_id" json:"node_id" binding:"required"`
	ButtonId      int    `form:"button_id" json:"button_id" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required"`
	UserList      string `form:"user_list" json:"user_list"`
	IssueRoleList string `form:"issue_role_list" json:"issue_role_list"`
}

type ButtonLimitDeleteReq struct {
	Id       int `form:"id" json:"id" binding:"required"`
	WsId     int `form:"ws_id" json:"ws_id" binding:"required"`
	TmplId   int `form:"tmpl_id" json:"tmpl_id" binding:"required"`
	NodeId   int `form:"node_id" json:"node_id" binding:"required"`
	ButtonId int `form:"button_id" json:"button_id" binding:"required"`
}
