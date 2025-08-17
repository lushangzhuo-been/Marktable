package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
)

type screenRepository struct {
	db *gorm.DB
}

func NewScreenRepository(db *gorm.DB) repository.ScreenRepository {
	return &screenRepository{db: db}
}

func (r *screenRepository) CreateScreens(screens []entity.ScreenModel) error {
	if err := r.db.Create(&screens).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
