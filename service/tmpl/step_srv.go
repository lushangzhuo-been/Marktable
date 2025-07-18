package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"

	"gorm.io/gorm"
)

type StepSrv struct{}

func (s *StepSrv) Info(req types.StepInfoReq) (resp interface{}, err error) {
	var step model.StepModel
	if err = global.GVA_DB.Where("id=?", req.Id).First(&step).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("步骤不存在")
	}
	type StepVo struct {
		Id               int    `json:"id"`
		WsId             int    `json:"ws_id"`
		TmplId           int    `json:"tmpl_id"`
		StartStatusId    int    `json:"start_status_id"`
		StartStatusName  string `json:"start_status_name"`
		StartStatusColor string `json:"start_status_color"`
		EndStatusId      int    `json:"end_status_id"`
		EndStatusName    string `json:"end_status_name"`
		EndStatusColor   string `json:"end_status_color"`
	}
	var stepVo = StepVo{
		Id:            step.Id,
		WsId:          step.WsId,
		TmplId:        step.TmplId,
		StartStatusId: step.StartStatusId,
		EndStatusId:   step.EndStatusId,
	}
	var startStatus model.StatusModel
	if err = global.GVA_DB.Where("id=?", step.StartStatusId).First(&startStatus).Error; err == nil {
		stepVo.StartStatusName = startStatus.Name
		stepVo.StartStatusColor = startStatus.Color
	}
	var endStatus model.StatusModel
	if err = global.GVA_DB.Where("id=?", step.EndStatusId).First(&endStatus).Error; err == nil {
		stepVo.EndStatusName = endStatus.Name
		stepVo.EndStatusColor = endStatus.Color
	}
	return stepVo, nil
}

func (s *StepSrv) Create(req types.StepCreateReq) (resp interface{}, err error) {
	var step model.StepModel
	if err = global.GVA_DB.Where("tmpl_id=? and start_status_id=? and end_status_id=?", req.TmplId, req.StartStatusId, req.EndStatusId).First(&step).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理找不到记录的情况
			if err = global.GVA_DB.Create(&model.StepModel{
				WsId:          req.WsId,
				TmplId:        req.TmplId,
				StartStatusId: req.StartStatusId,
				EndStatusId:   req.EndStatusId,
				CreatedAt:     common.GetCurrentTime(),
				UpdatedAt:     common.GetCurrentTime(),
			}).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
				return nil, errors.New("操作失败")
			}
		} else {
			//记录错误
			global.GVA_LOG.Error(err.Error())
			return nil, err
		}
	} else {
		return nil, errors.New("步骤已存在")
	}
	return
}

func (s *StepSrv) Delete(req types.StepDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Delete(&model.StepModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
