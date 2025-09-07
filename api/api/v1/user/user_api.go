package user

import (
	"mark3/pkg/ctl"
	srv "mark3/service/user"
	"mark3/types/user"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (a *UserApi) Register(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	//todo 判断username, password 位数

	l := new(srv.UserSrv)
	_, err := l.Register(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) Login(ctx *gin.Context) {
	var req user.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	l := new(srv.UserSrv)
	resp, err := l.Login(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UserApi) GetNewToken(ctx *gin.Context) {
	l := new(srv.UserSrv)
	resp, err := l.GetNewToken(ctx)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UserApi) Info(ctx *gin.Context) {
	userid, _ := ctx.Get("userid")
	l := new(srv.UserSrv)
	resp, err := l.Info(userid.(int))
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UserApi) Update(ctx *gin.Context) {
	var req user.UpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.UserSrv)
	_, err := l.Update(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.UserSrv)
	resp, err := l.UploadAvatar(userid.(int), file)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UserApi) UpdatePwd(ctx *gin.Context) {
	var req user.UpdatePwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.UserSrv)
	_, err := l.UpdatePwd(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) RegisterEmailCode(ctx *gin.Context) {
	// 用户注册，发送邮箱验证码
	var req user.EmailParamsReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	l := new(srv.UserSrv)
	_, err := l.RegisterEmailCode(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) RegisterV2(ctx *gin.Context) {
	var req user.RegisterV2Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.UserSrv)
	_, err := l.RegisterV2(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) ResetPwdEmailCode(ctx *gin.Context) {
	// 重置密码，发送邮箱验证码
	var req user.ResetPwdEmailCodeReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	l := new(srv.UserSrv)
	_, err := l.ResetPwdEmailCode(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) ResetPwd(ctx *gin.Context) {
	var req user.ResetPwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(srv.UserSrv)
	_, err := l.ResetPwd(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}

func (a *UserApi) UpdateEmail(ctx *gin.Context) {
	var req user.UpdateEmailReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	l := new(srv.UserSrv)
	_, err := l.UpdateEmail(userid.(int), req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.Ok(ctx)
}
