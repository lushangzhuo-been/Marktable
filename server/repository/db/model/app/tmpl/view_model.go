package tmpl

type UserViewModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Userid    int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	Name      string `json:"name" gorm:"column:name;comment:名称;type:varchar(50);"`
	Mode      string `json:"mode" gorm:"column:mode;comment:类型;type:varchar(50);"`
	Filter    string `json:"filter" gorm:"column:filter;comment:条件;type:text;"`
	Lor       string `json:"lor" gorm:"column:lor;comment:逻辑关系;type:varchar(50);"`
	SortOrder string `json:"sort_order" gorm:"column:sort_order;comment:排序;type:text;"`
	Columns   string `json:"columns" gorm:"column:columns;comment:字段列;type:text;"`
	Pin       string `json:"pin" gorm:"column:pin;comment:是否钉住;type:varchar(50);"`
	CardGroup string `json:"card_group" gorm:"column:card_group;comment:看板分组条件;type:text;"`
}

func (UserViewModel) TableName() string {
	return "user_view"
}

type UserViewCoordinateModel struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId       int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId     int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Userid     int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	Coordinate string `json:"coordinate" gorm:"column:coordinate;comment:坐标;type:text;"`
}

func (UserViewCoordinateModel) TableName() string {
	return "user_view_coordinate"
}
