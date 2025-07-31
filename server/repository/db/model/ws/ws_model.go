package ws

type WsModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name      string `json:"name" gorm:"column:name;comment:空间名称;type:varchar(50);"`
	Desc      string `json:"desc" gorm:"column:desc;comment:空间描述;type:text;"`
	Creator   int    `json:"creator" gorm:"column:creator;comment:创建人;type:int(11);"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (WsModel) TableName() string {
	return "ws"
}
