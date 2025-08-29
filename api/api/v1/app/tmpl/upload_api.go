package tmpl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mark3/global"
	"mark3/pkg/ctl"
	"mark3/pkg/right"
	fileModel "mark3/repository/db/model/app/tmpl"
	"mark3/service/app/tmpl"
	types "mark3/types/app/tmpl"
	"net/url"
	"os"
	"path"
	"strings"
)

type UploadApi struct{}

func (a *UploadApi) GetUploadExtConfig(ctx *gin.Context) {
	l := new(tmpl.UploadSrv)
	resp, err := l.GetFileExtConfig()
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) UploadImageForHtml(ctx *gin.Context) {
	file, err := ctx.FormFile("tmpl_html_image")
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	var req types.UploadImageForHtml
	if err := ctx.ShouldBind(&req); err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	userid, _ := ctx.Get("userid")
	//判断是否为模版成员
	_, err = right.NewUserTmplRight(userid.(int), req.WsId, req.TmplId)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	l := new(tmpl.UploadSrv)
	resp, err := l.UploadImageForHtml(req, file)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) GetFileList(ctx *gin.Context) {
	var req types.GetFileList
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

	l := new(tmpl.UploadSrv)
	resp, err := l.GetFileList(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) GetFileOne(ctx *gin.Context) {
	var req types.DownloadFile
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

	l := new(tmpl.UploadSrv)
	resp, err := l.GetFile(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	file := resp.(fileModel.FileModel)
	filePath := file.RelativePath

	if file.TransformedStatus == fileModel.Transforming {
		ctl.FailWithMessage("文件正在转换中，请稍后", ctx)
		return
	}

	if file.TransformedStatus == fileModel.Failed {
		ctl.FailWithMessage("文件转换失败，请重新上传", ctx)
		return
	}

	if req.DownloadFileType == "transformed_original_name" {
		filePath = file.TransformedRelativePath
		exists, _ := global.GVA_RDB.Get(strings.Split(path.Base(filePath), ".")[0]).Result()

		if exists == "awaiting" {
			ctl.FailWithMessage("文件正在转换中，请稍后", ctx)
			return
		}
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("tmpl_file")
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	var req types.UploadFile
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

	l := new(tmpl.UploadSrv)
	resp, err := l.UploadFile(userid.(int), req, file, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) DownloadFile(ctx *gin.Context) {
	var req types.DownloadFile
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

	l := new(tmpl.UploadSrv)
	resp, err := l.GetFile(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	file := resp.(fileModel.FileModel)
	fileName := file.OriginalName
	filePath := file.RelativePath

	if req.DownloadFileType == "transformed_original_name" {
		fileName = file.TransformedOriginalName
		filePath = file.TransformedRelativePath
	}
	_, err = os.Open(global.GVA_CONFIG.Upload.Path + filePath)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", url.PathEscape(fileName)))
	ctx.Writer.Header().Add("filename", url.PathEscape(fileName))
	ctx.Writer.Header().Add("Access-Control-Expose-Headers", "filename")
	ctx.File(global.GVA_CONFIG.Upload.Path + filePath)
}

func (a *UploadApi) DeleteFile(ctx *gin.Context) {
	var req types.DeleteFile
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

	l := new(tmpl.UploadSrv)
	resp, err := l.DeleteFile(userid.(int), req, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}

func (a *UploadApi) UpdateFileIsCurrentVersion(ctx *gin.Context) {
	var req types.UpdateFileIsCurrentVersion
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

	var needUpdateFileReq types.DownloadFile
	needUpdateFileReq.Id = req.Id
	needUpdateFileReq.WsId = req.WsId
	needUpdateFileReq.TmplId = req.TmplId
	needUpdateFileReq.IssueId = req.IssueId
	l := new(tmpl.UploadSrv)
	_, err = l.GetFile(userid.(int), needUpdateFileReq, userTmplRight)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, err := l.UpdateFileIsCurrentVersion(req)
	if err != nil {
		ctl.FailWithMessage(err.Error(), ctx)
		return
	}
	ctl.OkWithData(resp, ctx)
}
