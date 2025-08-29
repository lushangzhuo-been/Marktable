package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/app/tmpl"
	types "mark3/types/app/tmpl"

	"github.com/gin-gonic/gin"
)

type ProgressApi struct{}

func (a *ProgressApi) AddProgress(ctx *gin.Context) {
	var req types.AddProgress
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	_, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.AddProgress(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ProgressApi) UpdateProgress(ctx *gin.Context) {
	var req types.UpdateProgress
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	_, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.UpdateProgress(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ProgressApi) DeleteProgress(ctx *gin.Context) {
	var req types.DeleteProgress
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	_, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.DeleteProgress(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ProgressApi) GetProgressList(ctx *gin.Context) {
	var req types.GetProgressList
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	_, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetProgressList(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
