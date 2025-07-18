package tmpl

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"
)

type TabApi struct{}

func (a *TabApi) TabConfig(ctx *gin.Context) {
	var req types.TmplTabConfigReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TabSrv)
	resp, err := l.TabConfig(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TabApi) TmplTabList(ctx *gin.Context) {
	var req types.TmplTabListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")
	//判断是否为模板管理员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := userTmplRight.CanOperate(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.TabSrv)
	resp, err := l.TmplTabList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TabApi) TmplTabCreate(ctx *gin.Context) {
	var req types.TmplTabCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")
	//判断是否为模板管理员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := userTmplRight.CanOperate(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.TabSrv)
	_, err = l.TmplTabCreate(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TabApi) TmplTabDelete(ctx *gin.Context) {
	var req types.TmplTabDeleteReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")
	//判断是否为模板管理员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := userTmplRight.CanOperate(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.TabSrv)
	_, err = l.TmplTabDelete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
