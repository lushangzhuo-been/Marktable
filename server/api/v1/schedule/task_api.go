package schedule

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	"mark3/schedule"
	srv "mark3/service/rule"
	"mark3/types/rule"
)

type TaskApi struct{}

func (a *TaskApi) GetSchedule(ctx *gin.Context) {
	ctl.OkWithData(schedule.IsStarted, ctx)
}

func (a *TaskApi) SetSchedule(ctx *gin.Context) {
	start := ctx.Query("start")
	if start == "1" {
		schedule.StartAll()
	} else if start == "0" {
		schedule.StopAll()
	}
	ctl.Ok(ctx)
}

func (a *TaskApi) AddOrUpdateSchedule(ctx *gin.Context) {
	var req rule.DetailReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	//userid, _ := ctx.Get("userid")
	////判断是否为模板管理员
	//userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	//if err != nil {
	//	ctl.FailWithMessage(err.Error(), ctx)
	//	return
	//}
	//if err := userTmplRight.CanOperate(); err != nil {
	//	ctl.UnPermission(err.Error(), ctx)
	//	return
	//}

	l := new(srv.RuleSrv)
	resp, err := l.AddOrUpdateSchedule(req.Id)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *TaskApi) DoRuleCron(ctx *gin.Context) {
	var req rule.DetailReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	//userid, _ := ctx.Get("userid")
	////判断是否为模板管理员
	//userTmplRight, err := right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	//if err != nil {
	//	ctl.FailWithMessage(err.Error(), ctx)
	//	return
	//}
	//if err := userTmplRight.CanOperate(); err != nil {
	//	ctl.UnPermission(err.Error(), ctx)
	//	return
	//}

	l := new(srv.RuleSrv)
	resp, err := l.DoRuleCron(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
