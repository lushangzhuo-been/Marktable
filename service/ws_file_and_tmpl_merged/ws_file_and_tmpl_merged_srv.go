package ws_file_and_tmpl_merged

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/repository/db/model/ws_file"
)

type WsFileAndTmplMergedSrv struct{}

func (s *WsFileAndTmplMergedSrv) GetWsFileAndTmplMergedModel(WsId int, ImplTableName string, ImplTableId int) (resp ws_file.WsFileAndTmplMergedModel, err error) {
	var WsFileAndTmplMergedModel ws_file.WsFileAndTmplMergedModel
	if err = global.GVA_DB.Where("ws_id=? and impl_table_name=? and impl_table_id=?", WsId, ImplTableName, ImplTableId).First(&WsFileAndTmplMergedModel).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedModel, errors.New("操作失败, 获取文件夹或流程的合并表数据异常")
	}
	return WsFileAndTmplMergedModel, nil
}

func (s *WsFileAndTmplMergedSrv) GetWsFileAndTmplMergedModelList(WsId int) (resp []ws_file.WsFileAndTmplMergedModel, err error) {
	var WsFileAndTmplMergedModelList []ws_file.WsFileAndTmplMergedModel
	if err = global.GVA_DB.Where("ws_id=?", WsId).Find(&WsFileAndTmplMergedModelList).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return WsFileAndTmplMergedModelList, errors.New("操作失败, 获取文件夹或流程的合并表数据异常")
	}
	return WsFileAndTmplMergedModelList, nil
}

func (s *WsFileAndTmplMergedSrv) AddWsFileAndTmplMerged(wsFileAndTmplMerged ws_file.WsFileAndTmplMergedModel) (resp ws_file.WsFileAndTmplMergedModel, err error) {
	wsFileAndTmplMerged.CreatedAt = common.GetCurrentTime()
	wsFileAndTmplMerged.UpdatedAt = common.GetCurrentTime()
	tx := global.GVA_DB.Begin()
	if err = tx.Create(&wsFileAndTmplMerged).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return wsFileAndTmplMerged, errors.New("操作失败, 创建文件夹或流程的合并表数据异常")
	}
	tx.Commit()
	return wsFileAndTmplMerged, nil
}

func (s *WsFileAndTmplMergedSrv) DeleteWsFileAndTmplMerged(Id int) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", Id).Delete(&ws_file.WsFileAndTmplMergedModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败, 删除文件夹或流程的合并表数据异常")
	}
	return true, nil
}
