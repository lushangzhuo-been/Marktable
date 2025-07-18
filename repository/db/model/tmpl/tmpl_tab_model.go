package tmpl

type TmplTabModel struct {
	Id     int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId   int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Tab    string `json:"tab" gorm:"column:tab;comment:页签;type:varchar(255);"`
}

func (TmplTabModel) TableName() string {
	return "tmpl_tab"
}
