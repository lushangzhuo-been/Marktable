package usecase

import (
	"github.com/gin-gonic/gin"
	"mark3/global"
	"mark3/internal/i18n"
	tmplDto "mark3/internal/modules/tmpl/application/dto"
	"mark3/internal/modules/tmpl/application/enum"
	tmplDomainService "mark3/internal/modules/tmpl/domain/service"
	"mark3/internal/modules/user/application/dto"
	userDomainService "mark3/internal/modules/user/domain/service"
	wsDto "mark3/internal/modules/ws/application/dto"
	"mark3/internal/modules/ws/domain/entity"
	wsDomainService "mark3/internal/modules/ws/domain/service"
	wsFileModel "mark3/repository/db/model/ws_file"
	"strconv"
)

// RegisterUserUseCase 创建空间
type RegisterUserUseCase struct {
	userService     *userDomainService.UserService
	tmplService     *tmplDomainService.TmplService
	statusService   *tmplDomainService.StatusService
	stepService     *tmplDomainService.StepService
	screenService   *tmplDomainService.ScreenService
	fieldService    *tmplDomainService.FieldService
	wsService       *wsDomainService.WsService
	wsFolderService *wsDomainService.WsFolderService
	memberService   *wsDomainService.MemberService
}

func NewRegisterUserUseCase(
	userService *userDomainService.UserService,
	tmplService *tmplDomainService.TmplService,
	fieldService *tmplDomainService.FieldService,
	statusService *tmplDomainService.StatusService,
	stepService *tmplDomainService.StepService,
	screenService *tmplDomainService.ScreenService,
	wsService *wsDomainService.WsService,
	wsFileService *wsDomainService.WsFolderService,
	memberService *wsDomainService.MemberService,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userService:     userService,
		tmplService:     tmplService,
		fieldService:    fieldService,
		statusService:   statusService,
		stepService:     stepService,
		screenService:   screenService,
		wsService:       wsService,
		wsFolderService: wsFileService,
		memberService:   memberService,
	}
}

// Execute 执行用户注册
func (uc *RegisterUserUseCase) Execute(req dto.RegisterReq, c *gin.Context) (resp interface{}, err error) {
	//1. 注册用户
	user, err := uc.userService.Register(req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 2. 空间创建
	ws, err := uc.wsService.CreateWs(user.Id, wsDto.WsCreateReq{
		Name: i18n.FromGinContext(c).T("defaultWsName"),
		Desc: i18n.FromGinContext(c).T("defaultWsName"),
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 3. 空间记忆创建
	_, err = uc.memberService.CreateMember(user.Id, ws.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 4. 创建默认流程
	tmpl, err := uc.tmplService.CreateTmpl(user.Id, tmplDto.TmplCreateReq{
		WsId:     ws.Id,
		Name:     i18n.FromGinContext(c).T("defaultTmplName"),
		Desc:     i18n.FromGinContext(c).T("defaultTmplName"),
		Mode:     enum.TmplModePublic,
		WsFileId: 0,
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 5. 流程和文件合并表里面添加ID
	wsFileAndTmplMerged, err := uc.wsFolderService.CreateWsFolderAndTmplMerged(entity.WsFileAndTmplMergedModel(wsFileModel.WsFileAndTmplMergedModel{
		WsId:          ws.Id,
		ImplTableName: "tmpl",
		ImplTableId:   tmpl.Id,
	}))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	//6. 创建空间文件夹、流程合并排序数据
	_, err = uc.wsFolderService.CreateWsFolder(ws.Id, ","+strconv.Itoa(wsFileAndTmplMerged.Id)+",")
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	//7. 初始化默认字段
	defaultFieldsL3, err := uc.fieldService.CreateFields(ws.Id, tmpl.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	//8. 初始化界面
	err = uc.screenService.CreateScreens(ws.Id, tmpl.Id, defaultFieldsL3)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	//9. 初始化状态
	statusIdList, err := uc.statusService.CreateStatus(ws.Id, tmpl.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	//10. 初始化步骤
	err = uc.stepService.CreateSteps(ws.Id, tmpl.Id, statusIdList)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	return ws, nil
}
