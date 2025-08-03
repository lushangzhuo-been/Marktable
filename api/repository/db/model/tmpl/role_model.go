package tmpl

type RoleModel struct {
	Id             int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId           int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId         int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Sign           string `json:"sign" gorm:"column:sign;comment:是否为管理员标示;type:varchar(50);"`
	Name           string `json:"name" gorm:"column:name;comment:角色名称;type:varchar(50);"`
	Desc           string `json:"desc" gorm:"column:desc;comment:角色描述;type:varchar(200);"`
	ViewPermission string `json:"view_permission" gorm:"column:view_permission;comment:浏览权限;type:varchar(50);"`
	EditPermission string `json:"edit_permission" gorm:"column:edit_permission;comment:编辑时间;type:varchar(50);"`
	CreatedAt      string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt      string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (RoleModel) TableName() string {
	return "tmpl_role"
}
