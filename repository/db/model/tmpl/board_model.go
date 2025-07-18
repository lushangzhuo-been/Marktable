package tmpl

type BoardModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId        int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	ViewId        int    `json:"view_id" gorm:"column:view_id;comment:user_view表ID;type:int(50);"`
	Userid        int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	Mode          string `json:"mode" gorm:"column:mode;comment:图表类型;type:varchar(50);"`
	Title         string `json:"title" gorm:"column:title;comment:标题;type:varchar(50);"`
	XAxis         string `json:"x_axis" gorm:"column:x_axis;comment:X轴;type:varchar(50);"`
	AxisDateMode  string `json:"axis_date_mode" gorm:"column:axis_date_mode;comment:X轴时间类型;type:varchar(50);"`
	YAxis         string `json:"y_axis" gorm:"column:y_axis;comment:Y轴;type:varchar(50);"`
	GroupAxis     string `json:"group_axis" gorm:"column:group_axis;comment:分组;type:varchar(50);"`
	GroupDateMode string `json:"group_date_mode" gorm:"column:group_date_mode;comment:分组为时间类型;type:varchar(50);"`
	Order         string `json:"order" gorm:"column:order;comment:排序;type:varchar(50);"`
	ShowEmpty     string `json:"show_empty" gorm:"column:show_empty;comment:是否显示空数据;type:varchar(50);"`
	Filter        string `json:"filter" gorm:"column:filter;comment:过滤条件;type:text;"`
	W             int    `json:"w" gorm:"column:w;comment:宽度;type:int(50);"`
	H             int    `json:"h" gorm:"column:h;comment:高度;type:int(50);"`
	X             int    `json:"x" gorm:"column:x;comment:X坐标;type:int(50);"`
	Y             int    `json:"y" gorm:"column:y;comment:Y坐标;type:int(50);"`
}

func (BoardModel) TableName() string {
	return "user_board"
}
