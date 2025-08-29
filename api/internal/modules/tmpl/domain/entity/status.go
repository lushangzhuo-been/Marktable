package entity

type StatusModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Name      string `json:"name" gorm:"column:name;comment:状态名;type:varchar(50);"`
	Color     string `json:"color" gorm:"column:color;comment:状态颜色;type:varchar(50);"`
	Mode      string `json:"mode" gorm:"column:mode;comment:状态类型;type:varchar(50);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (StatusModel) TableName() string {
	return "tmpl_status"
}

type StatusCoordinateModel struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId       int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId     int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Coordinate string `json:"module" gorm:"column:coordinate;comment:坐标;type:text;"`
}

func (StatusCoordinateModel) TableName() string {
	return "tmpl_status_coordinate"
}
