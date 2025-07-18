package tmpl

type TmplModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;index;comment:空间ID;type:int(50);"`
	WsFileId      int    `json:"ws_file_id" gorm:"column:ws_file_id;index;comment:流程归属文件ID;type:int(50);"`
	Name          string `json:"name" gorm:"column:name;comment:模板名;type:varchar(50);"`
	Desc          string `json:"desc" gorm:"column:desc;comment:模板描述;type:text;"`
	Mode          string `json:"mode" gorm:"column:mode;comment:模式;type:varchar(10);"`
	CreatedAt     string `json:"created_at" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt     string `json:"updated_at" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
	CreatedUserId int    `json:"created_user_id" gorm:"column:created_user_id;comment:创建人ID;type:int(11);"`
	UpdatedUserId int    `json:"updated_user_id" gorm:"column:updated_user_id;comment:修改人ID;type:int(11);"`
}

func (TmplModel) TableName() string {
	return "tmpl"
}
