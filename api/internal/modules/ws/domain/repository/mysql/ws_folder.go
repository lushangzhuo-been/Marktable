package mysql

import (
	"gorm.io/gorm"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
)

type wsFolderRepository struct {
	db *gorm.DB
}

func NewWsFileRepository(db *gorm.DB) repository.WsFolderRepository {
	return &wsFolderRepository{db: db}
}

func (r *wsFolderRepository) CreateWsFileAndTmplMergedCoordinate(wsFileAndTmplMergedCoordinate *entity.WsFileAndTmplMergedCoordinateModel) error {
	return r.db.Create(wsFileAndTmplMergedCoordinate).Error
}

func (r *wsFolderRepository) CreateWsFolderAndTmplMerged(wsFileAndTmplMerged *entity.WsFileAndTmplMergedModel) error {
	return r.db.Create(wsFileAndTmplMerged).Error
}
