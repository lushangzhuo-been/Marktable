package tmpl

type FieldModel struct {
	ID            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId        int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	FieldKey      string `json:"field_key" gorm:"column:field_key;comment:字段key;type:varchar(200);"`
	Name          string `json:"name" gorm:"column:name;comment:字段名;type:varchar(50);"`
	Desc          string `json:"desc" gorm:"column:desc;comment:字段描述;type:varchar(200);"`
	Mode          string `json:"mode" gorm:"column:mode;comment:字段类型;type:varchar(50);"`
	EnumValues    string `json:"enum_values" gorm:"column:enum_values;comment:枚举值;type:text;"`
	Expr          string `json:"expr" gorm:"column:expr;comment:表达式;type:text;"`
	Level         string `json:"level" gorm:"column:level;comment:级别;type:varchar(50);"`
	ReadOnlyRule  string `json:"read_only_rule" gorm:"column:read_only_rule;comment:只读规则;type:text;"`
	RelatedTmplId int    `json:"related_tmpl_id" gorm:"column:related_tmpl_id;comment:关联模板ID;type:int(50);"`
	CreatedAt     string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt     string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (FieldModel) TableName() string {
	return "tmpl_field"
}
