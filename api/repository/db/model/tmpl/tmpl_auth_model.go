package tmpl

type TmplAuthModel struct {
	Id            int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId          int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId        int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Mode          string `json:"mode" gorm:"column:mode;comment:权限类型;type:varchar(50);"`
	WsRoles       string `json:"ws_roles" gorm:"column:ws_roles;comment:空间角色;type:varchar(50);"`
	UserList      string `json:"user_list" gorm:"column:user_list;comment:人ID;type:varchar(50);"`
	IssueRoleList string `json:"issue_role_list" gorm:"column:issue_role_list;comment:人员字段;type:varchar(50);"`
}

func (TmplAuthModel) TableName() string {
	return "tmpl_auth"
}
