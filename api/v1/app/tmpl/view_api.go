package tmpl

import (
	"mark3/pkg/ctl"
	"mark3/service/app/tmpl"
	types "mark3/types/app/tmpl"

	"github.com/gin-gonic/gin"
)

type ViewApi struct {
}

func (a *ViewApi) GetUserViews(ctx *gin.Context) {
	var req types.GetUserViewsReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")

	l := new(tmpl.ViewSrv)
	resp, err := l.GetUserViews(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ViewApi) GetViewInfo(ctx *gin.Context) {
	var req types.GetViewInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")

	l := new(tmpl.ViewSrv)
	resp, err := l.GetViewInfo(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *ViewApi) CreateView(ctx *gin.Context) {
	var req types.CreateViewReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")

	l := new(tmpl.ViewSrv)
	_, err := l.CreateView(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) RenameView(ctx *gin.Context) {
	var req types.RenameViewReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	userid, _ := ctx.Get("userid")

	l := new(tmpl.ViewSrv)
	_, err := l.RenameView(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) UpdateViewFilter(ctx *gin.Context) {
	var req types.UpdateViewFilterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.UpdateViewFilter(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) UpdateViewSortOrder(ctx *gin.Context) {
	var req types.UpdateViewSortOrderReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.UpdateViewSortOrder(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) UpdateViewColumns(ctx *gin.Context) {
	var req types.UpdateViewColumnsReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.UpdateViewColumns(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) UpdateViewCardGroup(ctx *gin.Context) {
	var req types.UpdateViewCardGroupReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.UpdateViewCardGroup(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) DeleteView(ctx *gin.Context) {
	var req types.DeleteViewReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.DeleteView(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) PinView(ctx *gin.Context) {
	var req types.PinViewReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.PinView(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) UnPinView(ctx *gin.Context) {
	var req types.UnPinViewReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.UnPinView(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *ViewApi) SetViewCoordinate(ctx *gin.Context) {
	var req types.SetViewCoordinateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(tmpl.ViewSrv)
	_, err := l.SetViewCoordinate(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
