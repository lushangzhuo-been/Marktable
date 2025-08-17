package service

import (
	"errors"
	"mark3/internal/modules/tmpl/application/enum"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
	"mark3/pkg/common"
)

// FieldService 领域服务（纯领域逻辑）
type FieldService struct {
	repo repository.FieldRepository
}

func NewFieldService(repo repository.FieldRepository) *FieldService {
	return &FieldService{repo: repo}
}

func (s *FieldService) CreateFields(wsId int, tmplId int) (resp []entity.FieldModel, err error) {
	//初始化默认字段
	var defaultFields []entity.FieldModel
	for _, field := range append(enum.FieldLevel2List, enum.FieldLevel1List...) {
		field.WsId = wsId
		field.TmplId = tmplId
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
	}
	// 初始化字段level3： 优先级（高，中，低），截止日期，描述（textarea_com）
	var defaultFieldsL3 []entity.FieldModel
	for _, field := range enum.FieldLevel3List {
		field.WsId = wsId
		field.TmplId = tmplId
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
		defaultFieldsL3 = append(defaultFieldsL3, field)
	}

	if err := s.repo.CreateFields(defaultFields).Error; err != nil {
		return nil, errors.New("流程创建失败")
	}
	return defaultFieldsL3, err
}
