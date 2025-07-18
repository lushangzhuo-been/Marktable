package tmpl

type UserListScreenModel struct {
	Id     int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId   int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Userid int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	Screen string `json:"screen" gorm:"column:screen;comment:屏幕信息;type:text;"`
}

func (UserListScreenModel) TableName() string {
	return "user_list_screen"
}

type UserRecentlyVisitedLogModel struct {
	Id          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId        int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId      int    `json:"tmpl_id" gorm:"column:tmpl_id;comment:模板ID;type:int(50);"`
	Userid      int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	VisitedTime string `json:"visited_time" gorm:"column:visited_time;comment:访问时间;type:varchar(50);"`
}

func (UserRecentlyVisitedLogModel) TableName() string {
	return "user_recently_visited_log"
}

type UserHandleCommonlyUsedModel struct {
	Id                     int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId                   int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	Userid                 int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	CommonlyUsedUseridList string `json:"commonly_used_userid_list" gorm:"column:commonly_used_userid_list;comment:常用人ID;type:varchar(255);"`
	CreatedAt              string `json:"created_at" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (UserHandleCommonlyUsedModel) TableName() string {
	return "user_handle_commonly_used"
}

type UserCloseUserLogModel struct {
	Id          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId        int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	CloseUserid int    `json:"close_userid" gorm:"column:close_userid;comment:亲密用户;type:int(50);"`
	Userid      int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	VisitedTime string `json:"visited_time" gorm:"column:visited_time;comment:访问时间;type:varchar(50);"`
}
