package tmpl

type MemberModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(50);"`
	TmplId    int    `json:"tmpl_id" gorm:"column:tmpl_id;index;comment:模板ID;type:int(50);"`
	Userid    int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(50);"`
	RoleId    int    `json:"role_id" gorm:"column:role_id;comment:角色ID;type:int(50);"`
	Status    string `json:"status" gorm:"column:status;comment:状态;type:varchar(50);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (MemberModel) TableName() string {
	return "tmpl_member"
}
