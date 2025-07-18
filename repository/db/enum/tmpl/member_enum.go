package tmpl

import "mark3/types/common"

const (
	TmplStatusOk     = "ok"
	TmplStatusForbid = "forbid"
)

var TmplStatusMap = map[string]string{
	TmplStatusOk:     "正常",
	TmplStatusForbid: "禁用中",
}

var TmplStatusConfig = []common.BaseConfig{
	{
		Value: "ok",
		Label: "正常",
	},
	{
		Value: "forbid",
		Label: "禁用中",
	},
}
