package tmpl

type NodeModel struct {
	Id          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId        int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId      int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Mode        string `json:"mode" gorm:"column:mode;comment:节点类型;type:varchar(50);"`
	Name        string `json:"name" gorm:"column:name;comment:节点名称;type:varchar(50);"`
	StatusText  string `json:"status_text" gorm:"column:status_text;comment:状态描述;type:varchar(50);"`
	StatusColor string `json:"status_color" gorm:"column:status_color;comment:状态颜色;type:varchar(50);"`
	CreatedAt   string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt   string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (NodeModel) TableName() string {
	return "tmpl_node"
}
