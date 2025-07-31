package ws

type MemberModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	WsId      int    `json:"ws_id" gorm:"column:ws_id;comment:空间ID;type:int(11);"`
	Userid    int    `json:"userid" gorm:"column:userid;comment:用户ID;type:int(11);"`
	Role      string `json:"role" gorm:"column:role;comment:角色;type:varchar(50);"`
	Status    string `json:"status" gorm:"column:status;comment:状态;type:varchar(50);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(200);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(200);"`
}

func (MemberModel) TableName() string {
	return "ws_member"
}
