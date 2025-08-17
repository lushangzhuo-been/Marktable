package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"

	"github.com/gin-gonic/gin"
)

type NodeApi struct{}

func (a *NodeApi) Config(ctx *gin.Context) {
	l := new(srv.NodeSrv)
	resp, err := l.Config()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *NodeApi) List(ctx *gin.Context) {
	var req types.NodeListReq
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

	l := new(srv.NodeSrv)
	resp, err := l.List(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *NodeApi) Info(ctx *gin.Context) {
	var req types.NodeInfoReq
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

	l := new(srv.NodeSrv)
	resp, err := l.Info(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *NodeApi) Create(ctx *gin.Context) {
	var req types.NodeCreateReq
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

	l := new(srv.NodeSrv)
	_, err = l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *NodeApi) Update(ctx *gin.Context) {
	var req types.NodeUpdateReq
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

	l := new(srv.NodeSrv)
	_, err = l.Update(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *NodeApi) GetAllStatus(ctx *gin.Context) {
	var req types.NodeGetAllStatusesReq
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

	l := new(srv.NodeSrv)
	resp, err := l.GetAllStatus(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *NodeApi) Delete(ctx *gin.Context) {
	var req types.NodeDeleteReq
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

	l := new(srv.NodeSrv)
	_, err = l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
