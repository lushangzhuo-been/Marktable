package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
)

type tmplRepository struct {
	db *gorm.DB
}

func NewTmplRepository(db *gorm.DB) repository.TmplRepository {
	return &tmplRepository{db: db}
}

func (r *tmplRepository) CreateTmpl(tmpl *entity.TmplModel) error {
	if err := r.db.Create(tmpl).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}
