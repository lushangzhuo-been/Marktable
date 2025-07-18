package ws

import "mark3/types/common"

const (
	WsStatusOk     = "ok"
	WsStatusForbid = "forbid"
)

var WsStatusMap = map[string]string{
	WsStatusOk:     "开启",
	WsStatusForbid: "禁用",
}

var WsStatusConfig = []common.BaseConfig{
	{
		Value: "ok",
		Label: "开启",
	},
	{
		Value: "forbid",
		Label: "禁用",
	},
}

const (
	WsRoleAdmin   = "admin"
	WsRoleGeneral = "general"
)

var WsRoleMap = map[string]string{
	WsRoleAdmin:   "管理员",
	WsRoleGeneral: "普通成员",
}

var WsRoleConfig = []common.BaseConfig{
	{
		Value: "admin",
		Label: "管理员",
	},
	{
		Value: "general",
		Label: "普通成员",
	},
}
