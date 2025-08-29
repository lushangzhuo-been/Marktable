package service

import (
	"errors"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/ws"
)

// MemberService 领域服务（纯领域逻辑）
type MemberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) CreateMember(userid int, wdsId int) (resp *entity.MemberModel, err error) {
	member := entity.MemberModel{
		WsId:      wdsId,
		Userid:    userid,
		Role:      enum.WsRoleAdmin,
		Status:    enum.WsStatusOk,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}
	if err := s.repo.CreateMember(&member).Error; err != nil {
		return nil, errors.New("空间记忆创建失败")
	}
	return &member, err
}
