package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"

	"github.com/gin-gonic/gin"
)

type StepScreenApi struct{}

func (a *StepScreenApi) Config(ctx *gin.Context) {
	l := new(srv.StepScreenSrv)
	resp, err := l.Config()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *StepScreenApi) Surplus(ctx *gin.Context) {
	var req types.StepScreenSurplusReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	resp, err := l.Surplus(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *StepScreenApi) List(ctx *gin.Context) {
	var req types.StepScreenListReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	resp, err := l.List(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *StepScreenApi) Create(ctx *gin.Context) {
	var req types.StepScreenCreateReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	_, err = l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *StepScreenApi) SetRequired(ctx *gin.Context) {
	var req types.StepScreenSetRequiredReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	_, err = l.SetRequired(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *StepScreenApi) Move(ctx *gin.Context) {
	var req types.StepScreenMoveReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	_, err = l.Move(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *StepScreenApi) Delete(ctx *gin.Context) {
	var req types.StepScreenDeleteReq
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
	if err := userTmplRight.CanManage(); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.StepScreenSrv)
	_, err = l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
