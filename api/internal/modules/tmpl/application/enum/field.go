package enum

import (
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/types/common"
)

var FieldLevel1List = []entity.FieldModel{
	{
		FieldKey:   "status",
		Name:       "状态",
		Desc:       "",
		Mode:       FieldStatusCom,
		EnumValues: "",
		Expr:       "",
		Level:      "level1",
	},
	{
		FieldKey:   "creator",
		Name:       "创建人",
		Desc:       "",
		Mode:       FieldPersonCom,
		EnumValues: "",
		Expr:       "single",
		Level:      "level1",
	},
	{
		FieldKey:   "created_at",
		Name:       "创建时间",
		Desc:       "",
		Mode:       FieldTimeCom,
		EnumValues: "",
		Expr:       "YmdHis",
		Level:      "level1",
	},
	{
		FieldKey:   "updated_at",
		Name:       "最后更新时间",
		Desc:       "",
		Mode:       FieldTimeCom,
		EnumValues: "",
		Expr:       "YmdHis",
		Level:      "level1",
	},
}

var FieldLevel2List = []entity.FieldModel{
	{
		FieldKey:   "title",
		Name:       "标题",
		Desc:       "",
		Mode:       FieldTextCom,
		EnumValues: "",
		Expr:       "",
		Level:      "level2",
	},
	{
		FieldKey:   "handler",
		Name:       "处理人",
		Desc:       "",
		Mode:       FieldPersonCom,
		EnumValues: "",
		Expr:       "single",
		Level:      "level2",
	},
}

var FieldLevel3List = []entity.FieldModel{
	{
		FieldKey:   "custom_field_priority",
		Name:       "优先级",
		Desc:       "",
		Mode:       FieldSelectCom,
		EnumValues: "[{\"value\":\"高\",\"color\":\"rgb(235, 85, 79)\"},{\"value\":\"中\",\"color\":\"rgb(250, 179, 3)\"},{\"value\":\"低\",\"color\":\"rgb(61, 197, 255)\"}]",
		Expr:       "",
		Level:      "level3",
	},
	{
		FieldKey:   "custom_field_deadline",
		Name:       "截止日期",
		Desc:       "",
		Mode:       FieldDateCom,
		EnumValues: "",
		Expr:       "",
		Level:      "level3",
	},
	{
		FieldKey:   "custom_field_description",
		Name:       "描述",
		Desc:       "",
		Mode:       FieldTextareaCom,
		EnumValues: "",
		Expr:       "",
		Level:      "level3",
	},
}

const (
	FieldTextCom        string = "text_com"
	FieldNumberCom      string = "number_com"
	FieldTextareaCom    string = "textarea_com"
	FieldHtmlTextCom    string = "html_text_com"
	FieldSelectCom      string = "select_com"
	FieldMultiSelectCom string = "multi_select_com"
	FieldLinkCom        string = "link_com"
	FieldTimeCom        string = "time_com"
	FieldDateCom        string = "date_com"
	FieldProgressCom    string = "progress_com"
	FieldPersonCom      string = "person_com"
	FieldRelatedCom     string = "related_com"
	//FieldFileCom   string = "file_com"
	FieldStatusCom string = "status_com"
	FieldBlankCom  string = "blank_com"
)

var FieldModeMap = map[string]string{
	FieldTextCom:        "文本框",
	FieldNumberCom:      "数字",
	FieldTextareaCom:    "富文本框",
	FieldHtmlTextCom:    "Html文本编辑器",
	FieldSelectCom:      "单选",
	FieldMultiSelectCom: "多选",
	FieldLinkCom:        "链接",
	FieldTimeCom:        "时间",
	FieldDateCom:        "日期",
	FieldProgressCom:    "进度条",
	FieldPersonCom:      "人员",
	FieldRelatedCom:     "关联字段",
	//FieldFileCom:   "文件",
}

var FieldModeOpMap = map[string][]common.BaseConfig{
	FieldTextCom: {{
		Value: "like",
		Label: "相似",
	}},
	FieldNumberCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$gt",
		Label: "大于",
	}, {
		Value: "$gte",
		Label: "大于等于",
	}, {
		Value: "$lt",
		Label: "小于",
	}, {
		Value: "$lte",
		Label: "小于等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldTextareaCom: {{
		Value: "like",
		Label: "相似",
	}},
	FieldSelectCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldMultiSelectCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldTimeCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$gt",
		Label: "大于",
	}, {
		Value: "$gte",
		Label: "大于等于",
	}, {
		Value: "$lt",
		Label: "小于",
	}, {
		Value: "$lte",
		Label: "小于等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldDateCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$gt",
		Label: "大于",
	}, {
		Value: "$gte",
		Label: "大于等于",
	}, {
		Value: "$lt",
		Label: "小于",
	}, {
		Value: "$lte",
		Label: "小于等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldPersonCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldRelatedCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldLinkCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldProgressCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$gt",
		Label: "大于",
	}, {
		Value: "$gte",
		Label: "大于等于",
	}, {
		Value: "$lt",
		Label: "小于",
	}, {
		Value: "$lte",
		Label: "小于等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
	FieldStatusCom: {{
		Value: "$eq",
		Label: "等于",
	}, {
		Value: "$ne",
		Label: "不等于",
	}},
}

var FieldNumberComMap = map[string]string{
	"":  "无",
	"￥": "￥",
	"$": "$",
	"%": "%",
}

var FieldTimeComMap = map[string]string{
	"Ymd":    "年-月-日",
	"YmdHis": "年-月-日 时-分-秒",
}

var FieldPersonComMap = map[string]string{
	"single":   "单人",
	"multiple": "多人",
}

var FieldNumberComConfig = []common.BaseConfig{
	{
		Value: "",
		Label: "无",
	},
	{
		Value: "CNY",
		Label: "￥",
	},
	{
		Value: "USD",
		Label: "$",
	},
	{
		Value: "PCT",
		Label: "%",
	},
}
