package entity

type StepModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId        int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	StartStatusId int    `json:"start_status_id" gorm:"column:start_status_id;comment:状态Id;type:int(50);"`
	EndStatusId   int    `json:"end_status_id" gorm:"column:end_status_id;comment:状态Id;type:int(50);"`
	CreatedAt     string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt     string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (StepModel) TableName() string {
	return "tmpl_step"
}
