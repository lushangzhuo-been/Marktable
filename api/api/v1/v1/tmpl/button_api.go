package tmpl

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"
)

type ButtonApi struct{}

func (a *ButtonApi) Info(ctx *gin.Context) {
	var req types.ButtonInfoReq
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

	l := new(srv.ButtonSrv)
	resp, err := l.Info(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ButtonApi) Create(ctx *gin.Context) {
	var req types.ButtonCreateReq
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

	l := new(srv.ButtonSrv)
	_, err = l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ButtonApi) Update(ctx *gin.Context) {
	var req types.ButtonUpdateReq
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

	l := new(srv.ButtonSrv)
	_, err = l.Update(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ButtonApi) Delete(ctx *gin.Context) {
	var req types.ButtonDeleteReq
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

	l := new(srv.ButtonSrv)
	_, err = l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
