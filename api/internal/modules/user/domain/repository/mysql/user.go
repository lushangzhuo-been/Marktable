package mysql

import (
	"errors"
	"gorm.io/gorm"
	"mark3/global"
	"mark3/internal/modules/user/domain/entity"
	"mark3/internal/modules/user/domain/repository"
	model "mark3/repository/db/model/user"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(user *entity.UserModel) error {
	if err := r.db.Create(user).Error; err != nil {
		r.db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return errors.New("操作失败")
	}
	return nil
}

func (r *UserRepository) FindUserByUserName(username string) (model.UserModel, error) {
	var user model.UserModel
	r.db.Model(&model.UserModel{}).Where("username=?", username).First(user)
	return user, nil
}
