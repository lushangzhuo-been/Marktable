package ws

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/ws"
	wsFileModel "mark3/repository/db/model/ws_file"
	types "mark3/types/ws"
)

type WsSrv struct{}

func (s *WsSrv) Info(req types.WsInfoReq) (resp interface{}, err error) {
	var ws types.WsVo
	if err = global.GVA_DB.Model(&model.WsModel{}).Where("id=?", req.Id).First(&ws).Error; err != nil {
		return
	}

	var admins []types.MemberVo
	global.GVA_DB.Raw("select a.id, a.userid, b.username, b.full_name, b.avatar from ws_member as a inner join user as b on a.userid=b.id where a.ws_id=? and a.role=? and status=?", req.Id, enum.WsRoleAdmin, enum.WsStatusOk).Scan(&admins)

	var user types.MemberVo
	global.GVA_DB.Raw("select full_name from user where id=?", ws.Creator).Find(&user)

	ws.CreatorName = user.FullName
	ws.Admin = admins

	return ws, nil
}

func (s *WsSrv) Create(userid int, req types.WsCreateReq) (resp interface{}, err error) {
	tx := global.GVA_DB.Begin()

	ws := model.WsModel{
		Name:      req.Name,
		Desc:      req.Desc,
		Creator:   userid,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}
	if err = tx.Create(&ws).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	if err = tx.Create(&model.MemberModel{
		WsId:      ws.Id,
		Userid:    userid,
		Role:      enum.WsRoleAdmin,
		Status:    enum.WsStatusOk,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 创建空间文件夹、流程合并排序数据
	if err = tx.Create(&wsFileModel.WsFileAndTmplMergedCoordinateModel{
		WsId:       ws.Id,
		Coordinate: "",
		CreatedAt:  common.GetCurrentTime(),
		UpdatedAt:  common.GetCurrentTime(),
	}).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	tx.Commit()
	return
}

func (s *WsSrv) Update(req types.WsUpdateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Updates(&model.WsModel{
		Name: req.Name,
		Desc: req.Desc,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *WsSrv) Delete(req types.WsDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", req.Id).Delete(&model.WsModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
