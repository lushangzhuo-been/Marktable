package ws

import (
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	srv "mark3/service/app/ws"
	types "mark3/types/app/ws"

	"github.com/gin-gonic/gin"
)

type WsApi struct{}

func (a *WsApi) GetMyWsList(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")
	l := new(srv.WsSrv)
	resp, err := l.GetMyWsList(userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *WsApi) GetWsInfo(ctx *gin.Context) {
	var req types.WsInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为空间成员
	if err := right.JudgeWsMember(userid.(int), req.WsId); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.WsSrv)
	resp, err := l.GetWsInfo(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
