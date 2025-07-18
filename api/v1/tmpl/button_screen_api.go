package tmpl

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"
)

type ButtonScreenApi struct{}

func (a *ButtonScreenApi) Config(ctx *gin.Context) {
	l := new(srv.ButtonScreenSrv)
	resp, err := l.Config()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ButtonScreenApi) Surplus(ctx *gin.Context) {
	var req types.ButtonScreenSurplusReq
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

	l := new(srv.ButtonScreenSrv)
	resp, err := l.Surplus(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ButtonScreenApi) List(ctx *gin.Context) {
	var req types.ButtonScreenListReq
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

	l := new(srv.ButtonScreenSrv)
	resp, err := l.List(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ButtonScreenApi) Create(ctx *gin.Context) {
	var req types.ButtonScreenCreateReq
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

	l := new(srv.ButtonScreenSrv)
	_, err = l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ButtonScreenApi) SetRequired(ctx *gin.Context) {
	var req types.ButtonScreenSetRequiredReq
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

	l := new(srv.ButtonScreenSrv)
	_, err = l.SetRequired(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ButtonScreenApi) Move(ctx *gin.Context) {
	var req types.ButtonScreenMoveReq
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

	l := new(srv.ButtonScreenSrv)
	_, err = l.Move(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ButtonScreenApi) Delete(ctx *gin.Context) {
	var req types.ButtonScreenDeleteReq
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

	l := new(srv.ButtonScreenSrv)
	_, err = l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
