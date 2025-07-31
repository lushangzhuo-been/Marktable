package ws_file

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	model "mark3/repository/db/model/ws_file"
	wsfileandtmplmergedsrv "mark3/service/ws_file_and_tmpl_merged"
	types "mark3/types/ws_file"
)

type AddWsFileSrv struct{}

func (s *AddWsFileSrv) AddWsFile(userid int, req types.AddWsFileReq) (resp interface {
}, err error) {
	l := new(WsFileSrv)
	wsFile, err := l.InfoByName(req.WsId, req.Name)
	if err == nil {

		return nil, errors.New("操作失败，该空间下此文件夹已存在")
	}

	tx := global.GVA_DB.Begin()

	wsFile = model.WsFileModel{
		WsId:          req.WsId,
		Name:          req.Name,
		Desc:          req.Desc,
		CreatedUserId: userid,
		UpdatedUserId: userid,
		Coordinate:    "",
		CreatedAt:     common.GetCurrentTime(),
		UpdatedAt:     common.GetCurrentTime(),
	}
	if err = tx.Create(&wsFile).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 流程和文件合并表里面添加ID
	wsFileAndTmplMergedSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
	wsFileAndTmplMerged, err := wsFileAndTmplMergedSrv.AddWsFileAndTmplMerged(model.WsFileAndTmplMergedModel{
		WsId:          req.WsId,
		ImplTableName: "ws_file",
		ImplTableId:   wsFile.Id,
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
	tx.Commit()
	return wsFile, nil
}
