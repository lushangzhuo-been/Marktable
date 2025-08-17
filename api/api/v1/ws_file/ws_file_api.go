package ws_file

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	wsfilemodel "mark3/repository/db/model/ws_file"
	tmplsrv "mark3/service/tmpl"
	srv "mark3/service/ws_file"
	wsfileandtmplmergedsrv "mark3/service/ws_file_and_tmpl_merged"
	wsfileandtmplmergedlistsrv "mark3/service/ws_file_and_tmpl_merged_list"
	tmpltypes "mark3/types/tmpl"
	types "mark3/types/ws_file"
	"mark3/utils"
	"strconv"
	"strings"
)

type WsFileApi struct{}

func (a *WsFileApi) Create(ctx *gin.Context) {
	var req types.AddWsFileReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.AddWsFileSrv)
	_, err := l.AddWsFile(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *WsFileApi) Info(ctx *gin.Context) {
	var req types.WsFileInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.WsFileSrv)
	resp, err := l.InfoView(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *WsFileApi) UpdateWsFileName(ctx *gin.Context) {
	var req types.UpdateWsFileName
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.WsFileSrv)
	_, err := l.UpdateWsFileName(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *WsFileApi) UpdateWsFileCoordinate(ctx *gin.Context) {
	var req types.UpdateWsFileCoordinate
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.WsFileSrv)
	_, err := l.UpdateWsFileCoordinate(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *WsFileApi) Delete(ctx *gin.Context) {
	var req types.WsFileDeleteReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.WsFileSrv)

	WsFile, err := l.Info(types.WsFileInfoReq{WsId: req.WsId, Id: req.Id})
	if err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	if len(WsFile.Coordinate) != 0 {
		Coordinate := strings.Trim(WsFile.Coordinate, ",")
		TmplIdList := strings.Split(Coordinate, ",")
		for _, v := range TmplIdList {
			TmplId, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			//判断是否为模板管理员
			userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, TmplId)
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
			if err = userTmplRight.CanManage(); err != nil {
				ctl.UnPermission(err.Error(), ctx)
				return
			}
		}
	}

	d := new(srv.DeleteWsFileSrv)
	_, err = d.DeleteWsFile(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *WsFileApi) TmplWsFileMove(ctx *gin.Context) {
	var req types.TmplWsFileMoveReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	wsfileandtmplmergedlistsrv := new(wsfileandtmplmergedlistsrv.WsFileAndTmplMergedListSrv)

	wsfiletmplmergedsrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
	wsFileAndTmplMergedCoordinateSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedCoordinateSrv)

	tmplSrv := new(tmplsrv.TmplSrv)

	wsFileSrv := new(srv.WsFileSrv)

	TmplOrWsFileList, err := wsfileandtmplmergedlistsrv.TmplMoveOutWsFileList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	if len(TmplOrWsFileList) != 2 {
		ctl.FailWithMessage("参数异常", ctx)
		return
	}

	// 外层往里层扔
	if (req.Order == "inner" && TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[0].WsFileId == 0 &&
		TmplOrWsFileList[1].ImplTableName == "ws_file") || (TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[0].WsFileId == 0 &&
		TmplOrWsFileList[1].ImplTableName == "tmpl" && TmplOrWsFileList[1].WsFileId > 0) {
		_, err := tmplSrv.UpdateWsFileId(types.TmplMoveInWsFile{
			TmplId: req.FromId,
			WsFileInfoReq: types.WsFileInfoReq{
				Id:   TmplOrWsFileList[1].WsFileId,
				WsId: req.WsId,
			},
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}
		model, err := wsfiletmplmergedsrv.GetWsFileAndTmplMergedModel(req.WsId, "tmpl", req.FromId)
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		_, err = wsFileAndTmplMergedCoordinateSrv.RemoveWsFileAndTmplMergedCoordinate(model)
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		_, err = wsfiletmplmergedsrv.DeleteWsFileAndTmplMerged(model.Id)
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		if TmplOrWsFileList[1].ImplTableName == "ws_file" && req.Order == "inner" {
			_, err := wsFileSrv.AddTmplId(req.FromId, types.WsFileInfoReq{
				WsId: req.WsId,
				Id:   TmplOrWsFileList[1].WsFileId,
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
		} else {
			wsFileModel, err := wsFileSrv.Info(types.WsFileInfoReq{
				WsId: req.WsId,
				Id:   TmplOrWsFileList[1].WsFileId,
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
			Coordinate := utils.LeftListSort(wsFileModel.Coordinate, req.FromId, req.ToId, req.Order)
			wsFileModel.Coordinate = Coordinate
			_, err = wsFileSrv.UpdateWsFileCoordinate(types.UpdateWsFileCoordinate{
				Coordinate: Coordinate,
				WsFileInfoReq: types.WsFileInfoReq{
					WsId: wsFileModel.WsId,
					Id:   wsFileModel.Id,
				},
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
		}
	} else if (req.Order == "inner" && TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[0].WsFileId > 0 &&
		TmplOrWsFileList[1].ImplTableName == "ws_file" &&
		TmplOrWsFileList[0].WsFileId != TmplOrWsFileList[1].WsFileId) ||
		(TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[1].ImplTableName == "tmpl" &&
			TmplOrWsFileList[0].WsFileId > 0 && TmplOrWsFileList[1].WsFileId > 0 &&
			TmplOrWsFileList[0].WsFileId != TmplOrWsFileList[1].WsFileId) {
		// 里层互相之间扔
		_, err := tmplSrv.UpdateWsFileId(types.TmplMoveInWsFile{
			TmplId: req.FromId,
			WsFileInfoReq: types.WsFileInfoReq{
				Id:   TmplOrWsFileList[1].WsFileId,
				WsId: req.WsId,
			},
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		_, err = wsFileSrv.RemoveTmplId(req.FromId, types.WsFileInfoReq{
			WsId: req.WsId,
			Id:   TmplOrWsFileList[0].WsFileId,
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		if TmplOrWsFileList[1].ImplTableName == "ws_file" && req.Order == "inner" {
			_, err = wsFileSrv.AddTmplId(req.FromId, types.WsFileInfoReq{
				WsId: req.WsId,
				Id:   TmplOrWsFileList[1].WsFileId,
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
		} else {
			wsFileModel, err := wsFileSrv.Info(types.WsFileInfoReq{
				WsId: req.WsId,
				Id:   TmplOrWsFileList[1].WsFileId,
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
			Coordinate := utils.LeftListSort(wsFileModel.Coordinate, req.FromId, req.ToId, req.Order)
			wsFileModel.Coordinate = Coordinate
			_, err = wsFileSrv.UpdateWsFileCoordinate(types.UpdateWsFileCoordinate{
				Coordinate: Coordinate,
				WsFileInfoReq: types.WsFileInfoReq{
					WsId: wsFileModel.WsId,
					Id:   wsFileModel.Id,
				},
			})
			if err != nil {
				ctl.FailWithMessage(err.Error(), ctx)
				return
			}
		}
	} else if (TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[0].WsFileId > 0 &&
		TmplOrWsFileList[1].ImplTableName == "ws_file" && req.Order != "inner") || (TmplOrWsFileList[0].ImplTableName == "tmpl" &&
		TmplOrWsFileList[0].WsFileId > 0 && TmplOrWsFileList[1].ImplTableName == "tmpl" &&
		TmplOrWsFileList[1].WsFileId == 0) {
		// 里层往外扔
		_, err := wsFileSrv.RemoveTmplId(req.FromId, types.WsFileInfoReq{
			WsId: req.WsId,
			Id:   TmplOrWsFileList[0].WsFileId,
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		_, err = tmplSrv.UpdateWsFileId(types.TmplMoveInWsFile{
			TmplId: req.FromId,
			WsFileInfoReq: types.WsFileInfoReq{
				Id:   0,
				WsId: req.WsId,
			},
		})

		model, err := wsfiletmplmergedsrv.AddWsFileAndTmplMerged(wsfilemodel.WsFileAndTmplMergedModel{
			WsId:          req.WsId,
			ImplTableName: "tmpl",
			ImplTableId:   req.FromId,
		})

		TomergedModel, err := wsfiletmplmergedsrv.GetWsFileAndTmplMergedModel(req.WsId, TmplOrWsFileList[1].ImplTableName, TmplOrWsFileList[1].Id)
		if err != nil {
			return
		}

		coordinate, err := wsFileAndTmplMergedCoordinateSrv.GetWsFileAndTmplMergedCoordinate(req.WsId)
		if err != nil {
			return
		}

		Coordinate := utils.LeftListSort(coordinate.Coordinate, model.Id, TomergedModel.Id, req.Order)
		_, err = wsFileAndTmplMergedCoordinateSrv.UpdateWsFileAndTmplMergedCoordinate(types.UpdateWsFileAndTmplMergedCoordinate{
			Coordinate: Coordinate,
			WsId:       req.WsId,
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

	} else if TmplOrWsFileList[0].ImplTableName == "tmpl" && TmplOrWsFileList[1].ImplTableName == "tmpl" &&
		TmplOrWsFileList[0].WsFileId > 0 && TmplOrWsFileList[1].WsFileId > 0 &&
		TmplOrWsFileList[0].WsFileId == TmplOrWsFileList[1].WsFileId {
		// 单个里层排序变更
		wsFileModel, err := wsFileSrv.Info(types.WsFileInfoReq{
			WsId: req.WsId,
			Id:   TmplOrWsFileList[1].WsFileId,
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

		Coordinate := utils.LeftListSort(wsFileModel.Coordinate, req.FromId, req.ToId, req.Order)
		wsFileModel.Coordinate = Coordinate
		_, err = wsFileSrv.UpdateWsFileCoordinate(types.UpdateWsFileCoordinate{
			Coordinate: Coordinate,
			WsFileInfoReq: types.WsFileInfoReq{
				WsId: wsFileModel.WsId,
				Id:   wsFileModel.Id,
			},
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}
	} else {
		// 外层排序变更
		model, err := wsfiletmplmergedsrv.GetWsFileAndTmplMergedModel(req.WsId, TmplOrWsFileList[0].ImplTableName,
			TmplOrWsFileList[0].Id)
		if err != nil {
			return
		}

		TomergedModel, err := wsfiletmplmergedsrv.GetWsFileAndTmplMergedModel(req.WsId, TmplOrWsFileList[1].ImplTableName, TmplOrWsFileList[1].Id)
		if err != nil {
			return
		}

		coordinate, err := wsFileAndTmplMergedCoordinateSrv.GetWsFileAndTmplMergedCoordinate(req.WsId)
		if err != nil {
			return
		}

		Coordinate := utils.LeftListSort(coordinate.Coordinate, model.Id, TomergedModel.Id, req.Order)
		_, err = wsFileAndTmplMergedCoordinateSrv.UpdateWsFileAndTmplMergedCoordinate(types.UpdateWsFileAndTmplMergedCoordinate{
			Coordinate: Coordinate,
			WsId:       req.WsId,
		})
		if err != nil {
			ctl.FailWithMessage(err.Error(), ctx)
			return
		}

	}
	ctl.Ok(ctx)
	return
}

func (a *WsFileApi) List(ctx *gin.Context) {
	var req tmpltypes.TmplListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(wsfileandtmplmergedlistsrv.WsFileAndTmplMergedListSrv)
	resp, err := l.List(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
