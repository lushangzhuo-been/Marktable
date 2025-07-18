package ws_file

import (
	"errors"
	"mark3/global"
	tmplmodel "mark3/repository/db/model/tmpl"
	wsfilemodel "mark3/repository/db/model/ws_file"
	wsfileandtmplmergedsrv "mark3/service/ws_file_and_tmpl_merged"
	types "mark3/types/ws_file"
	wsfiletypes "mark3/types/ws_file"
	"strconv"
	"strings"
)

type DeleteWsFileSrv struct{}

func (s *DeleteWsFileSrv) DeleteWsFile(req types.WsFileDeleteReq) (resp interface{}, err error) {
	l := new(WsFileSrv)
	tx := global.GVA_DB.Begin()

	WsFile, err := l.Info(wsfiletypes.WsFileInfoReq{WsId: req.WsId, Id: req.Id})
	if len(WsFile.Coordinate) != 0 {
		Coordinate := strings.Trim(WsFile.Coordinate, ",")
		TmplIdList := strings.Split(Coordinate, ",")
		var TmplIdListInt []int
		for _, v := range TmplIdList {
			TmplId, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			TmplIdListInt = append(TmplIdListInt, TmplId)
		}
		if err = tx.Where("ws_id=? and id in?", req.WsId, TmplIdListInt).Delete(&tmplmodel.TmplModel{}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return nil, errors.New("操作失败, 删除流程异常")
		}
	}

	if err = tx.Where("id=?", req.Id).Delete(&wsfilemodel.WsFileModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return nil, errors.New("操作失败，删除该文件夹异常")
	}

	// 空间下的排序ID去除此文件夹ID 从合并表里面移除ID
	wsFileAndTmplMergedSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
	WsFileAndTmplMergedModel, err := wsFileAndTmplMergedSrv.GetWsFileAndTmplMergedModel(req.WsId, "ws_file", req.Id)
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
	tx.Commit()
	return nil, nil

}
