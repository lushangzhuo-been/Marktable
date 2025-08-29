package http

import (
	"github.com/gin-gonic/gin"
	"mark3/internal/i18n"
	"mark3/internal/modules/ws/application/dto"
	"mark3/internal/modules/ws/application/usecase"
	"mark3/pkg/ctl"
	"mark3/pkg/log"
)

type WsHandler struct {
	creatWsUseCase *usecase.CreatWsUseCase // 依赖领域服务
}

func NewWsHandler(service *usecase.CreatWsUseCase) *WsHandler {
	return &WsHandler{creatWsUseCase: service}
}

func (h *WsHandler) CreatWs(c *gin.Context) {
	// 1. 解析DTO
	var req dto.WsCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userid, _ := c.Get("userid")
	i := i18n.FromGinContext(c)
	title := i.T("title")
	log.InitLog().Info(title)
	// 3. 调用领域服务（不处理业务逻辑）
	if _, err := h.creatWsUseCase.Execute(userid.(int), req); err != nil {
		ctl.FailWithMessage(err.Error(), c)
		return
	}

	// 4. 返回标准化响应
	ctl.Ok(c)
}
