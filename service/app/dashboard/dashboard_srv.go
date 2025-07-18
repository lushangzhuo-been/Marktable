package dashboard

import (
	"errors"
	"fmt"
	"mark3/const/dashboard"
	"mark3/global"
	model "mark3/repository/db/model/tmpl"
	"strconv"
)

// 服务接口
type BoardService struct {
}

func (d *BoardService) GetBoardConfig(wsID int, tmplID int) (resp map[string]interface{}, err error) {
	srv := DashBoardBaseSrv{}
	tmplFieldList, err := srv.GetTmplFieldList(wsID, tmplID)
	if err != nil {
		return nil, err
	}
	var fieldList []map[string]string
	for _, field := range tmplFieldList {
		field := map[string]string{
			"label":           field.Name,
			"value":           field.FieldKey,
			"mode":            field.Mode,
			"related_tmpl_id": strconv.Itoa(field.RelatedTmplId),
		}
		fieldList = append(fieldList, field)
	}
	res := make(map[string]interface{})
	res["field_list"] = fieldList
	constant := dashboard.TmplBoardConstant{}
	res["statistics_func"] = constant.GetBoardYaxisList()
	res["date_mode"] = []map[string]string{
		{"label": "年", "value": "year"},
		{"label": "月", "value": "month"},
		{"label": "日", "value": "day"},
	}
	res["yes_no"] = []map[string]string{
		{"label": "是", "value": "yes"},
		{"label": "否", "value": "no"},
	}
	return res, nil
}

func (d *BoardService) AddBoard(board model.BoardModel) (resp interface {
}, err error) {
	if err := global.GVA_DB.Save(&board).Error; err != nil {
		global.GVA_DB.Rollback()
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return board, nil
}

func (d *BoardService) UpdateBoard(board model.BoardModel) (resp model.BoardModel, err error) {
	if err = global.GVA_DB.Model(&model.BoardModel{}).Where("id=?", board.Id).Save(board).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return board, errors.New("操作失败")
	}
	return board, nil
}

func (d *BoardService) DeleteBoard(Id int) (resp interface {
}, err error) {
	if err = global.GVA_DB.Where("id=?", Id).Delete(model.BoardModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return true, nil
}

func (d *BoardService) SelectBoard(Id int) (resp model.BoardModel, err error) {
	var board model.BoardModel
	global.GVA_DB.Where("id=?", Id).First(&board)
	return board, nil
}

func (d *BoardService) SelectListBoard(WsId int, TmplId int, ViewId int) (resp []model.BoardModel, err error) {
	var boards []model.BoardModel
	global.GVA_DB.Model(&model.BoardModel{}).Where("ws_id=? and tmpl_id=? and view_id=?", WsId, TmplId, ViewId).Find(&boards)
	return boards, nil
}

func (d *BoardService) CopyBoard(boardMap map[string]int, userId int) (resp interface {
}, err error) {
	board, err := d.SelectBoard(boardMap["id"])
	if err != nil {
		return model.BoardModel{}, err
	}
	if board.Userid != userId {
		return model.BoardModel{}, errors.New("数据不存在")
	}
	NewBoard := model.BoardModel{
		WsId:          board.WsId,
		TmplId:        board.TmplId,
		ViewId:        board.ViewId,
		Userid:        board.Userid,
		Title:         board.Title + "(复制)",
		Mode:          board.Mode,
		XAxis:         board.XAxis,
		AxisDateMode:  board.AxisDateMode,
		YAxis:         board.YAxis,
		GroupAxis:     board.GroupAxis,
		GroupDateMode: board.GroupDateMode,
		Order:         board.Order,
		ShowEmpty:     board.ShowEmpty,
		Filter:        board.Filter,
		W:             boardMap["w"],
		H:             boardMap["h"],
		X:             boardMap["x"],
		Y:             boardMap["y"],
	}
	return d.AddBoard(NewBoard)
}

type BoardLocation struct {
	Id int `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	W  int `json:"w" gorm:"column:w;comment:宽度;type:int(50);"`
	H  int `json:"h" gorm:"column:h;comment:高度;type:int(50);"`
	X  int `json:"x" gorm:"column:x;comment:X坐标;type:int(50);"`
	Y  int `json:"y" gorm:"column:y;comment:Y坐标;type:int(50);"`
}

type BoardLocations struct {
	Locations []BoardLocation
}

func (d *BoardService) ModifyLocation(locations BoardLocations) error {
	for _, location := range locations.Locations {
		// 使用 gorm 的 ToMap 函数（如果可用），或手动转换
		err := global.GVA_DB.Model(&model.BoardModel{}).
			Where("id = ?", location.Id).
			Updates(map[string]interface{}{
				"w": location.W,
				"h": location.H,
				"x": location.X,
				"y": location.Y,
			}).Error
		if err != nil {
			return fmt.Errorf("更新看板位置失败(ID %d): %v", location.Id, err)
		}
	}
	return nil
}
