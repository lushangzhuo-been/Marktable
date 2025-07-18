package tmpl

import "mark3/types/common"

const (
	TmplModePublic  = "public"
	TmplModePrivate = "private"
)

var TmplModeMap = map[string]string{
	TmplModePublic:  "公共的",
	TmplModePrivate: "私有的",
}

var TmplModeConfig = []common.BaseConfig{
	{
		Value: "public",
		Label: "公共的",
	},
	{
		Value: "private",
		Label: "私有的",
	},
}
