package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
)

type wsRepository struct {
	db *gorm.DB
}

func NewWsRepository(db *gorm.DB) repository.WsRepository {
	return &wsRepository{db: db}
}

func (r *wsRepository) CreateWs(ws *entity.WsModel) error {
	if err := r.db.Create(ws).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
