package tmpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/pkg/right"
	fileModel "mark3/repository/db/model/app/tmpl"
	types "mark3/types/app/tmpl"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadSrv struct{}

func (s *UploadSrv) GetFileExtConfig() (resp interface{}, err error) {
	allExt := global.GVA_CONFIG.Upload.AllExt
	return allExt, nil
}

func (s *UploadSrv) UploadImageForHtml(req types.UploadImageForHtml, file *multipart.FileHeader) (resp interface{}, err error) {
	//判断文件大小
	if file.Size > global.GVA_CONFIG.Upload.Size*1024 {
		return nil, fmt.Errorf("上传文件过大（最大 %d M）", global.GVA_CONFIG.Upload.Size)
	}

	fileExt := path.Ext(file.Filename)

	var extConfig []string
	for _, ext := range strings.Split(global.GVA_CONFIG.Upload.ImageExt, ",") {
		ext := "." + ext
		extConfig = append(extConfig, ext, strings.ToUpper(ext))
	}

	//判断文件类型
	if !common.InArray(fileExt, extConfig) {
		return nil, fmt.Errorf("只能上传（文件格式：%s）", global.GVA_CONFIG.Upload.ImageExt)
	}
	src, err := file.Open()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	defer src.Close()

	dirName := "/image/tmpl/" + strconv.Itoa(req.TmplId)
	fileName := uuid.New().String() + fileExt

	if _, err = os.Stat(global.GVA_CONFIG.Upload.Path + "/" + dirName); os.IsNotExist(err) {
		if err = os.MkdirAll(global.GVA_CONFIG.Upload.Path+"/"+dirName, os.ModePerm); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("创建目录失败")
		}
	}

	out, err := os.Create(global.GVA_CONFIG.Upload.Path + dirName + "/" + fileName)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	return dirName + "/" + fileName, nil
}

func (s *UploadSrv) UploadFile(userid int, req types.UploadFile, file *multipart.FileHeader, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.IssueId)
	if err != nil {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"_id": objectID},
		},
	}

	var oldDocument bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	if getDataEditRight(userid, oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	//判断文件大小
	if file.Size > global.GVA_CONFIG.Upload.Size*1024 {
		return nil, fmt.Errorf("上传文件过大（最大 %d M）", global.GVA_CONFIG.Upload.Size)
	}

	var extConfig []string
	for _, ext := range strings.Split(global.GVA_CONFIG.Upload.AllExt, ",") {
		ext = "." + ext
		extConfig = append(extConfig, ext, strings.ToUpper(ext))
	}

	fileExt := path.Ext(file.Filename)
	if !common.InArray(fileExt, extConfig) {
		return nil, fmt.Errorf("只能上传（文件格式：%s）", global.GVA_CONFIG.Upload.AllExt)
	}
	src, err := file.Open()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	defer src.Close()

	dirName := "/file/tmpl/" + strconv.Itoa(req.TmplId)
	fileName := uuid.New().String() + fileExt

	if _, err = os.Stat(global.GVA_CONFIG.Upload.Path + "/" + dirName); os.IsNotExist(err) {
		if err = os.MkdirAll(global.GVA_CONFIG.Upload.Path+"/"+dirName, os.ModePerm); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("创建目录失败")
		}
	}

	out, err := os.Create(global.GVA_CONFIG.Upload.Path + dirName + "/" + fileName)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}

	TransformedOriginalName := file.Filename
	TransformedRelativePath := dirName + "/" + fileName
	var needTransformed2PDFExtConfig []string
	for _, ext := range strings.Split(global.GVA_CONFIG.Upload.NeedTransformed2PDFExt, ",") {
		ext = "." + ext
		needTransformed2PDFExtConfig = append(needTransformed2PDFExtConfig, ext, strings.ToUpper(ext))
	}

	var needTransformed2EndXExtConfig []string
	for _, ext := range strings.Split(global.GVA_CONFIG.Upload.NeedTransformed2EndX, ",") {
		ext = "." + ext
		needTransformed2EndXExtConfig = append(needTransformed2EndXExtConfig, ext, strings.ToUpper(ext))
	}

	versionCode := 1
	if req.GroupId <= 0 {
		maxGroupId := []int{}
		global.GVA_DB.Model(fileModel.FileModel{}).Select("IFNULL(MAX(group_id), 0) as maxGroupId").Where("issue_id=? and ws_id=? and tmpl_id=? ", req.IssueId, req.WsId, req.TmplId).Pluck("maxGroupId", &maxGroupId)
		req.GroupId = maxGroupId[0] + 1
	}
	if req.GroupId > 0 {
		maxVersionCode := []int{}
		global.GVA_DB.Model(fileModel.FileModel{}).Select("IFNULL(MAX(version_code), 0) as maxVersionCode").Where("issue_id=? and ws_id=? and tmpl_id=? and group_id=?", req.IssueId, req.WsId, req.TmplId, req.GroupId).Pluck("maxVersionCode", &maxVersionCode)
		versionCode = maxVersionCode[0] + 1
		if err = global.GVA_DB.Model(fileModel.FileModel{}).
			Where("issue_id=? and group_id=?", req.IssueId, req.GroupId).
			Update(
				"IsCurrentVersion", 0).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return false, errors.New("操作失败")
		}
	}

	if common.InArray(fileExt, needTransformed2PDFExtConfig) {
		TransformedRelativePath = dirName + "/" + strings.Split(path.Base(fileName), ".")[0] + ".pdf"
		TransformedOriginalName = strings.Split(path.Base(file.Filename), ".")[0] + ".pdf"
	}

	if common.InArray(fileExt, needTransformed2EndXExtConfig) {
		TransformedRelativePath = dirName + "/" + strings.Split(path.Base(fileName), ".")[0] + "." + fileExt + "x"
		TransformedOriginalName = strings.Split(path.Base(file.Filename), ".")[0] + "." + fileExt + "x"
	}

	transformedStatus := fileModel.NoTransform
	userTmplFile := fileModel.FileModel{
		WsId:                    req.WsId,
		TmplId:                  req.TmplId,
		IssueId:                 req.IssueId,
		GroupId:                 req.GroupId,
		IsCurrentVersion:        req.IsCurrentVersion,
		VersionCode:             versionCode,
		OriginalName:            file.Filename,
		RelativePath:            dirName + "/" + fileName,
		FileSize:                int(file.Size),
		TransformedStatus:       transformedStatus,
		TransformedRelativePath: TransformedRelativePath,
		TransformedOriginalName: TransformedOriginalName,
		Userid:                  userid,
		CreatedAt:               common.GetCurrentTime(),
	}

	if err = global.GVA_DB.Create(&userTmplFile).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}

	if common.InArray(fileExt, needTransformed2PDFExtConfig) || common.InArray(fileExt, needTransformed2EndXExtConfig) {
		fileType := "pdf"
		if common.InArray(fileExt, needTransformed2EndXExtConfig) {
			fileType = fileExt + "x"
		}

		go func() {
			goUserTmplFile := userTmplFile
			goUserTmplFile.TransformedStatus = fileModel.Transforming
			if err = global.GVA_DB.Model(&fileModel.FileModel{}).Where("id=?", goUserTmplFile.Id).Save(goUserTmplFile).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
			}
			_, err := common.TestToPdf(global.GVA_CONFIG.Upload.Path+dirName+"/"+fileName, global.GVA_CONFIG.Upload.Path+dirName, fileType)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				goUserTmplFile.TransformedStatus = fileModel.Failed
			} else {
				goUserTmplFile.TransformedStatus = fileModel.Succeed
			}
			if err = global.GVA_DB.Model(&fileModel.FileModel{}).Where("id=?", goUserTmplFile.Id).Save(goUserTmplFile).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
			}
		}()
	}

	logForFile(userid, req.WsId, req.TmplId, req.IssueId, file.Filename, "upload_file")

	return nil, nil
}

func (s *UploadSrv) GetFileList(req types.GetFileList) (resp interface{}, err error) {
	type fileVo struct {
		Id                      int    `json:"id"`
		GroupId                 int    `json:"group_id"`
		IsCurrentVersion        int    `json:"is_current_version"`
		VersionCode             int    `json:"version_code"`
		FileName                string `json:"file_name"`
		RelativePath            string `json:"relative_path"`
		TransformedFileName     string `json:"transformed_file_name"`
		TransformedRelativePath string `json:"transformed_relative_path"`
		FileSize                int    `json:"file_size"`
		FileSizeDesc            string `json:"file_size_desc"`
		Creator                 string `json:"creator"`
		CreatedAt               string `json:"created_at"`
	}

	sql_text := "select a.id as id, a.group_id as group_id, a.is_current_version as is_current_version, " +
		"a.version_code as version_code, a.original_name as file_name, " +
		"a.relative_path as relative_path, a.transformed_original_name as transformed_file_name, " +
		"a.transformed_relative_path as transformed_relative_path," +
		"a.file_size as file_size, b.full_name as creator, " +
		"a.created_at as created_at from user_tmpl_file as a left join user as b on a.userid=b.id " +
		"where a.issue_id=@issue_id"
	if req.GroupId > 0 {
		sql_text += " AND a.group_id=@group_id"
	}
	if req.IsCurrentVersion > -1 {
		sql_text += " AND a.is_current_version=@is_current_version"
	}

	sql_text += " ORDER BY a.created_at DESC"

	var fileVoList []fileVo
	global.GVA_DB.Raw(sql_text, sql.Named("issue_id", req.IssueId), sql.Named("group_id", req.GroupId), sql.Named("is_current_version", req.IsCurrentVersion)).Scan(&fileVoList)

	for k, v := range fileVoList {
		fileVoList[k].FileSizeDesc = common.ByteToOther(int64(v.FileSize))
	}

	return fileVoList, nil
}

func (s *UploadSrv) GetFile(userid int, req types.DownloadFile, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.IssueId)
	if err != nil {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"_id": objectID},
		},
	}

	var oldDocument bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	if getDataEditRight(userid, oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	var file fileModel.FileModel
	if err = global.GVA_DB.Where("issue_id=? and id=?", req.IssueId, req.Id).First(&file).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	return file, nil
}

func (s *UploadSrv) UpdateFileIsCurrentVersion(req types.UpdateFileIsCurrentVersion) (resp interface{}, err error) {

	if err = global.GVA_DB.Model(fileModel.FileModel{}).
		Where("issue_id=? and group_id=?", req.IssueId, req.GroupId).
		Update("IsCurrentVersion", 0).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return false, errors.New("操作失败")
	}

	if err = global.GVA_DB.Model(fileModel.FileModel{}).Where("id=?", req.Id).Update("IsCurrentVersion", 1).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return false, errors.New("操作失败")
	}

	return true, nil
}

func (s *UploadSrv) DeleteFile(userid int, req types.DeleteFile, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.IssueId)
	if err != nil {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"_id": objectID},
		},
	}

	var oldDocument bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	if getDataEditRight(userid, oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	idInt, err := strconv.Atoi(req.Id)
	if idInt > 0 {
		var file fileModel.FileModel
		if err = global.GVA_DB.Where("issue_id=? and id=?", req.IssueId, req.Id).First(&file).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误")
		}

		if err = os.Remove(global.GVA_CONFIG.Upload.Path + file.RelativePath); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("删除失败")
		}

		if err = global.GVA_DB.Where("issue_id=? and id=?", req.IssueId, req.Id).Delete(&fileModel.FileModel{}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("删除失败")
		}
		go logForFile(userid, req.WsId, req.TmplId, req.IssueId, file.OriginalName, "delete_file")
	}

	if req.GroupId > 0 {
		var fileList []fileModel.FileModel
		if err = global.GVA_DB.Where("issue_id=? and group_id=?", req.IssueId, req.GroupId).Find(&fileList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误")
		}
		for _, file := range fileList {
			if err = os.Remove(global.GVA_CONFIG.Upload.Path + file.RelativePath); err != nil {
				global.GVA_LOG.Error(err.Error())
				return nil, errors.New("删除失败")
			}
			if err = global.GVA_DB.Where("issue_id=? and id=?", req.IssueId, file.Id).Delete(&fileModel.FileModel{}).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
				return nil, errors.New("删除失败")
			}
			go logForFile(userid, req.WsId, req.TmplId, req.IssueId, file.OriginalName, "delete_file")
		}
	}

	return
}
