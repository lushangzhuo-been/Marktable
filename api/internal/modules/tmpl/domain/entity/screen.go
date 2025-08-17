package entity

type ScreenModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Module    string `json:"module" gorm:"column:module;comment:页面类型;type:varchar(50);"`
	FieldKey  string `json:"field_key" gorm:"column:field_key;comment:字段key;type:varchar(50);"`
	CanModify string `json:"can_modify" gorm:"column:can_modify;comment:是否可编辑;type:varchar(50);"`
	Required  string `json:"required" gorm:"column:required;comment:是否必填;type:varchar(50);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (ScreenModel) TableName() string {
	return "tmpl_screen"
}

type ScreenCoordinateModel struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId       int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId     int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Module     string `json:"module" gorm:"column:module;comment:页面类型;type:varchar(50);"`
	Coordinate string `json:"coordinate" gorm:"column:coordinate;comment:坐标;type:text;"`
}

func (ScreenCoordinateModel) TableName() string {
	return "tmpl_screen_coordinate"
}
