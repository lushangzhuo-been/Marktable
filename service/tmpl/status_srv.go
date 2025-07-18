package tmpl

import (
	"context"
	"errors"
	"fmt"
	"mark3/global"
	"mark3/pkg/common"
	model "mark3/repository/db/model/tmpl"
	types "mark3/types/tmpl"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type StatusSrv struct{}

func (s *StatusSrv) OverallView(req types.StatusOverallViewReq) (resp interface{}, err error) {
	var statusList []model.StatusModel
	if err = global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&statusList).Error; err != nil {
		return nil, err
	}
	type StepVo struct {
		Id            int `json:"id"`
		StartStatusId int `json:"start_status_id"`
		EndStatusId   int `json:"end_status_id"`
	}
	type StatusVo struct {
		Id    int      `json:"id"`
		Name  string   `json:"name"`
		Color string   `json:"color"`
		Mode  string   `json:"mode"`
		Steps []StepVo `json:"steps"`
	}
	var statusVoList []StatusVo
	for _, status := range statusList {
		var stepList []model.StepModel
		if err = global.GVA_DB.Where("start_status_id=?", status.Id).Find(&stepList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		var stepVoList []StepVo
		for _, step := range stepList {
			var stepVo = StepVo{
				Id:            step.Id,
				StartStatusId: step.StartStatusId,
				EndStatusId:   step.EndStatusId,
			}
			stepVoList = append(stepVoList, stepVo)
		}
		var statusVo = StatusVo{
			Id:    status.Id,
			Name:  status.Name,
			Color: status.Color,
			Mode:  status.Mode,
			Steps: stepVoList,
		}
		statusVoList = append(statusVoList, statusVo)
	}

	var coordinate model.StatusCoordinateModel
	global.GVA_DB.Where("tmpl_id=?", req.TmplId).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var statusMap = make(map[int]StatusVo)
		for _, s := range statusVoList {
			statusMap[s.Id] = s
		}

		var sort []StatusVo
		for _, sIdStr := range coordinateSlice {
			sId, err := strconv.Atoi(sIdStr)
			if err != nil {
				continue
			}
			if s, ok := statusMap[sId]; ok {
				sort = append(sort, s)
				delete(statusMap, sId)
			}
		}

		for _, s := range statusMap {
			sort = append(sort, s)
		}
		return sort, nil
	}
	return statusVoList, nil
}

func (s *StatusSrv) GetAll(req types.StatusGetAllReq) (resp interface{}, err error) {
	var statusList []model.StatusModel
	if err = global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&statusList).Error; err != nil {
		return nil, err
	}

	// 排序
	var coordinate model.StatusCoordinateModel
	global.GVA_DB.Where("tmpl_id=?", req.TmplId).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var statusMap = make(map[int]model.StatusModel)
		for _, s := range statusList {
			statusMap[s.Id] = s
		}

		var sort []model.StatusModel
		for _, sIdStr := range coordinateSlice {
			sId, err := strconv.Atoi(sIdStr)
			if err != nil {
				continue
			}
			if s, ok := statusMap[sId]; ok {
				sort = append(sort, s)
				delete(statusMap, sId)
			}
		}

		for _, s := range statusMap {
			sort = append(sort, s)
		}
		return sort, nil
	}

	return statusList, nil
}

func (s *StatusSrv) Create(req types.StatusCreateReq) (resp interface{}, err error) {
	//判断状态名称是否重复
	var status model.StatusModel
	if err = global.GVA_DB.Where("tmpl_id=? and name=?", req.TmplId, req.Name).First(&status).Error; err == nil {
		return nil, errors.New("字段命名重复，请重新命名")
	}
	if err = global.GVA_DB.Create(&model.StatusModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		Name:      req.Name,
		Color:     req.Color,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StatusSrv) Rename(req types.StatusRenameReq) (resp interface{}, err error) {
	var status model.StatusModel
	//判断是否重名
	if err = global.GVA_DB.Where("tmpl_id=? and id !=? and name=?", req.TmplId, req.Id, req.Name).First(&status).Error; err == nil {
		return nil, fmt.Errorf("[%s]状态命名重复，请重新命名", req.Name)
	}
	//判断是否存在
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&status).Error; err != nil {
		return nil, errors.New("参数错误")
	}
	status.Name = req.Name
	status.Color = req.Color
	status.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Save(&status).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StatusSrv) SetFirst(req types.StatusSetFirstReq) (resp interface{}, err error) {
	db := global.GVA_DB.Begin()
	var statusList []model.StatusModel
	if err = db.Where("tmpl_id=? and mode=?", req.TmplId, "first").Find(&statusList).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	if len(statusList) > 0 {
		for k := range statusList {
			statusList[k].Mode = "other"
		}
		if err = db.Save(statusList).Error; err != nil {
			db.Rollback()
			return nil, err
		}
	}

	if err = db.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Updates(&model.StatusModel{
		Mode: "first",
	}).Error; err != nil {
		db.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	db.Commit()
	return
}

func (s *StatusSrv) Move(req types.StatusMoveReq) (resp interface{}, err error) {
	var coordinate model.StatusCoordinateModel
	global.GVA_DB.Where("tmpl_id=?", req.TmplId).First(&coordinate)
	if coordinate.Id == 0 {
		if err = global.GVA_DB.Create(&model.StatusCoordinateModel{
			WsId:       req.WsId,
			TmplId:     req.TmplId,
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else {
		if err = global.GVA_DB.Where("tmpl_id=?", req.TmplId).Updates(&model.StatusCoordinateModel{
			Coordinate: req.Coordinate,
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	}
	return
}

func (s *StatusSrv) PreCheckDelete(req types.StatusPreCheckDeleteReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"status": req.Id},
		},
	}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	//tips : 该状态已关联以下工作项,请修改以下工作项的状态后,再进行删除
	return count, nil
}

// 删除状态方案 (先采用方案1)
// 方案1： 如果该状态下还有任务，则提醒用户修改这些状态下任务的状态
// 方案2： 不提供删除状态功能，只有废弃该状态
// 方案3： 直接可以删除，但任务中的涉及该状态的值会变为空
func (s *StatusSrv) Delete(req types.StatusDeleteReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"status": req.Id},
		},
	}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("该状态已关联以下工作项,请修改以下工作项的状态后,再进行删除")
	}
	var firstStatus model.StatusModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=? and mode=?", req.TmplId, req.Id, "first").Find(&firstStatus).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	if firstStatus.Id != 0 {
		return nil, errors.New("初始状态不能删除")
	}

	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Delete(&model.StatusModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StatusSrv) StatusNext(req types.StatusNextReq) (resp interface{}, err error) {
	type StepVo struct {
		Id            int `json:"id"`
		StartStatusId int `json:"start_status_id"`
		EndStatusId   int `json:"end_status_id"`
	}
	statusMap := make(map[int]map[int]StepVo)
	var statusList []model.StatusModel
	if err = global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&statusList).Error; err != nil {
		return nil, err
	}

	for _, status := range statusList {
		var stepList []model.StepModel
		if err = global.GVA_DB.Where("start_status_id=?", status.Id).Find(&stepList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		for _, step := range stepList {
			var stepVo = StepVo{
				Id:            step.Id,
				StartStatusId: step.StartStatusId,
				EndStatusId:   step.EndStatusId,
			}
			if statusMap[step.StartStatusId] == nil {
				statusMap[step.StartStatusId] = make(map[int]StepVo)
			}
			statusMap[step.StartStatusId][step.EndStatusId] = stepVo
		}
	}

	return statusMap, nil
}
