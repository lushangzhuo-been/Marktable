package http

import (
	"github.com/gin-gonic/gin"
	"mark3/global"
	"mark3/internal/modules/ws/application/usecase"
	domain_mysql "mark3/internal/modules/ws/domain/repository/mysql"
	domain_service "mark3/internal/modules/ws/domain/service"
	"mark3/middleware"
)

// 加载依赖以及接口
func Init(r *gin.Engine) {

	wsRepository := domain_mysql.NewWsRepository(global.GVA_DB)
	wsService := domain_service.NewWsService(wsRepository)
	memberRepository := domain_mysql.NewMemberRepository(global.GVA_DB)
	memberService := domain_service.NewMemberService(memberRepository)
	wsFileRepository := domain_mysql.NewWsFileRepository(global.GVA_DB)
	wsFileService := domain_service.NewWsFolderService(wsFileRepository)
	creatWsUseCase := usecase.NewCreatWsUseCase(wsService, wsFileService, memberService)
	wsApi := NewWsHandler(creatWsUseCase)

	r.POST("/ws/create", middleware.AuthMiddleware(), wsApi.CreatWs)

}
