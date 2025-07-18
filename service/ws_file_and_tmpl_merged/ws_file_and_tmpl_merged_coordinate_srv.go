package ws_file_and_tmpl_merged

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/repository/db/model/ws_file"
	types "mark3/types/ws_file"
	"strconv"
	"strings"
)

type WsFileAndTmplMergedCoordinateSrv struct{}

func (s *WsFileAndTmplMergedCoordinateSrv) GetWsFileAndTmplMergedCoordinate(WsId int) (resp ws_file.WsFileAndTmplMergedCoordinateModel, err error) {
	var WsFileAndTmplMergedCoordinate ws_file.WsFileAndTmplMergedCoordinateModel
	if err = global.GVA_DB.Where("ws_id=?", WsId).First(&WsFileAndTmplMergedCoordinate).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedCoordinate, errors.New("操作失败, 获取文件夹或流程的合并排序表数据异常")
	}
	return WsFileAndTmplMergedCoordinate, nil
}

func (s *WsFileAndTmplMergedCoordinateSrv) AddWsFileAndTmplMergedCoordinate(WsFileAndTmplMerged ws_file.WsFileAndTmplMergedModel) (resp ws_file.WsFileAndTmplMergedCoordinateModel, err error) {
	WsFileAndTmplMergedCoordinate, err := s.GetWsFileAndTmplMergedCoordinate(WsFileAndTmplMerged.WsId)
	if err != nil {
		return WsFileAndTmplMergedCoordinate, err
	}
	if strings.Contains(WsFileAndTmplMergedCoordinate.Coordinate, ","+strconv.Itoa(WsFileAndTmplMerged.Id)+",") {
		return WsFileAndTmplMergedCoordinate, nil
	}
	var Coordinate string
	if len(WsFileAndTmplMergedCoordinate.Coordinate) == 0 {
		Coordinate = "," + strconv.Itoa(WsFileAndTmplMerged.Id) + ","
	} else {
		Coordinate = "," + strconv.Itoa(WsFileAndTmplMerged.Id) + WsFileAndTmplMergedCoordinate.Coordinate
	}
	WsFileAndTmplMergedCoordinate.Coordinate = Coordinate
	WsFileAndTmplMergedCoordinate.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&ws_file.WsFileAndTmplMergedCoordinateModel{}).Where("ws_id=?", WsFileAndTmplMerged.WsId).Save(WsFileAndTmplMergedCoordinate).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedCoordinate, errors.New("操作失败")
	}
	return WsFileAndTmplMergedCoordinate, nil
}

func (s *WsFileAndTmplMergedCoordinateSrv) RemoveWsFileAndTmplMergedCoordinate(WsFileAndTmplMerged ws_file.WsFileAndTmplMergedModel) (resp ws_file.WsFileAndTmplMergedCoordinateModel, err error) {
	WsFileAndTmplMergedCoordinate, err := s.GetWsFileAndTmplMergedCoordinate(WsFileAndTmplMerged.WsId)
	if err != nil {
		return WsFileAndTmplMergedCoordinate, err
	}

	if !strings.Contains(WsFileAndTmplMergedCoordinate.Coordinate, ","+strconv.Itoa(WsFileAndTmplMerged.Id)+",") {
		return WsFileAndTmplMergedCoordinate, nil
	}
	Coordinate := strings.Replace(WsFileAndTmplMergedCoordinate.Coordinate, ","+strconv.Itoa(WsFileAndTmplMerged.Id)+",", ",", -1)
	if len(Coordinate) == 1 && Coordinate == "," {
		Coordinate = ""
	}
	WsFileAndTmplMergedCoordinate.Coordinate = Coordinate

	WsFileAndTmplMergedCoordinate.Coordinate = Coordinate
	WsFileAndTmplMergedCoordinate.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&ws_file.WsFileAndTmplMergedCoordinateModel{}).Where("ws_id=?", WsFileAndTmplMerged.WsId).Save(WsFileAndTmplMergedCoordinate).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedCoordinate, errors.New("操作失败")
	}
	return WsFileAndTmplMergedCoordinate, nil
}

func (s *WsFileAndTmplMergedCoordinateSrv) UpdateWsFileAndTmplMergedCoordinate(req types.UpdateWsFileAndTmplMergedCoordinate) (resp interface {
}, err error) {
	WsFileAndTmplMergedCoordinate, err := s.GetWsFileAndTmplMergedCoordinate(req.WsId)
	if err != nil {
		return WsFileAndTmplMergedCoordinate, err
	}
	WsFileAndTmplMergedCoordinate.Coordinate = req.Coordinate
	WsFileAndTmplMergedCoordinate.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Model(&ws_file.WsFileAndTmplMergedCoordinateModel{}).Where("ws_id=?", req.WsId).Save(WsFileAndTmplMergedCoordinate).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedCoordinate, errors.New("操作失败")
	}
	return WsFileAndTmplMergedCoordinate, nil
}
