package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	"mark3/service/app"
	srv "mark3/service/app/tmpl"
	types "mark3/types/app/tmpl"

	"github.com/gin-gonic/gin"
)

type TmplApi struct{}

func (a *TmplApi) GetTmplList(ctx *gin.Context) {
	var req types.GetTmplListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间成员
	if err := right.JudgeWsMember(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetTmplList(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetTmplOtherList(ctx *gin.Context) {
	var req types.GetTmplOtherListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间成员
	if err := right.JudgeWsMember(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetTmplOtherList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetTmplInfo(ctx *gin.Context) {
	var req types.GetTmplInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")
	//判断是否为空间成员
	if err := right.JudgeWsMember(userid.(int), req.WsId); err != nil {
		ctl.UnPermission(err.Error(), ctx)
		return
	}
	l := new(srv.TmplSrv)
	resp, err := l.GetTmplInfo(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	go log(userid.(int), req.WsId, req.TmplId)

	ctl.OkWithData(resp, ctx)
}

// 记录最近访问的模板
func log(userid int, wsId int, tmplId int) {
	l := new(app.CommonSrv)
	l.LogUserUserRecentlyVisitedTmpl(userid, wsId, tmplId)
}

func (a *TmplApi) GetConfig(ctx *gin.Context) {
	var req types.GetConfigReq
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
	resp, err := l.GetConfig(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetStatusList(ctx *gin.Context) {
	var req types.GetStatusListReq
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
	resp, err := l.GetStatusList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetScreen(ctx *gin.Context) {
	var req types.GetScreenReq
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
	resp, err := l.GetScreen(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetListData(ctx *gin.Context) {
	var req types.GetListDataReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetListData(ctx, userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetListDataSelect(ctx *gin.Context) {
	var req types.GetListDataSelectReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.TmplSrv)
	resp, err := l.GetListDataSelect(req, userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetUserAuth(ctx *gin.Context) {
	var req types.GetUserAuthReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetUserAuth(ctx, userid.(int), req, userTmplRight)
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetData(ctx *gin.Context) {
	var req types.GetDataReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetData(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetFileRight(ctx *gin.Context) {
	var req types.GetDataReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetFileRight(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) CreateAction(ctx *gin.Context) {
	var req types.CommonReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.CreateAction(ctx, req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) CreateSubAction(ctx *gin.Context) {
	var req types.CreateSubActionReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.CreateSubAction(ctx, req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) UpdateAction(ctx *gin.Context) {
	var req types.UpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.Update(ctx, req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) DeleteAction(ctx *gin.Context) {
	var req types.DeleteReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.DeleteAction(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) DeleteSubAction(ctx *gin.Context) {
	var req types.DeleteSubActionReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.DeleteSubAction(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) GetStepList(ctx *gin.Context) {
	var req types.GetStepListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetStepList(req, userid.(int), userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetStepScreen(ctx *gin.Context) {
	var req types.GetStepScreenReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	resp, err := l.GetStepScreen(ctx, req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) SwitchStepAction(ctx *gin.Context) {
	var req types.SwitchStepReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.TmplSrv)
	_, err = l.SwitchStep(ctx, req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *TmplApi) AddProgress(ctx *gin.Context) {
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

func (a *TmplApi) UpdateProgress(ctx *gin.Context) {
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

func (a *TmplApi) DeleteProgress(ctx *gin.Context) {
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

func (a *TmplApi) GetProgressList(ctx *gin.Context) {
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

func (a *TmplApi) GetLogList(ctx *gin.Context) {
	var req types.GetLogList
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
	resp, err := l.GetLogList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TmplApi) GetSubListCount(ctx *gin.Context) {
	var req types.GetSubListCountReq
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
	resp, err := l.GetSubListCount(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
