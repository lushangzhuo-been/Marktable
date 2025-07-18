package message

import (
	"mark3/pkg/ctl"
	srv "mark3/service/message"
	types "mark3/types/message"

	"github.com/gin-gonic/gin"
)

type MessageApi struct{}

func (a *MessageApi) GetUnReadCount(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")

	l := new(srv.MessageSrv)
	resp, err := l.GetUnReadCount(userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *MessageApi) GetMessageList(ctx *gin.Context) {
	var req types.GetMessageListReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")

	l := new(srv.MessageSrv)
	resp, err := l.GetMessageList(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *MessageApi) ReadAllMessage(ctx *gin.Context) {
	var req types.ReadAllMessageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")

	l := new(srv.MessageSrv)
	resp, err := l.ReadAllMessage(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)

}

func (a *MessageApi) ReadMessage(ctx *gin.Context) {
	var req types.ReadMessageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")

	l := new(srv.MessageSrv)
	resp, err := l.ReadMessage(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
