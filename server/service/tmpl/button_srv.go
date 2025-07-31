package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"
)

type ButtonSrv struct{}

func (s *ButtonSrv) Info(req types.ButtonInfoReq) (resp interface{}, err error) {
	var button model.ButtonModel
	if err = global.GVA_DB.Where("tmpl_id=? and node_id=? and id=?", req.TmplId, req.NodeId, req.Id).First(button).Error; err != nil {
		return
	}
	return button, nil
}

func (s *ButtonSrv) Create(req types.ButtonCreateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Create(&model.ButtonModel{
		WsId:         req.WsId,
		TmplId:       req.TmplId,
		NodeId:       req.NodeId,
		Name:         req.Name,
		TargetNodeId: req.TargetNodeId,
		CreatedAt:    common.GetCurrentTime(),
		UpdatedAt:    common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *ButtonSrv) Update(req types.ButtonUpdateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("tmpl_id=? and node_id=? and id=?", req.TmplId, req.NodeId, req.Id).Updates(&model.ButtonModel{
		Name:         req.Name,
		TargetNodeId: req.TargetNodeId,
		UpdatedAt:    common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *ButtonSrv) Delete(req types.ButtonDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("tmpl_id=? and node_id=? and id=?", req.TmplId, req.NodeId, req.Id).Delete(&model.ButtonModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
