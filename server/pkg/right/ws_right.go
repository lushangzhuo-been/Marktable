package right

import (
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/ws"
)

func JudgeWsMember(userid int, wsId int) error {
	//判断空间是否存在
	var ws model.WsModel
	if err := global.GVA_DB.Where("id=?", wsId).First(&ws).Error; err != nil {
		return errors.New("空间不存在，请刷新页面")
	}

	var member model.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=?", wsId, userid).First(&member).Error; err != nil {
		return errors.New("您不是空间成员，无权访问该空间")
	}
	if member.Status != enum.WsStatusOk {
		return errors.New("您已经被禁用，无权访问该空间")
	}
	return nil
}

func JudgeWsAdmin(userid int, wsId int) error {
	//判断空间是否存在
	var ws model.WsModel
	if err := global.GVA_DB.Where("id=?", wsId).First(&ws).Error; err != nil {
		return errors.New("空间不存在，请刷新页面")
	}

	var member model.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=?", wsId, userid).First(&member).Error; err != nil {
		return errors.New("您不是空间成员，无权访问该空间")
	}

	if member.Role != enum.WsRoleAdmin {
		return errors.New("您不是空间管理员，无权操作空间")
	}

	if member.Status != enum.WsStatusOk {
		return errors.New("您已经被禁用，无权访问该空间")
	}

	return nil
}
