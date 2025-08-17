package entity

// WsFileAndTmplMergedModel 空间下 文件夹和流程id 合并表
type WsFileAndTmplMergedModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;index;comment:空间ID;type:int(50);"`
	ImplTableName string `json:"impl_table_name" gorm:"column:impl_table_name;comment:合并源表名称;type:varchar(50);"`
	ImplTableId   int    `json:"impl_table_id" gorm:"column:impl_table_id;index;comment:空间ID;type:int(50);"`
	CreatedAt     string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt     string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (WsFileAndTmplMergedModel) TableName() string {
	return "ws_file_and_tmpl_merged"
}

// WsFileAndTmplMergedCoordinateModel  空间下 文件夹和流程id 合并表的排序表
type WsFileAndTmplMergedCoordinateModel struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId       int    `json:"ws_id" gorm:"column:ws_id;index;comment:空间ID;type:int(50);"`
	Coordinate string `json:"coordinate" gorm:"column:coordinate;comment:坐标-流程ids-排序;type:text;"`
	CreatedAt  string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt  string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (WsFileAndTmplMergedCoordinateModel) TableName() string {
	return "ws_file_and_tmpl_merged_coordinate"
}
