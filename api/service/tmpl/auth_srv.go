package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	enumWs "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/service/app"
	types "mark3/types/tmpl"
)

type AuthSrv struct{}

func (s *AuthSrv) Config() (resp interface{}, err error) {
	var config types.AuthConfigResp
	config.RoleConfig = common.FormatConfig(enumWs.WsRoleMap)
	config.AuthConfig = common.FormatConfig(enum.AuthMap)
	return config, nil
}

func (s *AuthSrv) Update(req types.AuthUpdateReq) (resp interface{}, err error) {
	//判断是否存在
	var detailVo model.TmplAuthModel
	global.GVA_DB.Model(&model.TmplAuthModel{}).Where("ws_id=? and tmpl_id=? and mode=?",
		req.WsId, req.TmplId, req.Mode).First(&detailVo)

	if detailVo.Id == 0 {
		// 创建
		detailVo.WsId = req.WsId
		detailVo.TmplId = req.TmplId
		detailVo.Mode = req.Mode
		detailVo.WsRoles = req.WsRoles
		detailVo.UserList = req.UserList
		detailVo.IssueRoleList = req.IssueRoleList

		if err = global.GVA_DB.Create(&detailVo).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}

	// 更新规则
	detailVo.WsRoles = req.WsRoles
	detailVo.UserList = req.UserList
	detailVo.IssueRoleList = req.IssueRoleList

	if err = global.GVA_DB.Save(&detailVo).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	return
}

func (s *AuthSrv) Detail(req types.AuthDetailReq) (resp interface{}, err error) {
	type TmplAuthVo struct {
		model.TmplAuthModel
		UserListInfo []user.UserModel `json:"user_list_info"`
	}
	var detailVo TmplAuthVo
	global.GVA_DB.Model(&model.TmplAuthModel{}).Where("ws_id=? and tmpl_id=? and mode=?",
		req.WsId, req.TmplId, req.Mode).First(&detailVo)

	if detailVo.Id != 0 {
		// 查询用户
		users := app.GetUsersByUserIdStr(detailVo.UserList)
		detailVo.UserListInfo = users

		return detailVo, nil
	} else {
		return nil, nil
	}
}
