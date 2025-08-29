package ws

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/ws"
	types "mark3/types/ws"
)

type MemberApi struct{}

func (a *MemberApi) Config(ctx *gin.Context) {
	l := new(srv.MemberSrv)
	resp, err := l.Config()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *MemberApi) List(ctx *gin.Context) {
	var req types.MemberListReq
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

	l := new(srv.MemberSrv)
	resp, err := l.List(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *MemberApi) Create(ctx *gin.Context) {
	var req types.MemberCreateReq
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

	l := new(srv.MemberSrv)
	_, err := l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *MemberApi) Update(ctx *gin.Context) {
	var req types.MemberUpdateReq
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

	l := new(srv.MemberSrv)
	_, err := l.Update(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *MemberApi) Delete(ctx *gin.Context) {
	var req types.MemberDeleteReq
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

	l := new(srv.MemberSrv)
	_, err := l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
