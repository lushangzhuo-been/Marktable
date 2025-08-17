package enum

import "mark3/types/common"

const (
	ScreenModuleCreate string = "create"
	ScreenModuleDetail string = "detail"
)

var ScreenModuleMap = map[string]string{
	ScreenModuleCreate: "创建界面",
	ScreenModuleDetail: "详情界面",
}

var ScreenModuleConfig = []common.BaseConfig{
	{
		Value: "create",
		Label: "创建界面",
	},
	{
		Value: "detail",
		Label: "详情界面",
	},
}

const (
	ScreenCanModifyYes string = "yes"
	ScreenCanModifyNo  string = "no"
)

var ScreenCanModifyMap = map[string]string{
	ScreenCanModifyYes: "是",
	ScreenCanModifyNo:  "否",
}

var ScreenCanModifyConfig = []common.BaseConfig{
	{
		Value: "yes",
		Label: "是",
	},
	{
		Value: "no",
		Label: "否",
	},
}

const (
	ScreenRequiredYes string = "yes"
	ScreenRequiredNo  string = "no"
)

var ScreenRequiredMap = map[string]string{
	ScreenRequiredYes: "是",
	ScreenRequiredNo:  "否",
}

var ScreenRequiredConfig = []common.BaseConfig{
	{
		Value: "yes",
		Label: "是",
	},
	{
		Value: "no",
		Label: "否",
	},
}

// tmpl_screen_coordinate 表 module枚举
const (
	ScreenCoordinateModuleCreate  string = "create"
	ScreenCoordinateModuleDetail  string = "detail"
	ScreenCoordinateModuleUpdate  string = "update"
	ScreenCoordinateModuleList    string = "list"
	ScreenCoordinateModuleField   string = "field"
	ScreenCoordinateModuleTab     string = "tab"
	ScreenCoordinateModuleSubTmpl string = "sub_tmpl"
)

var ScreenCoordinateModuleMap = map[string]string{
	ScreenCoordinateModuleCreate:  "创建界面",
	ScreenCoordinateModuleDetail:  "详情界面",
	ScreenCoordinateModuleField:   "配置字段页",
	ScreenCoordinateModuleTab:     "配置标签页",
	ScreenCoordinateModuleSubTmpl: "配置子任务页",
}
