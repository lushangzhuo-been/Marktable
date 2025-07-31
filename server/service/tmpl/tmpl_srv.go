package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	wsEnum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/tmpl"
	userModel "mark3/repository/db/model/user"
	wsModel "mark3/repository/db/model/ws"
	wsfilemodel "mark3/repository/db/model/ws_file"
	wsfilesrv "mark3/service/ws_file"
	wsfileandtmplmergedsrv "mark3/service/ws_file_and_tmpl_merged"
	types "mark3/types/tmpl"
	wsfiletypes "mark3/types/ws_file"
	"strconv"
	"strings"
)

type TmplSrv struct{}

func (s *TmplSrv) Config() (resp interface{}, err error) {
	var config types.TmplConfigResp
	config.ModeConfig = common.FormatConfig(enum.TmplModeMap)
	return config, nil
}

func (s *TmplSrv) List(req types.TmplListReq) (resp []model.TmplModel, err error) {
	var tmpls []model.TmplModel
	global.GVA_DB.Where("ws_id=?", req.WsId).Find(&tmpls)
	return tmpls, nil
}

func (s *TmplSrv) ListNotInWsFile(req types.WsFileTmplListReq) (resp []wsfiletypes.TmplVo, err error) {
	var tmpls []model.TmplModel
	var TmplVoList []wsfiletypes.TmplVo
	if err := global.GVA_DB.Where("ws_id=? and ws_file_id=0", req.WsId).Find(&tmpls).Error; err != nil {
		return TmplVoList, nil
	}

	if len(tmpls) == 0 {
		return TmplVoList, nil
	}
	for _, tmpl := range tmpls {
		TmplVoList = append(TmplVoList, wsfiletypes.TmplVo{
			Id:            tmpl.Id,
			WsId:          tmpl.WsId,
			WsFileId:      tmpl.WsFileId,
			ImplTableName: "tmpl",
			Name:          tmpl.Name,
			Desc:          tmpl.Desc,
			Mode:          tmpl.Mode,
			CreatedAt:     tmpl.CreatedAt,
			UpdatedAt:     tmpl.UpdatedAt,
		})

	}
	TmplVoList, err = s.getTmplIsAdminAndIsMember(req.UserId, req.WsId, TmplVoList)
	if err != nil {
		return nil, err
	}
	return TmplVoList, nil
}

func (s *TmplSrv) ListByIds(req types.TmplListReq) (resp []model.TmplModel, err error) {
	var tmpls []model.TmplModel
	global.GVA_DB.Where("id in ?", req.WsId).Find(&tmpls)
	return tmpls, nil
}

func (s *TmplSrv) WsFileTmplList(req types.WsFileTmplListReq) (resp []wsfiletypes.TmplVo, err error) {
	var TmplList []model.TmplModel
	var TmplVoList []wsfiletypes.TmplVo
	if err := global.GVA_DB.Where("ws_id=? and ws_file_id=?", req.WsId, req.WsFileId).Find(&TmplList).Error; err != nil {
		return TmplVoList, nil
	}
	if len(TmplList) == 0 {
		return TmplVoList, nil
	}
	WsFileSrv := new(wsfilesrv.WsFileSrv)
	WsFile, err := WsFileSrv.Info(wsfiletypes.WsFileInfoReq{
		WsId: req.WsId,
		Id:   req.WsFileId,
	})
	if err != nil {
		return TmplVoList, err
	}
	if len(WsFile.Coordinate) == 0 {
		return TmplVoList, nil
	}
	tmplMap := make(map[int]model.TmplModel)
	for _, v := range TmplList {
		tmplMap[v.Id] = v
	}
	Coordinates := strings.Split(WsFile.Coordinate, ",")
	for _, v := range Coordinates {
		if len(v) == 0 {
			continue
		}
		TmplId, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		tmpl, ok := tmplMap[TmplId]
		if !ok {
			continue
		}
		TmplVoList = append(TmplVoList, wsfiletypes.TmplVo{
			Id:            tmpl.Id,
			WsId:          tmpl.WsId,
			WsFileId:      tmpl.WsFileId,
			ImplTableName: "tmpl",
			Name:          tmpl.Name,
			Desc:          tmpl.Desc,
			Mode:          tmpl.Mode,
			CreatedAt:     tmpl.CreatedAt,
			UpdatedAt:     tmpl.UpdatedAt,
		})

	}
	TmplVoList, err = s.getTmplIsAdminAndIsMember(req.UserId, req.WsId, TmplVoList)
	if err != nil {
		return nil, err
	}
	return TmplVoList, nil
}

func (s *TmplSrv) getTmplIsAdminAndIsMember(userid int, wsId int, TmplVoList []wsfiletypes.TmplVo) (resp []wsfiletypes.TmplVo, err error) {
	var member wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=? and status=? and role=?", wsId, userid, wsEnum.WsStatusOk, wsEnum.WsRoleAdmin).First(&member).Error; err == nil {
		for index, vo := range TmplVoList {
			vo.IsMember = "yes"
			vo.IsAdmin = "yes"
			TmplVoList[index] = vo
		}

		return TmplVoList, nil
	}

	for index, vo := range TmplVoList {
		if vo.Mode == enum.TmplModePublic {
			vo.IsMember = "yes"
			vo.IsAdmin = "no"
			TmplVoList[index] = vo
		} else {
			var member model.MemberModel
			if err := global.GVA_DB.Where("tmpl_id=? and userid=? and status=?", vo.Id, userid, enum.TmplStatusOk).First(&member).Error; err != nil {
				vo.IsMember = "no"
			} else {
				vo.IsMember = "yes"
			}

			var role model.RoleModel
			if err := global.GVA_DB.Where("tmpl_id=? and id=? and sign=?", vo.Id, member.RoleId, enum.RoleSignAdmin).First(&role).Error; err != nil {
				vo.IsAdmin = "no"
			} else {
				vo.IsAdmin = "yes"
			}
			TmplVoList[index] = vo
		}
	}
	return TmplVoList, nil
}

func (s *TmplSrv) Info(req types.TmplInfoReq) (resp model.TmplModel, err error) {
	var tmpl model.TmplModel
	if err = global.GVA_DB.Where("id=? and ws_id=?", req.Id, req.WsId).First(&tmpl).Error; err != nil {
		return tmpl, errors.New("参数错误，流程不存在")
	}
	return tmpl, nil
}

func (s *TmplSrv) InfoView(req types.TmplInfoReq) (resp interface{}, err error) {
	type TmplVo struct {
		Id              int    `json:"id"`
		WsId            int    `json:"ws_id"`
		WsFileId        int    `json:"ws_file_id"`
		Name            string `json:"name"`
		Desc            string `json:"desc"`
		Mode            string `json:"mode"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		CreatedUserId   int    `json:"created_user_id"`
		UpdatedUserId   int    `json:"updated_user_id"`
		CreatedFullName string `json:"created_full_name"`
		CreatedUsername string `json:"created_username"`
		UpdatedFullName string `json:"updated_full_name"`
		UpdatedUsername string `json:"updated_username"`
	}
	var tmplVo TmplVo
	info, err := s.Info(req)
	if err != nil {
		return nil, err
	}
	tmplVo.Id = info.Id
	tmplVo.WsId = info.WsId
	tmplVo.WsFileId = info.WsFileId
	tmplVo.Name = info.Name
	tmplVo.Desc = info.Desc
	tmplVo.Mode = info.Mode
	tmplVo.CreatedAt = info.CreatedAt
	tmplVo.UpdatedAt = info.UpdatedAt
	tmplVo.CreatedUserId = info.CreatedUserId
	tmplVo.UpdatedUserId = info.UpdatedUserId
	if info.CreatedUserId > 0 {
		var user userModel.UserModel
		if err := global.GVA_DB.Where("id=?", info.CreatedUserId).Find(&user).Error; err != nil {
			return tmplVo, nil
		}
		tmplVo.CreatedFullName = user.FullName
		tmplVo.CreatedUsername = user.Username
	}
	if info.UpdatedUserId > 0 {
		var user userModel.UserModel
		if err := global.GVA_DB.Where("id=?", info.UpdatedUserId).Find(&user).Error; err != nil {
			return tmplVo, nil
		}
		tmplVo.UpdatedFullName = user.FullName
		tmplVo.UpdatedUsername = user.Username
	}
	return tmplVo, nil
}

func (s *TmplSrv) Create(userid int, req types.TmplCreateReq) (resp interface{}, err error) {
	tx := global.GVA_DB.Begin()
	//创建模板
	var tmpl = model.TmplModel{
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
	if err = tx.Create(&tmpl).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 空间文件夹下创建流程，需要将流程ID加入到空间文件夹表
	if req.WsFileId > 0 {
		WsFileSrv := new(wsfilesrv.WsFileSrv)
		_, err := WsFileSrv.AddTmplId(tmpl.Id, wsfiletypes.WsFileInfoReq{
			WsId: req.WsId,
			Id:   req.WsFileId,
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 流程和文件合并表里面添加ID
		wsFileAndTmplMergedSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
		wsFileAndTmplMerged, err := wsFileAndTmplMergedSrv.AddWsFileAndTmplMerged(wsfilemodel.WsFileAndTmplMergedModel{
			WsId:          req.WsId,
			ImplTableName: "tmpl",
			ImplTableId:   tmpl.Id,
		})
		if err != nil {
			return nil, err
		}
		// 将流程和文件合并表ID添加到排序表
		wsFileAndTmplMergedCoordinateSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedCoordinateSrv)
		_, err = wsFileAndTmplMergedCoordinateSrv.AddWsFileAndTmplMergedCoordinate(wsFileAndTmplMerged)
		if err != nil {
			return nil, err
		}
	}

	//初始化默认字段
	var defaultFields []model.FieldModel
	for _, field := range append(enum.FieldLevel2List, enum.FieldLevel1List...) {
		field.WsId = req.WsId
		field.TmplId = tmpl.Id
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
	}
	// 初始化字段level3： 优先级（高，中，低），截止日期，描述（textarea_com）
	var defaultFieldsL3 []model.FieldModel
	for _, field := range enum.FieldLevel3List {
		field.WsId = req.WsId
		field.TmplId = tmpl.Id
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
		defaultFieldsL3 = append(defaultFieldsL3, field)
	}

	if err = tx.Create(&defaultFields).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return
	}

	//初始化界面
	for module := range enum.ScreenModuleMap {
		var defaultScreenList []model.ScreenModel
		for _, field := range enum.FieldLevel2List {
			var screen = model.ScreenModel{
				WsId:      req.WsId,
				TmplId:    tmpl.Id,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: enum.ScreenCanModifyYes,
				Required:  enum.ScreenRequiredYes,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		for _, field := range defaultFieldsL3 {
			var screen = model.ScreenModel{
				WsId:      req.WsId,
				TmplId:    tmpl.Id,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: enum.ScreenCanModifyYes,
				Required:  enum.ScreenRequiredNo,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		if err = tx.Create(&defaultScreenList).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return
		}
	}

	//初始化状态
	var statusIdList []int
	for _, v := range enum.InitStatusList {
		var status = model.StatusModel{
			WsId:      req.WsId,
			TmplId:    tmpl.Id,
			Name:      v.Name,
			Color:     v.Color,
			Mode:      v.Mode,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}
		if err = tx.Create(&status).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return
		}
		statusIdList = append(statusIdList, status.Id)
	}
	//初始化步骤
	var stepList []model.StepModel
	for i := 0; i < len(statusIdList); i++ {
		for j := 0; j < len(statusIdList); j++ {
			if i != j {
				stepList = append(stepList, model.StepModel{
					WsId:          req.WsId,
					TmplId:        tmpl.Id,
					StartStatusId: statusIdList[i],
					EndStatusId:   statusIdList[j],
					CreatedAt:     common.GetCurrentTime(),
					UpdatedAt:     common.GetCurrentTime(),
				})
			}
		}
	}
	if err = tx.Create(&stepList).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return
	}

	//私有模板初始化角色、成员
	if req.Mode == enum.TmplModePrivate {
		var role = model.RoleModel{
			WsId:      req.WsId,
			TmplId:    tmpl.Id,
			Name:      "管理员",
			Desc:      "管理员",
			Sign:      enum.RoleSignAdmin,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}
		if err = tx.Create(&role).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		if err = tx.Create(&model.MemberModel{
			WsId:      req.WsId,
			TmplId:    tmpl.Id,
			Userid:    userid,
			RoleId:    role.Id,
			Status:    enum.TmplStatusOk,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}
	tx.Commit()
	return
}

func (s *TmplSrv) Update(userid int, req types.TmplUpdateReq) (resp interface{}, err error) {
	tx := global.GVA_DB.Begin()

	var tmpl model.TmplModel
	if err = tx.Where("id=? and ws_id=?", req.Id, req.WsId).First(&tmpl).Error; err != nil {
		tx.Rollback()
		return
	}

	if tmpl.Mode == enum.TmplModePublic && req.Mode == enum.TmplModePrivate {
		var role = model.RoleModel{
			WsId:      req.WsId,
			TmplId:    tmpl.Id,
			Name:      "管理员",
			Desc:      "管理员",
			Sign:      enum.RoleSignAdmin,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}
		if err = tx.Create(&role).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		if err = tx.Create(&model.MemberModel{
			WsId:      req.WsId,
			TmplId:    tmpl.Id,
			Userid:    userid,
			RoleId:    role.Id,
			Status:    enum.TmplStatusOk,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else if tmpl.Mode == enum.TmplModePrivate && req.Mode == enum.TmplModePublic {
		if err = tx.Where("ws_id=? and tmpl_id=?", req.WsId, req.Id).Delete(&model.RoleModel{}).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		if err = tx.Where("ws_id=? and tmpl_id=?", req.WsId, req.Id).Delete(&model.MemberModel{}).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}

	tmpl.Name = req.Name
	tmpl.Desc = req.Desc
	tmpl.Mode = req.Mode
	tmpl.UpdatedUserId = userid
	tmpl.UpdatedAt = common.GetCurrentTime()

	if err = tx.Save(&tmpl).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	tx.Commit()
	return
}

func (s *TmplSrv) Delete(req types.TmplDeleteReq) (resp interface{}, err error) {
	info, err := s.Info(types.TmplInfoReq{
		WsId: req.WsId,
		Id:   req.Id,
	})
	if err != nil {
		return nil, err
	}

	if err = global.GVA_DB.Where("ws_id=? and id=?", req.WsId, req.Id).Delete(&model.TmplModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	// 空间文件夹下创建流程，需要将流程ID加入到空间文件夹表
	if info.WsFileId > 0 {
		WsFileSrv := new(wsfilesrv.WsFileSrv)
		_, err := WsFileSrv.RemoveTmplId(req.Id, wsfiletypes.WsFileInfoReq{
			WsId: req.WsId,
			Id:   info.WsFileId,
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 空间下的排序ID去除此文件夹ID 从合并表里面移除ID
		wsFileAndTmplMergedSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
		WsFileAndTmplMergedModel, err := wsFileAndTmplMergedSrv.GetWsFileAndTmplMergedModel(req.WsId, "tmpl", req.Id)
		if err != nil {
			return nil, err
		}

		wsFileAndTmplMergedCoordinateSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedCoordinateSrv)
		_, err = wsFileAndTmplMergedCoordinateSrv.RemoveWsFileAndTmplMergedCoordinate(WsFileAndTmplMergedModel)
		if err != nil {
			return nil, err
		}

		_, err = wsFileAndTmplMergedSrv.DeleteWsFileAndTmplMerged(WsFileAndTmplMergedModel.Id)
		if err != nil {
			return nil, err
		}
	}

	return
}

func (s *TmplSrv) UpdateWsFileId(req wsfiletypes.TmplMoveInWsFile) (resp interface{}, err error) {
	info, err := s.Info(types.TmplInfoReq{
		WsId: req.WsId,
		Id:   req.TmplId,
	})
	if err != nil {
		return nil, err
	}
	info.WsFileId = req.Id
	if err = global.GVA_DB.Model(model.TmplModel{}).Where("id=?", req.TmplId).Save(info).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return false, errors.New("操作失败")
	}
	return true, nil
}
