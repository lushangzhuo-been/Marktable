package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
)

type fieldRepository struct {
	db *gorm.DB
}

func NewFieldRepository(db *gorm.DB) repository.FieldRepository {
	return &fieldRepository{db: db}
}

func (r *fieldRepository) CreateFields(fields []entity.FieldModel) error {
	if err := r.db.Create(&fields).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
