package service

import (
	"errors"
	"mark3/internal/modules/ws/application/dto"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
	"mark3/pkg/common"
)

// WsService 领域服务（纯领域逻辑）
type WsService struct {
	repo repository.WsRepository
}

func NewWsService(repo repository.WsRepository) *WsService {
	return &WsService{repo: repo}
}

func (s *WsService) CreateWs(userid int, req dto.WsCreateReq) (resp *entity.WsModel, err error) {
	ws := entity.WsModel{
		Name:      req.Name,
		Desc:      req.Desc,
		Creator:   userid,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}
	if err := s.repo.CreateWs(&ws).Error; err != nil {
		return nil, errors.New("空间创建失败")
	}
	return &ws, err
}
