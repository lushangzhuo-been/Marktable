package tmpl

type ButtonModel struct {
	Id           int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId         int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId       int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	NodeId       int    `json:"node_id" gorm:"column:node_id;comment:节点ID;type:int(50);"`
	Name         string `json:"name" gorm:"column:name;comment:按钮名称;type:varchar(50);"`
	TargetNodeId int    `json:"target_node_id" gorm:"column:target_node_id;comment:目标节点;type:int(50);"`
	CreatedAt    string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt    string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (ButtonModel) TableName() string {
	return "tmpl_button"
}
