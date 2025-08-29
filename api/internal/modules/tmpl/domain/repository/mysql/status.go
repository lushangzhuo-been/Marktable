package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
)

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) repository.StatusRepository {
	return &statusRepository{db: db}
}

func (r *statusRepository) CreateStatus(status *entity.StatusModel) error {
	if err := r.db.Create(status).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
