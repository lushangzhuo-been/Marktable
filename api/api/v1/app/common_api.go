package app

import (
	"mark3/pkg/ctl"
	srv "mark3/service/app"
	types "mark3/types/app"

	"github.com/gin-gonic/gin"
)

type CommonApi struct{}

func (a *CommonApi) GetUserRecentlyVisitedLog(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")
	l := new(srv.CommonSrv)
	resp, err := l.GetUserUserRecentlyVisitedLog(userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *CommonApi) GetWsUserList(ctx *gin.Context) {
	var req types.CommonGetWsUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.CommonSrv)
	resp, err := l.GetWsUserList(req, userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
