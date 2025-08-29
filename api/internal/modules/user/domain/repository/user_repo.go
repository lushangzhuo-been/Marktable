package repository

import (
	"mark3/internal/modules/user/domain/entity"
)

// WsRepository 空间仓储接口(领域层定义)
type UserRepository interface {
	Register(user *entity.UserModel) error
	FindUserByUserName(username string) entity.UserModel
	// ... 其他仓储接口
}
