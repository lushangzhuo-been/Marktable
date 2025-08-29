package service

import (
	"errors"
	"mark3/internal/modules/tmpl/application/dto"
	"mark3/internal/modules/tmpl/domain/entity"
	"mark3/internal/modules/tmpl/domain/repository"
	"mark3/pkg/common"
)

// TmplService 领域服务（纯领域逻辑）
type TmplService struct {
	repo repository.TmplRepository
}

func NewTmplService(repo repository.TmplRepository) *TmplService {
	return &TmplService{repo: repo}
}

func (s *TmplService) CreateTmpl(userid int, req dto.TmplCreateReq) (resp *entity.TmplModel, err error) {
	//创建模板
	var tmpl = entity.TmplModel{
		WsId:          req.WsId,
		WsFileId:      req.WsFileId,
		CreatedUserId: userid,
		UpdatedUserId: userid,
		Name:          req.Name,
		Desc:          req.Desc,
		Mode:          req.Mode,
		CreatedAt:     common.GetCurrentTime(),
		UpdatedAt:     common.GetCurrentTime(),
	}

	if err := s.repo.CreateTmpl(&tmpl).Error; err != nil {
		return nil, errors.New("流程创建失败")
	}
	return &tmpl, err
}
