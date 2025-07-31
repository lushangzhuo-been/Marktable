package rule

type RuleModel struct {
	Id          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId        int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId      int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Name        string `json:"name" gorm:"column:name;comment:规则名称;type:varchar(50);"`
	RuleType    string `json:"rule_type" gorm:"column:rule_type;comment:规则类型;type:varchar(50);"`
	IsEnable    string `json:"is_enable" gorm:"column:is_enable;comment:是否启用(yes,no);type:varchar(50);"`
	IsDeleted   string `json:"is_deleted" gorm:"column:is_deleted;comment:是否删除(已删除yes,未删除no);type:varchar(50);"`
	Description string `json:"description" gorm:"column:description;comment:描述;type:varchar(50);"`
	Filter      string `json:"filter" gorm:"column:filter;comment:过滤条件;type:text;"`
	TriggerDay  int    `json:"trigger_day" gorm:"column:trigger_day;comment:触发天;type:int(50);"`
	TriggerTime string `json:"trigger_time" gorm:"column:trigger_time;comment:触发时间(时:分);type:varchar(50);"`
	FieldKey    string `json:"field_key" gorm:"column:field_key;comment:触发字段;type:varchar(50);"`
	OldValue    string `json:"old_value" gorm:"column:old_value;comment:旧值;type:text;"`
	NewValue    string `json:"new_value" gorm:"column:new_value;comment:新值;type:text;"`
	CreatedBy   int    `json:"created_by" gorm:"column:created_by;comment:创建人ID;type:int(50);"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (RuleModel) TableName() string {
	return "rule"
}

type RuleActionModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	RuleId        int    `json:"rule_id" gorm:"column:rule_id;index;comment:规则ID;type:int(50);"`
	ActionType    string `json:"action_type" gorm:"column:action_type;comment:动作类型;type:varchar(50);"`
	FieldKey      string `json:"field_key" gorm:"column:field_key;comment:字段名;type:varchar(50);"`
	FieldValue    string `json:"field_value" gorm:"column:field_value;comment:字段值;type:varchar(50);"`
	EmailTitle    string `json:"email_title" gorm:"column:email_title;comment:邮件主题;type:varchar(50);"`
	EmailContents string `json:"email_contents" gorm:"column:email_contents;comment:邮件内容;type:text;"`
	UserList      string `json:"user_list" gorm:"column:user_list;comment:收件人ID;type:varchar(50);"`
	IssueRoleList string `json:"issue_role_list" gorm:"column:issue_role_list;comment:收件人员字段;type:varchar(50);"`
	IsDeleted     string `json:"is_deleted" gorm:"column:is_deleted;comment:是否删除(已删除yes,未删除no);type:varchar(50);"`
}

func (RuleActionModel) TableName() string {
	return "rule_action"
}

type RuleLogModel struct {
	Id              int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId            int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId          int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	RuleId          int    `json:"rule_id" gorm:"column:rule_id;comment:规则ID;type:int(50);"`
	RuleExecuteUuid string `json:"rule_execute_uuid" gorm:"column:rule_execute_uuid;comment:规则执行UUID;type:varchar(50);"`
	RuleName        string `json:"rule_name" gorm:"column:rule_name;comment:规则名称;type:varchar(50);"`
	RuleType        string `json:"rule_type" gorm:"column:rule_type;comment:规则类型;type:varchar(50);"`
	ExecuteAt       string `json:"execute_at" gorm:"column:execute_at;comment:执行时间;type:varchar(50);"`
	DataId          string `json:"data_id" gorm:"column:data_id;comment:数据ID;type:varchar(255);"`
	DataName        string `json:"data_name" gorm:"column:data_name;comment:数据名称;type:varchar(1024);"`
	RuleData        string `json:"rule_data" gorm:"column:rule_data;comment:规则表json;type:text;"`
	RuleActionData  string `json:"rule_action_data" gorm:"column:rule_action_data;comment:动作表list<json>;type:text;"`
}

func (RuleLogModel) TableName() string {
	return "rule_log"
}

type RuleActionLogModel struct {
	Id              int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId            int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId          int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	RuleId          int    `json:"rule_id" gorm:"column:rule_id;comment:规则ID;type:int(50);"`
	RuleActionId    int    `json:"rule_action_id" gorm:"column:rule_action_id;comment:动作ID;type:int(50);"`
	RuleExecuteUuid string `json:"rule_execute_uuid" gorm:"column:rule_execute_uuid;comment:规则执行UUID;type:varchar(50);"`
	Status          string `json:"status" gorm:"column:status;comment:执行状态(no,success,fail);type:varchar(50);"`
	ActionType      string `json:"action_type" gorm:"column:action_type;comment:动作类型;type:varchar(50);"`
	EmailTitle      string `json:"email_title" gorm:"column:email_title;comment:邮件主题;type:varchar(50);"`
	EmailContents   string `json:"email_contents" gorm:"column:email_contents;comment:邮件内容;type:text;"`
	EmailAddress    string `json:"email_address" gorm:"column:email_address;comment:收件邮箱;type:text;"`
	DataId          string `json:"data_id" gorm:"column:data_id;comment:数据ID;type:varchar(50);"`
	FieldKey        string `json:"field_key" gorm:"column:field_key;comment:字段名;type:varchar(50);"`
	FieldValue      string `json:"field_value" gorm:"column:field_value;comment:字段值;type:varchar(50);"`
	ExecuteAt       string `json:"execute_at" gorm:"column:execute_at;comment:执行时间;type:varchar(50);"`
}

func (RuleActionLogModel) TableName() string {
	return "rule_action_log"
}
