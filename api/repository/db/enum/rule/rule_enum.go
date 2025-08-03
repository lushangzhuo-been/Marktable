package rule

const (
	RuleEnableYes string = "yes"
	RuleEnableNo  string = "no"
)

var RuleEnableMap = map[string]string{
	RuleEnableYes: "是",
	RuleEnableNo:  "否",
}

const (
	RuleDeletedYes string = "yes"
	RuleDeletedNo  string = "no"
)

var RuleDeletedMap = map[string]string{
	RuleDeletedYes: "是",
	RuleDeletedNo:  "否",
}

const (
	RuleActionDeletedYes string = "yes"
	RuleActionDeletedNo  string = "no"
)

var RuleActionDeletedMap = map[string]string{
	RuleActionDeletedYes: "是",
	RuleActionDeletedNo:  "否",
}

const (
	RuleActionLogStatusNo      string = "no"
	RuleActionLogStatusSuccess string = "success"
	RuleActionLogStatusFail    string = "fail"
)

var RuleActionLogStatusMap = map[string]string{
	RuleActionLogStatusNo:      "未执行",
	RuleActionLogStatusSuccess: "成功",
	RuleActionLogStatusFail:    "失败",
}

const (
	RuleTypeUpdate   string = "update"
	RuleTypeSchedule string = "schedule"
	RuleTypeCreate   string = "create"
	RuleTypeDelete   string = "delete"
)

var RuleTypeMap = map[string]string{
	RuleTypeUpdate:   "字段变更",
	RuleTypeSchedule: "定时触发",
	RuleTypeCreate:   "创建任务",
	RuleTypeDelete:   "删除任务",
}

const (
	RuleActionTypeUpdate string = "update"
	RuleActionTypeEmail  string = "email"
)

var RuleActionTypeMap = map[string]string{
	RuleActionTypeUpdate: "修改字段",
	RuleActionTypeEmail:  "邮件通知",
}
