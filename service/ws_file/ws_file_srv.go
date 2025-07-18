package ws_file

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	userModel "mark3/repository/db/model/user"
	model "mark3/repository/db/model/ws_file"
	types "mark3/types/ws_file"
	"strconv"
	"strings"
)

type WsFileSrv struct{}

func (s *WsFileSrv) Info(req types.WsFileInfoReq) (resp model.WsFileModel, err error) {
	var wsFile model.WsFileModel
	if err = global.GVA_DB.Where("id=? and ws_id=?", req.Id, req.WsId).First(&wsFile).Error; err != nil {
		return wsFile, errors.New("参数错误，文件夹不存在")
	}
	return wsFile, nil
}

func (s *WsFileSrv) InfoView(req types.WsFileInfoReq) (resp interface{}, err error) {
	type WsFileVo struct {
		Id              int    `json:"id"`
		WsId            int    `json:"ws_id"`
		WsFileId        int    `json:"ws_file_id"`
		Name            string `json:"name"`
		Desc            string `json:"desc"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		CreatedUserId   int    `json:"created_user_id"`
		UpdatedUserId   int    `json:"updated_user_id"`
		CreatedFullName string `json:"created_full_name"`
		CreatedUsername string `json:"created_username"`
		UpdatedFullName string `json:"updated_full_name"`
		UpdatedUsername string `json:"updated_username"`
	}
	var wsFileVo WsFileVo
	info, err := s.Info(req)
	if err != nil {
		return nil, err
	}
	wsFileVo.Id = info.Id
	wsFileVo.WsId = info.WsId
	wsFileVo.Name = info.Name
	wsFileVo.Desc = info.Desc
	wsFileVo.CreatedAt = info.CreatedAt
	wsFileVo.UpdatedAt = info.UpdatedAt
	wsFileVo.CreatedUserId = info.CreatedUserId
	wsFileVo.UpdatedUserId = info.UpdatedUserId
	if info.CreatedUserId > 0 {
		var user userModel.UserModel
		if err := global.GVA_DB.Where("id=?", info.CreatedUserId).Find(&user).Error; err != nil {
			return wsFileVo, nil
		}
		wsFileVo.CreatedFullName = user.FullName
		wsFileVo.CreatedUsername = user.Username
	}
	if info.UpdatedUserId > 0 {
		var user userModel.UserModel
		if err := global.GVA_DB.Where("id=?", info.UpdatedUserId).Find(&user).Error; err != nil {
			return wsFileVo, nil
		}
		wsFileVo.UpdatedFullName = user.FullName
		wsFileVo.UpdatedUsername = user.Username
	}
	return wsFileVo, nil
}

func (s *WsFileSrv) InfoByName(WsId int, Name string) (resp model.WsFileModel, err error) {
	var wsFile model.WsFileModel
	if err = global.GVA_DB.Where("ws_id=? and name=?", WsId, Name).First(&wsFile).Error; err != nil {
		return wsFile, errors.New("参数错误，文件夹不存在")
	}
	return wsFile, nil
}

func (s *WsFileSrv) List(WsId int) (resp []model.WsFileModel, err error) {
	var wsFiles []model.WsFileModel
	if err = global.GVA_DB.Where("ws_id=?", WsId).Find(&wsFiles).Error; err != nil {
		return wsFiles, errors.New("参数错误，文件夹不存在")
	}
	return wsFiles, nil
}

func (s *WsFileSrv) AddTmplId(tmplId int, req types.WsFileInfoReq) (resp interface{}, err error) {
	wsFile, err := s.Info(req)
	if err != nil {
		return nil, err
	}
	var Coordinate string
	if len(wsFile.Coordinate) == 0 {
		Coordinate = "," + strconv.Itoa(tmplId) + ","
	} else {
		Coordinate = "," + strconv.Itoa(tmplId) + wsFile.Coordinate
	}
	wsFile.Coordinate = Coordinate
	wsFile.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&model.WsFileModel{}).Where("ws_id=? and id=?", req.WsId, req.Id).Save(wsFile).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	wsFile.Coordinate = Coordinate
	return wsFile, nil
}

func (s *WsFileSrv) RemoveTmplId(tmplId int, req types.WsFileInfoReq) (resp interface{}, err error) {
	wsFile, err := s.Info(req)
	if err != nil {
		return nil, err
	}
	subId := "," + strconv.Itoa(tmplId) + ","
	if wsFile.Coordinate != subId && !strings.Contains(wsFile.Coordinate, subId) {
		return nil, nil
	}
	Coordinate := strings.Replace(wsFile.Coordinate, ","+strconv.Itoa(tmplId)+",", ",", -1)
	if len(Coordinate) == 1 && Coordinate == "," {
		Coordinate = ""
	}
	wsFile.Coordinate = Coordinate
	wsFile.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&model.WsFileModel{}).Where("ws_id=? and id=?", req.WsId, req.Id).Save(wsFile).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return wsFile, nil
}

func (s *WsFileSrv) UpdateWsFileName(userid int, req types.UpdateWsFileName) (resp interface {
}, err error) {
	wsFile, err := s.InfoByName(req.WsId, req.Name)
	if err == nil && wsFile.Id != req.Id {
		return nil, errors.New("操作失败，该空间下此文件夹已存在")
	}
	wsFile.Name = req.Name
	wsFile.Desc = req.Desc
	wsFile.UpdatedUserId = userid
	wsFile.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&model.WsFileModel{}).Where("ws_id=? and id=?", req.WsId, req.Id).Save(wsFile).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return wsFile, nil
}

func (s *WsFileSrv) UpdateWsFileCoordinate(req types.UpdateWsFileCoordinate) (resp interface {
}, err error) {
	wsFile, err := s.Info(req.WsFileInfoReq)
	if err != nil {
		return nil, err
	}
	wsFile.Coordinate = req.Coordinate
	wsFile.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&model.WsFileModel{}).Where("ws_id=? and id=?", req.WsId, req.Id).Save(wsFile).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return wsFile, nil
}
