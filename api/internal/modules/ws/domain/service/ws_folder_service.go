package service

import (
	"errors"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
	"mark3/pkg/common"
)

// WsFileService 领域服务（纯领域逻辑）
type WsFolderService struct {
	repo repository.WsFolderRepository
}

func NewWsFolderService(repo repository.WsFolderRepository) *WsFolderService {
	return &WsFolderService{repo: repo}
}

func (s *WsFolderService) CreateWsFolder(wdsId int, coordinate string) (resp *entity.WsFileAndTmplMergedCoordinateModel, err error) {
	wsFile := entity.WsFileAndTmplMergedCoordinateModel{
		WsId:       wdsId,
		Coordinate: coordinate,
		CreatedAt:  common.GetCurrentTime(),
		UpdatedAt:  common.GetCurrentTime(),
	}
	if err := s.repo.CreateWsFileAndTmplMergedCoordinate(&wsFile).Error; err != nil {
		return nil, errors.New("空间排序创建失败")
	}
	return &wsFile, err
}

func (s *WsFolderService) CreateWsFolderAndTmplMerged(wsFileAndTmplMerged entity.WsFileAndTmplMergedModel) (resp *entity.WsFileAndTmplMergedModel, err error) {
	wsFileAndTmplMerged.CreatedAt = common.GetCurrentTime()
	wsFileAndTmplMerged.UpdatedAt = common.GetCurrentTime()

	if err := s.repo.CreateWsFolderAndTmplMerged(&wsFileAndTmplMerged).Error; err != nil {
		return nil, errors.New("空间排序创建失败")
	}
	return &wsFileAndTmplMerged, err
}
