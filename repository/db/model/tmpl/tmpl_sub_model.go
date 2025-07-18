package tmpl

type TmplSubConfigModel struct {
	Id        int `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	SubTmplId int `json:"sub_tmpl_id" gorm:"column:sub_tmpl_id;index;comment:子模板ID;type:int(50);"`
}

func (TmplSubConfigModel) TableName() string {
	return "tmpl_sub_config"
}

type TmplSubObjModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	SubTmplId int    `json:"sub_tmpl_id" gorm:"column:sub_tmpl_id;index;comment:子模板ID;type:int(50);"`
	ObjId     string `json:"obj_id" gorm:"column:obj_id;comment:主流程数据ID;type:varchar(64);"`
	SubObjId  string `json:"sub_obj_id" gorm:"column:sub_obj_id;comment:子流程数据ID;type:varchar(64);"`
}

func (TmplSubObjModel) TableName() string {
	return "tmpl_sub_obj"
}
