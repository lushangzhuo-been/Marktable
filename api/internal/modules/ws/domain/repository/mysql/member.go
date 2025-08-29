package mysql

import (
	"gorm.io/gorm"
	"mark3/internal/modules/ws/domain/entity"
	"mark3/internal/modules/ws/domain/repository"
)

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) repository.MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) CreateMember(member *entity.MemberModel) error {
	return r.db.Create(member).Error
}
