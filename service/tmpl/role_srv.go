package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	commonTypes "mark3/types/common"
	types "mark3/types/tmpl"
	"strings"
)

type RoleSrv struct{}

func (s *RoleSrv) Config() (resp interface{}, err error) {
	var config types.RoleConfigResp
	config.ViewPermission = common.FormatConfig(enum.RoleViewPermissionMap)
	config.EditPermission = common.FormatConfig(enum.RoleEditPermissionMap)
	return config, nil
}

func (s *RoleSrv) List(req types.RoleListReq) (resp interface{}, err error) {
	db := global.GVA_DB.Model(&model.RoleModel{})
	db = db.Where("tmpl_id=?", req.TmplId)
	if req.Key != "" {
		db = db.Where("name like ?", "%"+req.Key+"%")
	}
	var cnt int64
	db.Count(&cnt)

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	var roles []model.RoleModel
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&roles)

	resp = commonTypes.BasePageResp{
		List: roles,
		Cnt:  cnt,
	}
	return
}

func (s *RoleSrv) Create(req types.RoleCreateReq) (resp interface{}, err error) {
	_, ok := enum.RoleViewPermissionMap[req.ViewPermission]
	if !ok {
		return nil, errors.New("参数错误")
	}

	for _, v := range strings.Split(req.EditPermission, ",") {
		_, ok = enum.RoleEditPermissionMap[v]
		if !ok {
			return nil, errors.New("参数错误")
		}
	}

	if err = global.GVA_DB.Create(&model.RoleModel{
		WsId:           req.WsId,
		TmplId:         req.TmplId,
		Name:           req.Name,
		Desc:           req.Desc,
		Sign:           enum.RoleSignGeneral,
		ViewPermission: req.ViewPermission,
		EditPermission: req.EditPermission,
		CreatedAt:      common.GetCurrentTime(),
		UpdatedAt:      common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *RoleSrv) Update(req types.RoleUpdateReq) (resp interface{}, err error) {
	_, ok := enum.RoleViewPermissionMap[req.ViewPermission]
	if !ok {
		return nil, errors.New("参数错误")
	}

	for _, v := range strings.Split(req.EditPermission, ",") {
		_, ok = enum.RoleEditPermissionMap[v]
		if !ok {
			return nil, errors.New("参数错误")
		}
	}

	var role model.RoleModel
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).First(&role).Error; err != nil {
		return
	}

	if role.Sign == enum.RoleSignAdmin {
		return nil, errors.New("无法修改此项数据")
	}

	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Updates(&model.RoleModel{
		Name:           req.Name,
		Desc:           req.Desc,
		ViewPermission: req.ViewPermission,
		EditPermission: req.EditPermission,
		UpdatedAt:      common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *RoleSrv) Delete(req types.RoleDeleteReq) (resp interface{}, err error) {
	var role model.RoleModel
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).First(&role).Error; err != nil {
		return
	}

	if role.Sign == enum.RoleSignAdmin {
		return nil, errors.New("无法删除管理员角色")
	}

	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Delete(&model.RoleModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
