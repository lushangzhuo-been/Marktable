package tmpl

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"
)

type TmplSubConfigApi struct{}

func (a *TmplSubConfigApi) TmplSubConfigCheck(ctx *gin.Context) {
	var req types.TmplSubConfigCheckReq
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

	l := new(srv.TmplSubConfigSrv)
	resp, err := l.TmplSubConfigCheck(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplSubConfigApi) TmplSubConfigList(ctx *gin.Context) {
	var req types.TmplSubConfigListReq
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

	l := new(srv.TmplSubConfigSrv)
	resp, err := l.TmplSubConfigList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplSubConfigApi) TmplSubConfigCreate(ctx *gin.Context) {
	var req types.TmplSubConfigCreateReq
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

	l := new(srv.TmplSubConfigSrv)
	_, err = l.TmplSubConfigCreate(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplSubConfigApi) TmplSubConfigDelete(ctx *gin.Context) {
	var req types.TmplSubConfigDeleteReq
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

	l := new(srv.TmplSubConfigSrv)
	_, err = l.TmplSubConfigDelete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
