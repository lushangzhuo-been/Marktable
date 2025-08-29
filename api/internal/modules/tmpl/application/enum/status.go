package enum

import model "mark3/repository/db/model/tmpl"

var InitStatusList = []model.StatusModel{
	{
		Name:  "未开始",
		Color: "rgb(24, 144, 255)",
		Mode:  "first",
	},
	{
		Name:  "进行中",
		Color: "rgb(0, 221, 87)",
		Mode:  "other",
	},
	{
		Name:  "已完成",
		Color: "rgb(165, 167, 165)",
		Mode:  "other",
	},
}
