package user

import (
	"errors"
	"fmt"
	"io"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/pkg/jwt"
	tmplEnum "mark3/repository/db/enum/tmpl"
	wsEnum "mark3/repository/db/enum/ws"
	tmplModel "mark3/repository/db/model/tmpl"
	model "mark3/repository/db/model/user"
	wsModel "mark3/repository/db/model/ws"
	wsFileModel "mark3/repository/db/model/ws_file"
	"mark3/service/app/tmpl"
	wsfileandtmplmergedsrv "mark3/service/ws_file_and_tmpl_merged"
	types "mark3/types/user"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserSrv struct{}

func (s *UserSrv) Register(req types.RegisterReq) (resp interface{}, err error) {
	tx := global.GVA_DB.Begin()

	//【注册用户】
	var user model.UserModel
	tx.Where("username=?", req.Username).First(&user)
	if user.Id != 0 {
		return nil, errors.New("用户已存在")
	}

	hashPassword, err := common.EncryptPassword(req.Password)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("加密失败")
	}
	user.Username = req.Username
	user.Password = string(hashPassword)
	user.FullName = req.Username
	user.Email = req.Email
	user.CreatedAt = common.GetCurrentTime()
	user.UpdatedAt = common.GetCurrentTime()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	//【创建一个默认空间】
	ws := wsModel.WsModel{
		Name:      "默认空间",
		Desc:      "默认空间",
		Creator:   user.Id,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}
	if err = tx.Create(&ws).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	if err = tx.Create(&wsModel.MemberModel{
		WsId:      ws.Id,
		Userid:    user.Id,
		Role:      wsEnum.WsRoleAdmin,
		Status:    wsEnum.WsStatusOk,
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	//【创建默认模版】
	var tmpl = tmplModel.TmplModel{
		WsId:          ws.Id,
		Name:          "默认流程",
		Desc:          "默认流程",
		WsFileId:      0,
		CreatedUserId: user.Id,
		UpdatedUserId: user.Id,
		Mode:          tmplEnum.TmplModePublic,
		CreatedAt:     common.GetCurrentTime(),
		UpdatedAt:     common.GetCurrentTime(),
	}
	if err = tx.Create(&tmpl).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 流程和文件合并表里面添加ID
	wsFileAndTmplMergedSrv := new(wsfileandtmplmergedsrv.WsFileAndTmplMergedSrv)
	wsFileAndTmplMerged, err := wsFileAndTmplMergedSrv.AddWsFileAndTmplMerged(wsFileModel.WsFileAndTmplMergedModel{
		WsId:          ws.Id,
		ImplTableName: "tmpl",
		ImplTableId:   tmpl.Id,
	})
	if err != nil {
		return nil, err
	}

	// 创建空间文件夹、流程合并排序数据
	if err = tx.Create(&wsFileModel.WsFileAndTmplMergedCoordinateModel{
		WsId:       ws.Id,
		Coordinate: "," + strconv.Itoa(wsFileAndTmplMerged.Id) + ",",
		CreatedAt:  common.GetCurrentTime(),
		UpdatedAt:  common.GetCurrentTime(),
	}).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	//初始化默认字段
	var defaultFields []tmplModel.FieldModel
	for _, field := range append(tmplEnum.FieldLevel2List, tmplEnum.FieldLevel1List...) {
		field.WsId = ws.Id
		field.TmplId = tmpl.Id
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
	}
	// 初始化字段level3： 优先级（高，中，低），截止日期，描述（textarea_com）
	var defaultFieldsL3 []tmplModel.FieldModel
	for _, field := range tmplEnum.FieldLevel3List {
		field.WsId = ws.Id
		field.TmplId = tmpl.Id
		field.CreatedAt = common.GetCurrentTime()
		field.UpdatedAt = common.GetCurrentTime()
		defaultFields = append(defaultFields, field)
		defaultFieldsL3 = append(defaultFieldsL3, field)
	}

	if err = tx.Create(&defaultFields).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return
	}

	//初始化界面
	for module := range tmplEnum.ScreenModuleMap {
		var defaultScreenList []tmplModel.ScreenModel
		for _, field := range tmplEnum.FieldLevel2List {
			var screen = tmplModel.ScreenModel{
				WsId:      ws.Id,
				TmplId:    tmpl.Id,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: tmplEnum.ScreenCanModifyYes,
				Required:  tmplEnum.ScreenRequiredYes,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		for _, field := range defaultFieldsL3 {
			var screen = tmplModel.ScreenModel{
				WsId:      ws.Id,
				TmplId:    tmpl.Id,
				FieldKey:  field.FieldKey,
				Module:    module,
				CanModify: tmplEnum.ScreenCanModifyYes,
				Required:  tmplEnum.ScreenRequiredNo,
				CreatedAt: common.GetCurrentTime(),
				UpdatedAt: common.GetCurrentTime(),
			}
			defaultScreenList = append(defaultScreenList, screen)
		}
		if err = tx.Create(&defaultScreenList).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return
		}
	}
	//初始化状态
	var statusIdList []int
	for _, v := range tmplEnum.InitStatusList {
		var status = tmplModel.StatusModel{
			WsId:      ws.Id,
			TmplId:    tmpl.Id,
			Name:      v.Name,
			Color:     v.Color,
			Mode:      v.Mode,
			CreatedAt: common.GetCurrentTime(),
			UpdatedAt: common.GetCurrentTime(),
		}
		if err = tx.Create(&status).Error; err != nil {
			tx.Rollback()
			global.GVA_LOG.Error(err.Error())
			return
		}
		statusIdList = append(statusIdList, status.Id)
	}
	//初始化步骤
	var stepList []tmplModel.StepModel
	for i := 0; i < len(statusIdList); i++ {
		for j := 0; j < len(statusIdList); j++ {
			if i != j {
				stepList = append(stepList, tmplModel.StepModel{
					WsId:          ws.Id,
					TmplId:        tmpl.Id,
					StartStatusId: statusIdList[i],
					EndStatusId:   statusIdList[j],
					CreatedAt:     common.GetCurrentTime(),
					UpdatedAt:     common.GetCurrentTime(),
				})
			}
		}
	}
	if err = tx.Create(&stepList).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return
	}
	// 初始化权限
	var defaultTmplAuth []tmplModel.TmplAuthModel
	for _, auth := range tmplEnum.AuthList {
		auth.WsId = ws.Id
		auth.TmplId = tmpl.Id
		defaultTmplAuth = append(defaultTmplAuth, auth)
	}
	if err = tx.Create(&defaultTmplAuth).Error; err != nil {
		tx.Rollback()
		global.GVA_LOG.Error(err.Error())
		return
	}

	tx.Commit()

	return
}

func (s *UserSrv) Login(req types.LoginReq) (resp interface{}, err error) {
	var user model.UserModel
	if err = global.GVA_DB.Where("username=?", req.Username).Find(&user).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("登录失败")
	}
	if err = common.ValidatePassword([]byte(user.Password), req.Password); err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("登录失败")
	}

	token, err := jwt.GenerateToken(user.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("登录失败")
	}

	resp = types.LoginResp{
		User:  user,
		Token: token,
	}

	return
}

func (s *UserSrv) Info(userid int) (resp interface{}, err error) {
	var user model.UserModel
	global.GVA_DB.Where("id=?", userid).Find(&user)
	return user, nil
}

func (s *UserSrv) Update(userid int, req types.UpdateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=?", userid).Updates(model.UserModel{
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *UserSrv) UploadAvatar(userid int, file *multipart.FileHeader) (resp interface{}, err error) {
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
		return nil, fmt.Errorf("只能上传图片（文件格式：%s）", global.GVA_CONFIG.Upload.ImageExt)
	}
	src, err := file.Open()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	defer src.Close()

	dirName := "/image/avatar/" + strconv.Itoa(userid)
	fileName := uuid.New().String() + fileExt

	if _, err = os.Stat(global.GVA_CONFIG.Upload.Path + "/" + dirName); os.IsNotExist(err) {
		if err = os.MkdirAll(global.GVA_CONFIG.Upload.Path+"/"+dirName, os.ModePerm); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("创建目录失败")
		}
	}

	out, err := os.Create(fmt.Sprintf("%s%s/%s", global.GVA_CONFIG.Upload.Path, dirName, fileName))
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

	if err = global.GVA_DB.Where("id=?", userid).Updates(model.UserModel{
		Avatar:    dirName + "/" + fileName,
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("上传失败")
	}
	return dirName + "/" + fileName, nil
}

func (s *UserSrv) UpdatePwd(userid int, req types.UpdatePwdReq) (resp interface{}, err error) {
	oldPassword := req.OldPassword
	newPassword := req.NewPassword
	var user model.UserModel
	global.GVA_DB.Where("id=?", userid).Find(&user)

	if err = common.ValidatePassword([]byte(user.Password), oldPassword); err != nil {
		return nil, errors.New("当前密码错误")
	}

	newHashPassword, _ := common.EncryptPassword(newPassword)
	if err = global.GVA_DB.Where("id=?", userid).Updates(model.UserModel{
		Password:  string(newHashPassword),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *UserSrv) SendEmailCode(email string) (resp interface{}, err error) {
	code := common.RandCode(6)
	subject := "邮箱验证码"
	body := "你的验证码为：<h1>" + code + "</h1>"
	to := []string{email}
	if err := tmpl.SendEmail(to, nil, nil, subject, body, ""); err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	rk := email + code
	err = global.GVA_RDB.Set(rk, code, time.Minute*5).Err()
	if err != nil {
		return nil, errors.New("验证码缓存失败")
	}

	return
}

func (s *UserSrv) CheckEmailCode(rk string) (resp interface{}, err error) {
	exists, err := global.GVA_RDB.Exists(rk).Result()
	if err != nil {
		return nil, errors.New("验证码验证失败")
	}
	if exists == 0 {
		return nil, errors.New("验证码已失效，请重新发送")
	}
	return
}

func (s *UserSrv) RegisterEmailCode(req types.EmailParamsReq) (resp interface{}, err error) {
	var user model.UserModel
	global.GVA_DB.Where("email=?", req.Email).First(&user)
	if user.Id != 0 {
		return nil, errors.New("邮箱已存在")
	}
	_, err = s.SendEmailCode(req.Email)
	if err != nil {
		return nil, err
	}
	return
}

func (s *UserSrv) RegisterV2(req types.RegisterV2Req) (resp interface{}, err error) {
	// 校验验证码
	rk := req.Email + req.Code
	_, err = s.CheckEmailCode(rk)
	if err != nil {
		return nil, err
	}
	// 注册
	reqRegister := types.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	_, err = s.Register(reqRegister)
	if err != nil {
		return nil, err
	}

	return
}

func (s *UserSrv) ResetPwdEmailCode(req types.ResetPwdEmailCodeReq) (resp interface{}, err error) {
	var user model.UserModel
	global.GVA_DB.Where("username=?", req.Username).First(&user)
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}
	if user.Email == "" {
		return nil, errors.New("邮箱不存在")
	}
	_, err = s.SendEmailCode(user.Email)
	if err != nil {
		return nil, err
	}
	return
}

func (s *UserSrv) ResetPwd(req types.ResetPwdReq) (resp interface{}, err error) {
	var user model.UserModel
	global.GVA_DB.Where("username=?", req.Username).First(&user)
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}
	if user.Email == "" {
		return nil, errors.New("邮箱不存在")
	}
	// 校验验证码
	rk := user.Email + req.Code
	_, err = s.CheckEmailCode(rk)
	if err != nil {
		return nil, err
	}

	// 修改密码
	newHashPassword, _ := common.EncryptPassword(req.Password)
	if err = global.GVA_DB.Where("id=?", user.Id).Updates(model.UserModel{
		Password:  string(newHashPassword),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *UserSrv) UpdateEmail(userid int, req types.UpdateEmailReq) (resp interface{}, err error) {
	// 校验验证码
	rk := req.Email + req.Code
	_, err = s.CheckEmailCode(rk)
	if err != nil {
		return nil, err
	}
	// 更新
	if err = global.GVA_DB.Where("id=?", userid).Updates(model.UserModel{
		Email:     req.Email,
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
