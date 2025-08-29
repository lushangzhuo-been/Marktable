package ws

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/ws"
	types "mark3/types/ws"

	"github.com/gin-gonic/gin"
)

type WsApi struct{}

func (a *WsApi) Create(ctx *gin.Context) {
	var req types.WsCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.WsSrv)
	_, err := l.Create(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *WsApi) Info(ctx *gin.Context) {
	var req types.WsInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.Id); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.WsSrv)
	resp, err := l.Info(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (*WsApi) Update(ctx *gin.Context) {
	var req types.WsUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.Id); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.WsSrv)
	_, err := l.Update(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (*WsApi) Delete(ctx *gin.Context) {
	var req types.WsDeleteReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间管理员
	if err := right.JudgeWsAdmin(userid.(int), req.Id); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.WsSrv)
	_, err := l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
