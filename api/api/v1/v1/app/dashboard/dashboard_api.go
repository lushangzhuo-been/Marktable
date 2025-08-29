package dashboard

import (
	"github.com/gin-gonic/gin"
	"mark3/global"
	"mark3/pkg/ctl"
	model "mark3/repository/db/model/tmpl"
	srv "mark3/service/app/dashboard"
	"mark3/utils"
	"net/http"
)

type DashBoardApi struct{}

// GetBoardConfig 获取看板配置
func (r *DashBoardApi) GetBoardConfig(c *gin.Context) {
	type QueryParams struct {
		WsID   int `form:"ws_id" binding:"required"`
		TmplID int `form:"tmpl_id" binding:"required"`
	}

	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	res, err := b.GetBoardConfig(params.WsID, params.TmplID)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(res, c)
}

// AddBoard 添加看板
func (r *DashBoardApi) AddBoard(c *gin.Context) {
	var board model.BoardModel
	if err := c.ShouldBindJSON(&board); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	res, err := b.AddBoard(board)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}
	s := new(srv.BoardStatisticsService)
	boardDict, _ := utils.StructToMap(res)
	stats, err := s.GetBoardStatisticsPreview(boardDict, global.GVA_DB)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	boardDict["dashboard_data"] = stats
	ctl.OkWithData(boardDict, c)
}

// ModifyBoard 修改看板
func (r *DashBoardApi) ModifyBoard(c *gin.Context) {
	var board model.BoardModel
	if err := c.ShouldBindJSON(&board); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	if _, err := b.UpdateBoard(board); err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(nil, c)
}

// DeleteBoard 删除看板
func (r *DashBoardApi) DeleteBoard(c *gin.Context) {
	type BodyParams struct {
		BoardID int `json:"board_id" binding:"required"`
	}

	var params BodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	if _, err := b.DeleteBoard(params.BoardID); err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(nil, c)
}

// CopyBoard 复制看板
func (r *DashBoardApi) CopyBoard(c *gin.Context) {
	var boardMap map[string]int
	if err := c.ShouldBindJSON(&boardMap); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	userid, _ := c.Get("userid")

	b := new(srv.BoardService)
	res, err := b.CopyBoard(boardMap, userid.(int))
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}
	s := new(srv.BoardStatisticsService)
	boardDict, _ := utils.StructToMap(res)
	stats, err := s.GetBoardStatisticsPreview(boardDict, global.GVA_DB)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	boardDict["dashboard_data"] = stats
	ctl.OkWithData(boardDict, c)
}

// GetBoardByID 获取看板详情
func (r *DashBoardApi) GetBoardByID(c *gin.Context) {
	type QueryParams struct {
		BoardID int `form:"board_id" binding:"required"`
	}

	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	res, err := b.SelectBoard(params.BoardID)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(res, c)
}

// GetBoardList 获取看板列表
func (r *DashBoardApi) GetBoardList(c *gin.Context) {
	type QueryParams struct {
		WsID   int `form:"ws_id" binding:"required"`
		TmplID int `form:"tmpl_id" binding:"required"`
		ViewID int `form:"view_id" binding:"required"`
	}

	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	res, err := b.SelectListBoard(params.WsID, params.TmplID, params.ViewID)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(res, c)
}

// GetBoardStatistics 获取看板统计
func (r *DashBoardApi) GetBoardStatistics(c *gin.Context) {
	type BodyParams struct {
		BoardID    int    `json:"board_id" binding:"required"`
		Filter     string `json:"filter"`
		LOR        string `json:"lor"`
		FilterDown string `json:"filter_down"`
	}

	var params BodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	s := new(srv.BoardStatisticsService)
	res, err := s.GetBoardStatisticsByID(
		params.BoardID,
		params.Filter,
		params.FilterDown,
		params.LOR,
		global.GVA_DB,
	)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(res, c)
}

// GetBoardStatisticsPreview 获取看板统计预览
func (r *DashBoardApi) GetBoardStatisticsPreview(c *gin.Context) {
	var boardDict map[string]interface{}
	if err := c.ShouldBindJSON(&boardDict); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	s := new(srv.BoardStatisticsService)
	res, err := s.GetBoardStatisticsPreview(boardDict, global.GVA_DB)
	if err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(res, c)
}

// ModifyLocation 修改看板位置
func (r *DashBoardApi) ModifyLocation(c *gin.Context) {
	var locations srv.BoardLocations
	if err := c.ShouldBindJSON(&locations); err != nil {
		ctl.FailWithCodeMessage(http.StatusBadRequest, err.Error(), c)
		return
	}
	b := new(srv.BoardService)
	if err := b.ModifyLocation(locations); err != nil {
		ctl.FailWithCodeMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}

	ctl.OkWithData(nil, c)
}
