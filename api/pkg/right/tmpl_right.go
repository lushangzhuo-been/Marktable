package right

import (
	"errors"
	"mark3/global"
	wsEnum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/tmpl"
	wsModel "mark3/repository/db/model/ws"
)

type UserTmplRight struct {
	TmplId   int    `json:"tmpl_id"`
	TmplName string `json:"tmpl_name"`

	WsIsMember string `json:"ws_is_member"`
	WsStatus   string `json:"ws_status"`
	WsRoleSign string `json:"ws_role_sign"`
}

func NewUserTmplRight(userid int, wsId int, tmplId int) (UserTmplRight, error) {
	var UserTmplRightInfo UserTmplRight
	//判断空间是否存在
	var ws wsModel.WsModel
	if err := global.GVA_DB.Where("id=?", wsId).First(&ws).Error; err != nil {
		return UserTmplRightInfo, errors.New("空间不存在，请刷新页面")
	}

	var tmpl model.TmplModel
	if err := global.GVA_DB.Where("id=?", tmplId).First(&tmpl).Error; err != nil {
		return UserTmplRightInfo, errors.New("流程不存在，请刷新页面")
	}

	UserTmplRightInfo.TmplId = tmpl.Id
	UserTmplRightInfo.TmplName = tmpl.Name

	var wsMember wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=?", wsId, userid).First(&wsMember).Error; err != nil {
		return UserTmplRightInfo, errors.New("您不是空间成员，无权访问该空间")
	}
	if wsMember.Status != wsEnum.WsStatusOk {
		return UserTmplRightInfo, errors.New("您已经被禁用，无权访问该空间")
	}

	UserTmplRightInfo.WsIsMember = "yes"
	UserTmplRightInfo.WsRoleSign = wsMember.Role
	UserTmplRightInfo.WsStatus = wsMember.Status

	return UserTmplRightInfo, nil
}

func (r UserTmplRight) CanManage() error {
	// 判断是否为空间管理员
	if r.WsRoleSign != wsEnum.WsRoleAdmin {
		return errors.New("您不是空间管理员")
	}
	return nil
}
