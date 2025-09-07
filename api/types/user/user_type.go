package user

import (
	"mark3/repository/db/model/user"
)

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" `
}

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResp struct {
	User         user.UserModel `json:"user"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
}

type UpdateReq struct {
	FullName string `form:"full_name" json:"full_name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
}

type UpdatePwdReq struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}

type Tmpl struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Mode           string `json:"mode"`
	RoleId         int    `json:"role_id"`
	RoleName       string `json:"role_name"`
	RoleSign       string `json:"role_sign"`
	ViewPermission string `json:"view_permission"`
	EditPermission string `json:"edit_permission"`
	Status         string `json:"status"`
}
type Ws struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	TmplList []Tmpl `json:"tmpl_list"`
}

type EmailParamsReq struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type RegisterV2Req struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

type ResetPwdEmailCodeReq struct {
	Username string `form:"username" json:"username" binding:"required"`
}

type ResetPwdReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

type UpdateEmailReq struct {
	Email string `form:"email" json:"email" binding:"required"`
	Code  string `form:"code" json:"code" binding:"required"`
}
