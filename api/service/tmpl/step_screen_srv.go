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

type StepScreenSrv struct{}

func (s *StepScreenSrv) Config() (resp interface{}, err error) {
	var config types.StepScreenConfigResp
	config.RequiredConfig = common.FormatConfig(enum.ScreenRequiredMap)
	return config, nil
}

func (s *StepScreenSrv) Surplus(req types.StepScreenSurplusReq) (resp interface{}, err error) {
	//获取全部可用字段
	var allFields []model.FieldModel
	global.GVA_DB.Where("tmpl_id=? and level in ?", req.TmplId, []string{"level2", "level3"}).Find(&allFields)

	//获取已添加字段
	var addFields []model.StepScreenModel
	global.GVA_DB.Where("step_id=?", req.StepId).Find(&addFields)

	var addFieldsMap = make(map[string]model.StepScreenModel)
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

func (s *StepScreenSrv) List(req types.StepScreenListReq) (resp interface{}, err error) {
	var list []types.StepScreenVo
	global.GVA_DB.Raw("select a.id, a.field_key, b.name as field_name, a.required from tmpl_step_screen as a inner join tmpl_field as b on a.field_key=b.field_key and a.tmpl_id=b.tmpl_id where a.step_id=?", req.StepId).Scan(&list)

	var coordinate model.StepScreenCoordinateModel
	global.GVA_DB.Where("step_id=?", req.StepId).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var screensMap = make(map[string]types.StepScreenVo)
		for _, vo := range list {
			screensMap[vo.FieldKey] = vo
		}

		var sort []types.StepScreenVo
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

func (s *StepScreenSrv) Create(req types.StepScreenCreateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Create(&model.StepScreenModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		StepId:    req.StepId,
		FieldKey:  req.FieldKey,
		CanModify: "yes",
		Required:  enum.ButtonScreenRequiredNo,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StepScreenSrv) SetRequired(req types.StepScreenSetRequiredReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Updates(&model.StepScreenModel{
		Required: req.Required,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StepScreenSrv) Move(req types.StepScreenMoveReq) (resp interface{}, err error) {
	var coordinate model.StepScreenCoordinateModel
	global.GVA_DB.Where("step_id=?", req.StepId).First(&coordinate)
	if coordinate.Id == 0 {
		if err = global.GVA_DB.Create(&model.StepScreenCoordinateModel{
			WsId:       req.WsId,
			TmplId:     req.TmplId,
			StepId:     req.StepId,
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else {
		if err = global.GVA_DB.Where("step_id=?", req.StepId).Updates(&model.StepScreenCoordinateModel{
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}
	return
}

func (s *StepScreenSrv) Delete(req types.StepScreenDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Delete(&model.StepScreenModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
