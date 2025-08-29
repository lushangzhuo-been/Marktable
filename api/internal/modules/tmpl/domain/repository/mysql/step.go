package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
)

type stepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) repository.StepRepository {
	return &stepRepository{db: db}
}

func (r *stepRepository) CreateSteps(Steps []entity.StepModel) error {
	if err := r.db.Create(&Steps).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
