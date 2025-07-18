package ws_file_and_tmpl_merged_list

import (
	"errors"
	"mark3/global"
	wsEnum "mark3/repository/db/enum/ws"
	wsModel "mark3/repository/db/model/ws"
	"mark3/repository/db/model/ws_file"
	tmplsrv "mark3/service/tmpl"
	wsfilesrv "mark3/service/ws_file"
	"mark3/service/ws_file_and_tmpl_merged"
	types "mark3/types/tmpl"
	wsfiletypes "mark3/types/ws_file"
	"strconv"
	"strings"
)

type WsFileAndTmplMergedListSrv struct{}

func (s *WsFileAndTmplMergedListSrv) getWsFileIsAdmin(userid int, wsId int) (IsAdmin string, err error) {
	var member wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=? and status=? and role=?", wsId, userid, wsEnum.WsStatusOk, wsEnum.WsRoleAdmin).First(&member).Error; err == nil {
		return "yes", nil
	}
	return "no", nil
}

func (s *WsFileAndTmplMergedListSrv) List(userId int, req types.TmplListReq) (resp interface{}, err error) {
	WsFileAndTmplMergedSrv := new(ws_file_and_tmpl_merged.WsFileAndTmplMergedSrv)
	wsFileAndTmplMergedList, err := WsFileAndTmplMergedSrv.GetWsFileAndTmplMergedModelList(req.WsId)
	if err != nil {
		return nil, err
	}
	if len(wsFileAndTmplMergedList) == 0 {
		return nil, nil
	}

	wsFileAndTmplMergedMap := make(map[int]ws_file.WsFileAndTmplMergedModel)
	for _, v := range wsFileAndTmplMergedList {
		wsFileAndTmplMergedMap[v.Id] = v
	}
	var WsFileList []ws_file.WsFileModel
	var TmplList []wsfiletypes.TmplVo
	WsFileMap := make(map[int]ws_file.WsFileModel)
	TmplMap := make(map[int]wsfiletypes.TmplVo)

	WsFileSrv := new(wsfilesrv.WsFileSrv)
	WsFileList, err = WsFileSrv.List(req.WsId)
	if len(WsFileList) != 0 {
		for _, v := range WsFileList {
			WsFileMap[v.Id] = v
		}
	}

	TmplSrv := new(tmplsrv.TmplSrv)
	TmplList, err = TmplSrv.ListNotInWsFile(types.WsFileTmplListReq{
		UserId:   userId,
		WsId:     req.WsId,
		WsFileId: 0,
	})
	if len(TmplList) != 0 {
		for _, v := range TmplList {
			TmplMap[v.Id] = v
		}
	}

	wsFileAndTmplMergedCoordinateSrv := new(ws_file_and_tmpl_merged.WsFileAndTmplMergedCoordinateSrv)
	WsFileAndTmplMergedCoordinate, err := wsFileAndTmplMergedCoordinateSrv.GetWsFileAndTmplMergedCoordinate(req.WsId)
	if err != nil {
		return nil, err
	}

	if len(WsFileAndTmplMergedCoordinate.Coordinate) == 0 {
		return nil, errors.New("查询失败, 排序数据异常")
	}

	var res []wsfiletypes.WsFileAndTmplMergedVo
	IsAdmin, err := s.getWsFileIsAdmin(userId, req.WsId)
	if err != nil {
		return nil, err
	}
	Coordinates := strings.Split(WsFileAndTmplMergedCoordinate.Coordinate, ",")
	for _, v := range Coordinates {
		if len(v) == 0 {
			continue
		}
		WsFileAndTmplMergedId, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		wsFileAndTmplMerged, ok := wsFileAndTmplMergedMap[WsFileAndTmplMergedId]
		if !ok {
			continue
		}
		if wsFileAndTmplMerged.ImplTableName == "ws_file" {
			WsFile, wsFileOk := WsFileMap[wsFileAndTmplMerged.ImplTableId]
			if !wsFileOk {
				continue
			}
			WsFileTmplList, err := TmplSrv.WsFileTmplList(types.WsFileTmplListReq{
				UserId:   userId,
				WsId:     WsFile.WsId,
				WsFileId: WsFile.Id,
			})
			if err != nil {
				return nil, err
			}
			res = append(res, wsfiletypes.WsFileAndTmplMergedVo{
				MergedId:      wsFileAndTmplMerged.Id,
				Id:            WsFile.Id,
				WsId:          WsFile.WsId,
				WsFileId:      WsFile.Id,
				IsMember:      "yes",
				IsAdmin:       IsAdmin,
				Name:          WsFile.Name,
				ImplTableName: "ws_file",
				Children:      WsFileTmplList,
				Coordinate:    WsFile.Coordinate,
				CreatedAt:     WsFile.CreatedAt,
				UpdatedAt:     WsFile.UpdatedAt,
			})
		}

		if wsFileAndTmplMerged.ImplTableName == "tmpl" {
			Tmpl, TmplOk := TmplMap[wsFileAndTmplMerged.ImplTableId]
			if !TmplOk {
				continue
			}
			res = append(res, wsfiletypes.WsFileAndTmplMergedVo{
				MergedId:      wsFileAndTmplMerged.Id,
				Id:            Tmpl.Id,
				WsId:          Tmpl.WsId,
				WsFileId:      Tmpl.WsFileId,
				IsMember:      Tmpl.IsMember,
				IsAdmin:       Tmpl.IsAdmin,
				ImplTableName: "tmpl",
				Name:          Tmpl.Name,
				Desc:          Tmpl.Desc,
				Mode:          Tmpl.Mode,
				CreatedAt:     Tmpl.CreatedAt,
				UpdatedAt:     Tmpl.UpdatedAt,
			})
		}

	}
	return res, nil
}

func (s *WsFileAndTmplMergedListSrv) TmplMoveOutWsFileList(req wsfiletypes.TmplWsFileMoveReq) (resp []wsfiletypes.WsFileAndTmplMergedVo, err error) {

	var res []wsfiletypes.WsFileAndTmplMergedVo
	type FromAndToReq struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
	}
	var FromAndToList []FromAndToReq
	FromAndToList = append(FromAndToList, FromAndToReq{
		Id:   req.FromId,
		Type: req.FromType,
	})
	FromAndToList = append(FromAndToList, FromAndToReq{
		Id:   req.ToId,
		Type: req.ToType,
	})
	w := new(wsfilesrv.WsFileSrv)
	t := new(tmplsrv.TmplSrv)
	for _, v := range FromAndToList {
		if v.Type == "ws_file" {
			WsFile, err := w.Info(wsfiletypes.WsFileInfoReq{
				WsId: req.WsId,
				Id:   v.Id,
			})
			if err != nil {
				return nil, err
			}
			res = append(res, wsfiletypes.WsFileAndTmplMergedVo{
				Id:            WsFile.Id,
				WsId:          WsFile.WsId,
				WsFileId:      WsFile.Id,
				Name:          WsFile.Name,
				ImplTableName: "ws_file",
				Coordinate:    WsFile.Coordinate,
				CreatedAt:     WsFile.CreatedAt,
				UpdatedAt:     WsFile.UpdatedAt,
			})
		}

		if v.Type == "tmpl" {
			Tmpl, err := t.Info(types.TmplInfoReq{
				WsId: req.WsId,
				Id:   v.Id,
			})

			if err != nil {
				return nil, err
			}

			res = append(res, wsfiletypes.WsFileAndTmplMergedVo{
				Id:            Tmpl.Id,
				WsId:          Tmpl.WsId,
				WsFileId:      Tmpl.WsFileId,
				ImplTableName: "tmpl",
				Name:          Tmpl.Name,
				Desc:          Tmpl.Desc,
				Mode:          Tmpl.Mode,
				CreatedAt:     Tmpl.CreatedAt,
				UpdatedAt:     Tmpl.UpdatedAt,
			})
		}

	}
	return res, nil
}
