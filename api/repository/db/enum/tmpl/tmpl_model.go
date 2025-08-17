package tmpl

import "mark3/types/common"

const (
	TmplModePublic = "public"
)

var TmplModeMap = map[string]string{
	TmplModePublic: "公共的",
}

var TmplModeConfig = []common.BaseConfig{
	{
		Value: "public",
		Label: "公共的",
	},
}
