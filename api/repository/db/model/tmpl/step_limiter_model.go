package tmpl

type StepLimiterModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	StepId    int    `json:"step_id" gorm:"column:step_id;comment:步骤ID;type:int(50);"`
	ButtonId  int    `json:"button_id" gorm:"column:button_id;comment:按钮ID;type:int(50);"`
	Mode      string `json:"mode" gorm:"column:mode;comment:类型;type:varchar(50);"`
	Name      string `json:"name" gorm:"column:name;comment:描述;type:text;"`
	Rule      string `json:"rule" gorm:"column:rule;comment:规则;type:text;"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (StepLimiterModel) TableName() string {
	return "tmpl_step_limiter"
}
