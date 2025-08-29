package usecase

import (
	"mark3/global"
	"mark3/internal/modules/ws/application/dto"
	domainService "mark3/internal/modules/ws/domain/service"
)

// CreatWsUseCase 创建空间
type CreatWsUseCase struct {
	wsService     *domainService.WsService
	wsFileService *domainService.WsFileService
	memberService *domainService.MemberService
}

func NewCreatWsUseCaseUseCase(
	wsService *domainService.WsService,
	wsFileService *domainService.WsFileService,
	memberService *domainService.MemberService,
) *CreatWsUseCase {
	return &CreatWsUseCase{
		wsService:     wsService,
		wsFileService: wsFileService,
		memberService: memberService,
	}
}

// Execute 执行空间创建
func (uc *CreatWsUseCase) Execute(userid int, req dto.WsCreateReq) (resp interface{}, err error) {
	// 1. 空间创建
	ws, err := uc.wsService.CreateWs(userid, req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 2. 空间记忆创建
	_, err = uc.memberService.CreateMember(userid, ws.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 3. 空间文件夹排序创建
	_, err = uc.wsFileService.CreateWsFile(ws.Id)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	return ws, nil
}
