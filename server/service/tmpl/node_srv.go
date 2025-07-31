package tmpl

import (
	"errors"
	"fmt"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"
)

type NodeSrv struct{}

func (s *NodeSrv) Config() (resp interface{}, err error) {
	var config types.NodeConfigResp
	config.ModeConfig = common.FormatConfig(enum.NodeModeMap)
	return config, nil
}

func (s *NodeSrv) List(req types.NodeListReq) (resp interface{}, err error) {
	var nodes []model.NodeModel
	global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&nodes)

	nodesMap := make(map[int]model.NodeModel, len(nodes))
	for _, node := range nodes {
		nodesMap[node.Id] = node
	}

	//todo 代码需要重构
	var list []types.NodeVo
	for _, node := range nodes {
		var nodeResp types.NodeVo
		nodeResp.Node = node
		var buttons []types.ButtonVo
		global.GVA_DB.Where("node_id=?", node.Id).Find(&buttons)
		for k, button := range buttons {
			if tmp, ok := nodesMap[button.TargetNodeId]; ok {
				button.TargetNodeName = tmp.Name
			}
			buttons[k] = button
		}
		nodeResp.Buttons = buttons
		list = append(list, nodeResp)
	}
	return list, nil
}

func (s *NodeSrv) Info(req types.NodeInfoReq) (resp interface{}, err error) {
	var node model.NodeModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&node).Error; err != nil {
		return
	}
	return node, nil
}

func (s *NodeSrv) Create(req types.NodeCreateReq) (resp interface{}, err error) {
	//判断节点名称是否存在
	var node model.NodeModel
	global.GVA_DB.Where("tmpl_id=? and name=?", req.TmplId, req.Name).Find(&node)
	if node.Id != 0 {
		return nil, fmt.Errorf("[%s]节点名称已存在,请重新命名", req.Name)
	}

	//判断首节点是否已经存在
	var headNode model.NodeModel
	global.GVA_DB.Where("tmpl_id=? and mode=?", req.TmplId, enum.NodeModeHead).Find(&headNode)
	if headNode.Id != 0 && req.Mode == enum.NodeModeHead {
		return nil, errors.New("首节点已经存在,请重新选择节点类型")
	}
	if err = global.GVA_DB.Create(&model.NodeModel{
		WsId:        req.WsId,
		TmplId:      req.TmplId,
		Name:        req.Name,
		Mode:        req.Mode,
		StatusText:  req.StatusText,
		StatusColor: req.StatusColor,
		CreatedAt:   common.GetCurrentTime(),
		UpdatedAt:   common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *NodeSrv) Update(req types.NodeUpdateReq) (resp interface{}, err error) {
	//判断首节点是否已经存在
	if req.Mode == enum.NodeModeHead {
		var headNode model.NodeModel
		if err = global.GVA_DB.Where("tmpl_id=? and mode=? and id !=?", req.TmplId, enum.NodeModeHead, req.Id).First(&headNode).Error; err == nil {
			return nil, errors.New("首节点已经存在,请重新选择节点类型")
		}
	}

	//判断重命名
	var node model.NodeModel
	if err = global.GVA_DB.Where("tmpl_id=? and id !=? and name=?", req.TmplId, req.Id, req.Name).First(&node).Error; err == nil {
		return nil, fmt.Errorf("[%s]步骤名称已存在,请重新命名", req.Name)
	}

	//判断步骤是否存在
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&node).Error; err != nil {
		return nil, errors.New("参数错误")
	}

	if node.Mode == enum.NodeModeHead && req.Mode != enum.NodeModeHead {
		return nil, errors.New("首节点类型不能更改")
	}

	node.Name = req.Name
	node.Mode = req.Mode
	node.StatusText = req.StatusText
	node.StatusColor = req.StatusColor
	if err = global.GVA_DB.Save(node).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *NodeSrv) GetAllStatus(req types.NodeGetAllStatusesReq) (resp interface{}, err error) {
	var nodes []model.NodeModel
	if err = global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&nodes).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	type Status struct {
		StatusText  string `json:"status_text"`
		StatusColor string `json:"status_color"`
	}
	var allStatuses []Status
	for _, node := range nodes {
		var status Status
		status.StatusText = node.StatusText
		status.StatusColor = node.StatusColor
		allStatuses = append(allStatuses, status)
	}
	return allStatuses, nil
}

func (s *NodeSrv) Delete(req types.NodeDeleteReq) (resp interface{}, err error) {
	var node model.NodeModel
	if err = global.GVA_DB.Where("tmpl_id=? and id =?", req.TmplId, req.Id).First(&node).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("数据不存在")
	}
	if node.Mode == enum.NodeModeHead {
		return nil, errors.New("首节点类型不能删除")
	}

	if err = global.GVA_DB.Where("tmpl_id=? and id =?", req.TmplId, req.Id).Delete(&model.NodeModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("删除失败")
	}
	return
}
