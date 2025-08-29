package service

import (
	"errors"
	"mark3/global"
	"mark3/internal/modules/user/application/dto"
	"mark3/internal/modules/user/domain/entity"
	"mark3/internal/modules/user/domain/repository"
	"mark3/pkg/common"
)

// UserService 领域服务（纯领域逻辑）
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req dto.RegisterReq) (resp *entity.UserModel, err error) {
	user := s.repo.FindUserByUserName(req.Username)
	if user.Id != 0 {
		return nil, errors.New("用户已存在")
	}
	hashPassword, err := common.EncryptPassword(req.Password)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("加密失败")
	}
	user.Username = req.Username
	user.Password = string(hashPassword)
	user.FullName = req.Username
	user.Email = req.Email
	user.CreatedAt = common.GetCurrentTime()
	user.UpdatedAt = common.GetCurrentTime()
	err = s.repo.Register(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
