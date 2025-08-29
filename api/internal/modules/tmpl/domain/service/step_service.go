package service

import (
	"errors"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
	"mark3/pkg/common"
)

// StepService 领域服务（纯领域逻辑）
type StepService struct {
	repo repository.StepRepository
}

func NewStepService(repo repository.StepRepository) *StepService {
	return &StepService{repo: repo}
}

func (s *StepService) CreateSteps(wsId int, tmplId int, statusIdList []int) error {
	//初始化步骤
	var stepList []entity.StepModel
	for i := 0; i < len(statusIdList); i++ {
		for j := 0; j < len(statusIdList); j++ {
			if i != j {
				stepList = append(stepList, entity.StepModel{
					WsId:          wsId,
					TmplId:        tmplId,
					StartStatusId: statusIdList[i],
					EndStatusId:   statusIdList[j],
					CreatedAt:     common.GetCurrentTime(),
					UpdatedAt:     common.GetCurrentTime(),
				})
			}
		}
	}
	if err := s.repo.CreateSteps(stepList).Error; err != nil {
		return errors.New("初始化状态失败")
	}
	return nil
}
