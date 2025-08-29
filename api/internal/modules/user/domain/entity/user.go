package entity

type UserModel struct {
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string `json:"username" gorm:"column:username;comment:用户名;type:varchar(200);"`
	Password  string `json:"-" gorm:"column:password;comment:密码;type:varchar(200);"`
	FullName  string `json:"full_name" gorm:"column:full_name;comment:姓名;type:varchar(200);"`
	Email     string `json:"email" gorm:"column:email;comment:邮箱;type:varchar(50);"`
	Phone     string `json:"phone" gorm:"column:phone;comment:手机;type:varchar(50);"`
	Avatar    string `json:"avatar" gorm:"column:avatar;comment:头像;type:varchar(200);"`
	CreatedAt string `json:"-" gorm:"column:created_at;comment:创建时间;type:varchar(50);"`
	UpdatedAt string `json:"-" gorm:"column:updated_at;comment:修改时间;type:varchar(50);"`
}

func (UserModel) TableName() string {
	return "user"
}
