package ws_file

// WsFileModel 空间下的文件夹
type WsFileModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;index;comment:空间ID;type:int(50);"`
	Name          string `json:"name" gorm:"column:name;comment:文件夹名;type:varchar(50);"`
	Desc          string `json:"desc" gorm:"column:desc;comment:文件夹名;type:varchar(50);"`
	Coordinate    string `json:"coordinate" gorm:"column:coordinate;comment:坐标-流程ids-排序;type:text;"`
	CreatedAt     string `json:"created_at" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt     string `json:"updated_at" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
	CreatedUserId int    `json:"created_user_id" gorm:"column:created_user_id;comment:创建人ID;type:int(11);"`
	UpdatedUserId int    `json:"updated_user_id" gorm:"column:updated_user_id;comment:修改人ID;type:int(11);"`
}

func (WsFileModel) TableName() string {
	return "ws_file"
}
