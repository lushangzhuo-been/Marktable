package dashboard

import (
	"mark3/global"
	appModel "mark3/repository/db/model/app/tmpl"
	model "mark3/repository/db/model/tmpl"
)

type DashBoardBaseSrv struct{}

func (d *DashBoardBaseSrv) GetTmplFieldList(wsID int, tmplID int) (resp []model.FieldModel, err error) {
	var tmplList []model.FieldModel
	global.GVA_DB.Model(model.FieldModel{}).Where("ws_id=? and tmpl_id=? and mode in (\"status_com\", \"select_com\", \"multi_select_com\", \"person_com\", \"related_com\", \"time_com\", \"date_com\")", wsID, tmplID).Find(&tmplList)
	return tmplList, err
}

// GetTmplFieldDictAll 查询流程字段信息
func (s *BoardStatisticsService) GetTmplFieldDictAll(wsID int, tmplID int) (resp []model.FieldModel, err error) {
	var tmplList []model.FieldModel
	global.GVA_DB.Model(model.FieldModel{}).Where("ws_id=? and tmpl_id=?", wsID, tmplID).Find(&tmplList)
	return tmplList, err
}

// GetUserViewById 查询视图信息
func (s *BoardStatisticsService) GetUserViewById(userViewID int) (resp []appModel.UserViewModel, err error) {
	var userViewList []appModel.UserViewModel
	global.GVA_DB.Model(appModel.UserViewModel{}).Where("id=?", userViewID).Find(&userViewList)
	return userViewList, err
}
