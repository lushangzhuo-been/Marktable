package service

import (
	"errors"
	"mark3/internal/modules/tmpl/application/enum"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
	"mark3/pkg/common"
)

// ScreenService 领域服务（纯领域逻辑）
type ScreenService struct {
	repo repository.ScreenRepository
}

func NewScreenService(repo repository.ScreenRepository) *ScreenService {
	return &ScreenService{repo: repo}
}

func (s *ScreenService) CreateScreens(wsId int, tmplId int, defaultFieldsL3 []entity.FieldModel) error {
	//初始化界面
	for module := range enum.ScreenModuleMap {
		var defaultScreenList []entity.ScreenModel
		for _, field := range enum.FieldLevel2List {
			var screen = entity.ScreenModel{
				WsId:      wsId,
				TmplId:    tmplId,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: enum.ScreenCanModifyYes,
				Required:  enum.ScreenRequiredYes,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		for _, field := range defaultFieldsL3 {
			var screen = entity.ScreenModel{
				WsId:      wsId,
				TmplId:    tmplId,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: enum.ScreenCanModifyYes,
				Required:  enum.ScreenRequiredNo,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		if err := s.repo.CreateScreens(defaultScreenList).Error; err != nil {
			return errors.New("初始化界面失败")
		}
	}
	return nil
}
