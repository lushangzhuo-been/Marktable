package service

import (
	"mark3/internal/modules/tmpl/application/enum"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
	"mark3/pkg/common"
)

// StatusService 领域服务（纯领域逻辑）
type StatusService struct {
	repo repository.StatusRepository
}

func NewStatusService(repo repository.StatusRepository) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) CreateStatus(wsId int, tmplId int) (resp []int, err error) {
	//初始化状态
	var statusIdList []int
	for _, v := range enum.InitStatusList {
		var status = entity.StatusModel{
			WsId:      wsId,
			TmplId:    tmplId,
			Name:      v.Name,
			Color:     v.Color,
			Mode:      v.Mode,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}
		err := s.repo.CreateStatus(&status)
		if err != nil {
			return nil, err
		}
		statusIdList = append(statusIdList, status.Id)
	}
	return statusIdList, nil
}
