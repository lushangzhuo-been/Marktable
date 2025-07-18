package tmpl

import (
	"encoding/json"
	"mark3/global"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/app/tmpl"
	"strings"
)

type Caller struct{}

func (s *Caller) GetFields(tmplId int) ([]model.FieldModel, error) {
	var fields []model.FieldModel
	if err := global.GVA_DB.Where("tmpl_id=?", tmplId).Find(&fields).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 创建人、创建时间、最后更新时间默认展示字段末尾
	var fieldsRes []model.FieldModel
	var fieldsSome []model.FieldModel
	for _, field := range fields {
		if field.FieldKey == "creator" || field.FieldKey == "created_at" || field.FieldKey == "updated_at" {
			fieldsSome = append(fieldsSome, field)
		} else {
			fieldsRes = append(fieldsRes, field)
		}
	}
	for _, field := range fieldsSome {
		fieldsRes = append(fieldsRes, field)
	}

	return fieldsRes, nil
}

func (s *Caller) GetFieldsMap(tmplId int) map[string]model.FieldModel {
	fields, err := s.GetFields(tmplId)
	if err != nil {
		return nil
	}
	var fieldsMap = make(map[string]model.FieldModel)
	for _, field := range fields {
		fieldsMap[field.FieldKey] = field
	}
	return fieldsMap
}

func (s *Caller) GetFieldsMapName(tmplId int) map[string]model.FieldModel {
	fields, err := s.GetFields(tmplId)
	if err != nil {
		return nil
	}
	var fieldsMap = make(map[string]model.FieldModel)
	for _, field := range fields {
		fieldsMap[field.Name] = field
	}
	return fieldsMap
}

func (s *Caller) GetScreenByModule(tmplId int, module string) (screenVoList []types.ScreenVo, err error) {
	var screens []ScreenCommon
	if err = global.GVA_DB.Model(model.ScreenModel{}).Where("tmpl_id=? and module=?", tmplId, module).Find(&screens).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	if module == "detail" {
		for _, field := range enum.FieldLevel1List {
			var screen ScreenCommon
			//思考，是否需要重新写，将status的级别定义为1.5级
			if field.FieldKey == "status" {
				screen = ScreenCommon{
					FieldKey:  field.FieldKey,
					CanModify: enum.ScreenCanModifyYes,
					Required:  enum.ScreenRequiredYes,
				}
			} else {
				screen = ScreenCommon{
					FieldKey:  field.FieldKey,
					CanModify: enum.ScreenCanModifyNo,
					Required:  enum.ScreenRequiredNo,
				}
			}
			screens = append(screens, screen)
		}
	}

	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", tmplId, module).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var screensMap = make(map[string]ScreenCommon)
		for _, vo := range screens {
			screensMap[vo.FieldKey] = vo
		}

		var sort []ScreenCommon
		for _, key := range coordinateSlice {
			sort = append(sort, screensMap[key])
			delete(screensMap, key)
		}
		for _, vo := range screensMap {
			sort = append(sort, vo)
		}
		screenVoList = s.screenFormat(tmplId, sort)
		return
	}
	screenVoList = s.screenFormat(tmplId, screens)
	return
}

type ScreenCommon struct {
	FieldKey  string `json:"field_key"`
	CanModify string `json:"can_modify"`
	Required  string `json:"required"`
}

func (s *Caller) screenFormat(tmplId int, screens []ScreenCommon) []types.ScreenVo {
	var screenVoList []types.ScreenVo
	fieldsMap := s.GetFieldsMap(tmplId)
	for _, screen := range screens {
		var field model.FieldModel
		var ok bool
		if field, ok = fieldsMap[screen.FieldKey]; !ok {
			continue
		}
		var enumValues []interface{}
		json.Unmarshal([]byte(field.EnumValues), &enumValues)
		screenVo := types.ScreenVo{
			FieldKey:      field.FieldKey,
			Name:          field.Name,
			Desc:          field.Desc,
			Mode:          field.Mode,
			EnumValues:    enumValues,
			Expr:          field.Expr,
			Level:         field.Level,
			ReadOnlyRule:  field.ReadOnlyRule,
			RelatedTmplId: field.RelatedTmplId,
			CanModify:     screen.CanModify,
			Required:      screen.Required,
		}
		screenVoList = append(screenVoList, screenVo)
	}
	return screenVoList
}

func (s *Caller) GetFirstStatus(tmplId int) (*model.StatusModel, error) {
	var status model.StatusModel
	if err := global.GVA_DB.Where("tmpl_id=? and mode=?", tmplId, "first").First(&status).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return &status, nil
}

func (s *Caller) GetStatus(tmplId int, statusId int) (*model.StatusModel, error) {
	var status model.StatusModel
	if err := global.GVA_DB.Where("tmpl_id=? and id=?", tmplId, statusId).First(&status).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return &status, nil
}

func (s *Caller) GetStatusList(tmplId int) ([]model.StatusModel, error) {
	var statusList []model.StatusModel
	if err := global.GVA_DB.Where("tmpl_id=?", tmplId).Find(&statusList).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return statusList, nil
}

func (s *Caller) GetStep(tmplId int, stepId int) (*model.StepModel, error) {
	var step model.StepModel
	if err := global.GVA_DB.Where("tmpl_id=? and id=?", tmplId, stepId).First(&step).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return &step, nil
}

func (s *Caller) GetStepList(tmplId int, statusId int) ([]model.StepModel, error) {
	var stepList []model.StepModel
	if err := global.GVA_DB.Where("start_status_id=?", statusId).Find(&stepList).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return stepList, nil
}

func (s *Caller) GetStepScreen(tmplId int, stepId int) (screenVoList []types.ScreenVo, err error) {
	var screens []ScreenCommon
	if err := global.GVA_DB.Model(model.StepScreenModel{}).Where("tmpl_id=? and step_id=?", tmplId, stepId).Find(&screens).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	var coordinate model.StepScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and step_id=?", tmplId, stepId).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var screensMap = make(map[string]ScreenCommon)
		for _, vo := range screens {
			screensMap[vo.FieldKey] = vo
		}

		var sort []ScreenCommon
		for _, key := range coordinateSlice {
			sort = append(sort, screensMap[key])
			delete(screensMap, key)
		}
		for _, vo := range screensMap {
			sort = append(sort, vo)
		}
		screenVoList = s.screenFormat(tmplId, sort)
		return
	}
	screenVoList = s.screenFormat(tmplId, screens)
	return
}

func (s *Caller) GetStepLimiter(tmplId int, stepId int) ([]model.StepLimiterModel, error) {
	var limits []model.StepLimiterModel
	if err := global.GVA_DB.Model(model.StepLimiterModel{}).Where("tmpl_id=? and step_id=?", tmplId, stepId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return limits, nil
}
