package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/tmpl"
	types "mark3/types/tmpl"

	"github.com/gin-gonic/gin"
)

type FieldApi struct{}

func (a *FieldApi) Config(ctx *gin.Context) {
	l := new(srv.FieldSrv)
	resp, err := l.Config()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *FieldApi) List(ctx *gin.Context) {
	var req types.FieldListReq
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

	l := new(srv.FieldSrv)
	resp, err := l.List(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *FieldApi) Info(ctx *gin.Context) {
	var req types.FieldInfoReq
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

	l := new(srv.FieldSrv)
	resp, err := l.Info(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *FieldApi) Create(ctx *gin.Context) {
	var req types.FieldCreateReq
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

	l := new(srv.FieldSrv)
	_, err = l.Create(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *FieldApi) Update(ctx *gin.Context) {
	var req types.FieldUpdateReq
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

	l := new(srv.FieldSrv)
	_, err = l.Update(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *FieldApi) GetAllPersonCom(ctx *gin.Context) {
	var req types.FieldGetAllPersonComReq
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

	l := new(srv.FieldSrv)
	resp, err := l.GetAllPersonCom(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *FieldApi) GetReadOnlyRule(ctx *gin.Context) {
	var req types.FieldGetReadOnlyRuleReq
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

	l := new(srv.FieldSrv)
	resp, err := l.GetReadOnlyRule(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *FieldApi) UpdateReadOnlyRule(ctx *gin.Context) {
	var req types.FieldUpdateReadOnlyRuleReq
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

	l := new(srv.FieldSrv)
	_, err = l.UpdateReadOnlyRule(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *FieldApi) Delete(ctx *gin.Context) {
	var req types.FieldDeleteReq
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

	l := new(srv.FieldSrv)
	_, err = l.Delete(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *FieldApi) Enumeration(ctx *gin.Context) {
	var req types.FieldListReq
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

	l := new(srv.FieldSrv)
	resp, err := l.Enumeration(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
