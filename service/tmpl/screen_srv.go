package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"
	"strings"
)

type ScreenSrv struct{}

func (s *ScreenSrv) Config() (resp interface{}, err error) {
	var config types.ScreenConfigResp
	config.ModuleConfig = common.FormatConfig(enum.ScreenModuleMap)
	config.RequiredConfig = common.FormatConfig(enum.ScreenRequiredMap)
	return config, nil
}

func (s *ScreenSrv) Surplus(req types.ScreenSurplusReq) (resp interface{}, err error) {
	//获取全部可用字段
	var allFields []model.FieldModel
	global.GVA_DB.Where("tmpl_id=? and level in ?", req.TmplId, []string{"level2", "level3"}).Find(&allFields)

	//获取已添加字段
	var addFields []model.ScreenModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, req.Module).Find(&addFields)
	var addFieldsMap = make(map[string]model.ScreenModel)
	for _, screenField := range addFields {
		addFieldsMap[screenField.FieldKey] = screenField
	}

	var surplus []model.FieldModel
	for _, field := range allFields {
		if _, ok := addFieldsMap[field.FieldKey]; !ok {
			surplus = append(surplus, field)
		}
	}
	return surplus, nil
}

func (s *ScreenSrv) List(req types.ScreenListReq) (resp interface{}, err error) {
	var list []types.ScreenVo
	global.GVA_DB.Raw("select a.id, a.field_key, a.module, b.name as field_name, a.can_modify, a.required from tmpl_screen as a inner join tmpl_field as b on a.field_key=b.field_key and a.tmpl_id=b.tmpl_id where a.tmpl_id=? and a.module=?", req.TmplId, req.Module).Scan(&list)

	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, req.Module).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var screensMap = make(map[string]types.ScreenVo)
		for _, vo := range list {
			screensMap[vo.FieldKey] = vo
		}

		var sort []types.ScreenVo
		for _, key := range coordinateSlice {
			if screen, ok := screensMap[key]; ok {
				sort = append(sort, screen)
				delete(screensMap, key)
			}
		}

		for _, vo := range screensMap {
			sort = append(sort, vo)
		}
		return sort, nil
	}

	return list, nil
}

func (s *ScreenSrv) Create(req types.ScreenCreateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Create(&model.ScreenModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		FieldKey:  req.FieldKey,
		Module:    req.Module,
		CanModify: enum.ScreenCanModifyYes,
		Required:  enum.ScreenRequiredNo,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *ScreenSrv) Set(req types.ScreenSetReq) (resp interface{}, err error) {

	if _, ok := enum.ScreenRequiredMap[req.Required]; !ok {
		return nil, errors.New("参数错误")
	}

	//title, handle 字段不可操作
	var screen model.ScreenModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&screen).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("数据不存在")
	}

	if screen.FieldKey == "title" || screen.FieldKey == "handle" {
		return nil, errors.New("title,handle禁止操作")
	}

	screen.Required = req.Required
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Updates(&screen).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *ScreenSrv) Move(req types.ScreenMoveReq) (resp interface{}, err error) {
	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, req.Module).First(&coordinate)
	if coordinate.Id == 0 {
		if err = global.GVA_DB.Create(&model.ScreenCoordinateModel{
			WsId:       req.WsId,
			TmplId:     req.TmplId,
			Module:     req.Module,
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else {
		if err = global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, req.Module).Updates(&model.ScreenCoordinateModel{
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}
	return
}

func (s *ScreenSrv) Delete(req types.ScreenDeleteReq) (resp interface{}, err error) {
	//title, handle 字段不可操作
	var screen model.ScreenModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&screen).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("数据不存在")
	}

	if screen.FieldKey == "title" || screen.FieldKey == "handle" {
		return nil, errors.New("title,handle禁止操作")
	}

	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Delete(&model.ScreenModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
