package right

import (
	"errors"
	"fmt"
	"mark3/global"
	enum "mark3/repository/db/enum/tmpl"
	wsEnum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/tmpl"
	wsModel "mark3/repository/db/model/ws"
)

type UserTmplRight struct {
	TmplId   int    `json:"tmpl_id"`
	TmplName string `json:"tmpl_name"`
	TmplMode string `json:"tmpl_mode"`

	WsIsMember string `json:"ws_is_member"`
	WsStatus   string `json:"ws_status"`
	WsRoleSign string `json:"ws_role_sign"`

	TmplIsMember       string `json:"tmpl_is_member"`
	TmplStatus         string `json:"tmpl_status"`
	TmplRoleSign       string `json:"tmpl_role_sign"`
	TmplViewPermission string `json:"tmpl_view_permission"`
	TmplEditPermission string `json:"tmpl_edit_permission"`
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
	UserTmplRightInfo.TmplMode = tmpl.Mode

	var wsMember wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=?", wsId, userid).First(&wsMember).Error; err != nil {
		return UserTmplRightInfo, nil
	}

	UserTmplRightInfo.WsIsMember = "yes"
	UserTmplRightInfo.WsRoleSign = wsMember.Role
	UserTmplRightInfo.WsStatus = wsMember.Status

	if tmpl.Mode == enum.TmplModePrivate {
		type TmplMemberAndRole struct {
			Status         string `json:"status"`
			RoleSign       string `json:"role_sign"`
			ViewPermission string `json:"view_permission"`
			EditPermission string `json:"edit_permission"`
		}
		var tmplMemberAndRole TmplMemberAndRole
		sql := "select a.status, b.sign as role_sign, b.view_permission, b.edit_permission from tmpl_member as a left join tmpl_role as b on a.role_id=b.id where a.ws_id=? and a.tmpl_id=? and a.userid=?"
		if err := global.GVA_DB.Raw(sql, wsId, tmplId, userid).Scan(&tmplMemberAndRole).Error; err != nil {
			return UserTmplRightInfo, nil
		}

		UserTmplRightInfo.TmplIsMember = "yes"
		UserTmplRightInfo.TmplStatus = tmplMemberAndRole.Status
		UserTmplRightInfo.TmplRoleSign = tmplMemberAndRole.RoleSign
		if tmplMemberAndRole.RoleSign == enum.RoleSignAdmin {
			UserTmplRightInfo.TmplViewPermission = enum.RoleViewPermissionAll
			UserTmplRightInfo.TmplEditPermission = fmt.Sprintf("%s,%s", enum.RoleEditPermissionCreate, enum.RoleEditPermissionExportList)
		} else {
			UserTmplRightInfo.TmplViewPermission = tmplMemberAndRole.ViewPermission
			UserTmplRightInfo.TmplEditPermission = tmplMemberAndRole.EditPermission
		}
	}
	return UserTmplRightInfo, nil
}

func (r UserTmplRight) CanAccess() error {
	//判断是否为空间成员
	if r.WsIsMember != "yes" {
		return errors.New("您不是空间成员，无权访问该空间")
	}
	if r.WsStatus != wsEnum.WsStatusOk {
		return errors.New("您已经被禁用，无权访问该空间")
	}

	if r.WsRoleSign == wsEnum.WsRoleAdmin {
		return nil
	}

	if r.TmplMode == enum.TmplModePrivate {
		if r.TmplIsMember != "yes" {
			return errors.New("您不是该流程成员，无权访问该流程")
		}
		if r.TmplStatus != enum.TmplStatusOk {
			return errors.New("您已经被禁用，无权访问该流程")
		}
	}

	return nil
}

func (r UserTmplRight) CanOperate() error {
	//判断是否为空间成员
	if r.WsIsMember != "yes" {
		return errors.New("您不是空间成员，无权访问该空间")
	}
	if r.WsStatus != wsEnum.WsStatusOk {
		return errors.New("您已经被禁用，无权访问该空间")
	}

	if r.WsRoleSign == wsEnum.WsRoleAdmin {
		return nil
	}

	if r.TmplMode == enum.TmplModePublic {
		return errors.New("您不是空间管理员，无权操作该流程")
	} else {
		if r.TmplIsMember != "yes" {
			return errors.New("您不是该流程成员，无权访问该流程")
		}
		if r.TmplStatus != enum.TmplStatusOk {
			return errors.New("您已经被禁用，无权访问该流程")
		}
		if r.TmplRoleSign != enum.RoleSignAdmin {
			return errors.New("您不是流程管理员，无权操作流程")
		}
	}
	return nil
}
