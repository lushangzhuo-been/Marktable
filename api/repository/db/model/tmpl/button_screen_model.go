package tmpl

type ButtonScreenModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	NodeId    int    `json:"node_id" gorm:"column:node_id;comment:节点ID;type:int(50);"`
	ButtonId  int    `json:"button_id" gorm:"column:button_id;comment:按钮ID;type:int(50);"`
	FieldKey  string `json:"field_key" gorm:"column:field_key;comment:字段key;type:varchar(200);"`
	CanModify string `json:"can_modify" gorm:"column:can_modify;comment:是否可编辑;type:varchar(50);"`
	Required  string `json:"required" gorm:"column:required;comment:是否必填;type:varchar(50);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (ButtonScreenModel) TableName() string {
	return "tmpl_button_screen"
}

type ButtonScreenCoordinateModel struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId       int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId     int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	NodeId     int    `json:"node_id" gorm:"column:node_id;comment:节点ID;type:int(50);"`
	ButtonId   int    `json:"button_id" gorm:"column:button_id;comment:按钮ID;type:int(50);"`
	Coordinate string `json:"module" gorm:"column:coordinate;comment:坐标;type:text;"`
}

func (ButtonScreenCoordinateModel) TableName() string {
	return "tmpl_button_screen_coordinate"
}
