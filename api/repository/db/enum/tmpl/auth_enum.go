package tmpl

import model "mark3/repository/db/model/tmpl"

const (
	AuthSee      = "see"
	AuthExport   = "export"
	AuthCreate   = "create"
	AuthEdit     = "edit"
	AuthDelete   = "delete"
	AuthProgress = "progress"
	AuthFile     = "file"
	AuthSubTmpl  = "sub_tmpl"
)

var AuthMap = map[string]string{
	AuthSee:      "查看",
	AuthExport:   "导出",
	AuthCreate:   "创建",
	AuthEdit:     "编辑",
	AuthDelete:   "删除",
	AuthProgress: "进展",
	AuthFile:     "附件",
	AuthSubTmpl:  "子任务",
}

var AuthList = []model.TmplAuthModel{
	{
		Mode:          AuthSee,
		WsRoles:       "admin,general",
		UserList:      "",
		IssueRoleList: "",
	},
	{
		Mode:          AuthExport,
		WsRoles:       "admin,general",
		UserList:      "",
		IssueRoleList: "",
	},
	{
		Mode:          AuthCreate,
		WsRoles:       "admin,general",
		UserList:      "",
		IssueRoleList: "",
	},
	{
		Mode:          AuthEdit,
		WsRoles:       "",
		UserList:      "",
		IssueRoleList: "handler,creator",
	},
	{
		Mode:          AuthDelete,
		WsRoles:       "",
		UserList:      "",
		IssueRoleList: "creator",
	},
	{
		Mode:          AuthProgress,
		WsRoles:       "",
		UserList:      "",
		IssueRoleList: "handler,creator",
	},
	{
		Mode:          AuthFile,
		WsRoles:       "",
		UserList:      "",
		IssueRoleList: "handler,creator",
	},
	{
		Mode:          AuthSubTmpl,
		WsRoles:       "",
		UserList:      "",
		IssueRoleList: "handler,creator",
	},
}
