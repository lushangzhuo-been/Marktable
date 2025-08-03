package tmpl

import "mark3/types/common"

const (
	RoleSignAdmin   string = "admin"
	RoleSignGeneral string = "general"
)

const (
	//查看所有条目
	RoleViewPermissionAll string = "view_all"
	//仅能查看与自己相关的条目（创建,指派）
	RoleViewPermissionOnlyMe string = "view_only_me"
)

var RoleViewPermissionMap = map[string]string{
	RoleViewPermissionAll:    "能浏览所有条目",
	RoleViewPermissionOnlyMe: "仅能浏览自己创建的或者指派给自己的条目",
}

var RoleViewPermissionConfig = []common.BaseConfig{
	{
		Value: "view_all",
		Label: "能浏览所有条目",
	},
	{
		Value: "view_only_me",
		Label: "仅能浏览自己创建的或者指派给自己的条目",
	},
}

const (
	//可编辑所有
	RoleEditPermissionCreate string = "create"
	//仅能编辑进展
	RoleEditPermissionExportList string = "export_list"
)

var RoleEditPermissionMap = map[string]string{
	RoleEditPermissionCreate:     "创建",
	RoleEditPermissionExportList: "导出列表",
}

var RoleEditPermissionConfig = []common.BaseConfig{
	{
		Value: "create",
		Label: "创建",
	},
	{
		Value: "export_list",
		Label: "导出列表",
	},
}
