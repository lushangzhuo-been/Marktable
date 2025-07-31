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

type ButtonScreenSrv struct{}

func (s *ButtonScreenSrv) Config() (resp interface{}, err error) {
	var config types.ButtonScreenConfigResp
	config.RequiredConfig = common.FormatConfig(enum.ScreenRequiredMap)
	return config, nil
}

func (s *ButtonScreenSrv) Surplus(req types.ButtonScreenSurplusReq) (resp interface{}, err error) {
	//获取全部可用字段
	var allFields []model.FieldModel
	global.GVA_DB.Where("tmpl_id=? and level in ?", req.TmplId, []string{"level2", "level3"}).Find(&allFields)

	//获取已添加字段
	var addFields []model.ButtonScreenModel
	global.GVA_DB.Where("button_id=?", req.ButtonId).Find(&addFields)

	var addFieldsMap = make(map[string]model.ButtonScreenModel)
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

func (s *ButtonScreenSrv) List(req types.ButtonScreenListReq) (resp interface{}, err error) {
	var list []types.ButtonScreenVo
	global.GVA_DB.Raw("select a.id, a.field_key, b.name as field_name, a.required from tmpl_button_screen as a inner join tmpl_field as b on a.field_key=b.field_key and a.tmpl_id=b.tmpl_id where a.button_id=?", req.ButtonId).Scan(&list)

	var coordinate model.ButtonScreenCoordinateModel
	global.GVA_DB.Where("button_id=?", req.ButtonId).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var screensMap = make(map[string]types.ButtonScreenVo)
		for _, vo := range list {
			screensMap[vo.FieldKey] = vo
		}

		var sort []types.ButtonScreenVo
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

func (s *ButtonScreenSrv) Create(req types.ButtonScreenCreateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Create(&model.ButtonScreenModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		NodeId:    req.NodeId,
		ButtonId:  req.ButtonId,
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

func (s *ButtonScreenSrv) SetRequired(req types.ButtonScreenSetRequiredReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Updates(&model.ButtonScreenModel{
		Required: req.Required,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *ButtonScreenSrv) Move(req types.ButtonScreenMoveReq) (resp interface{}, err error) {
	var coordinate model.ButtonScreenCoordinateModel
	global.GVA_DB.Where("button_id=?", req.ButtonId).First(&coordinate)
	if coordinate.Id == 0 {
		if err = global.GVA_DB.Create(&model.ButtonScreenCoordinateModel{
			WsId:       req.WsId,
			TmplId:     req.TmplId,
			NodeId:     req.NodeId,
			ButtonId:   req.ButtonId,
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else {
		if err = global.GVA_DB.Where("button_id=?", req.ButtonId).Updates(&model.ButtonScreenCoordinateModel{
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}
	return
}

func (s *ButtonScreenSrv) Delete(req types.ButtonScreenDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Delete(model.ButtonScreenModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
