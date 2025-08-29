package ws

import (
	"github.com/gin-gonic/gin"
	"mark3/pkg/ctl"
	srv "mark3/service/ws"
	"mark3/types/ws"
)

type CommonApi struct{}

func (a *CommonApi) GetUserList(ctx *gin.Context) {
	var req ws.GetUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	l := new(srv.CommonSrv)
	resp, err := l.GetUserList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
