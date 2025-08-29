package tmpl

import (
	"encoding/json"
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/tmpl"
	"mark3/repository/db/model/app/tmpl"
	model "mark3/repository/db/model/tmpl"
	tmpl2 "mark3/service/tmpl"
	types "mark3/types/app/tmpl"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type ViewSrv struct{}

func (s *ViewSrv) GetUserViews(userid int, req types.GetUserViewsReq) (resp interface{}, err error) {
	var views []tmpl.UserViewModel
	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and userid=? order by id asc", req.WsId, req.TmplId, userid).Find(&views).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var sortViews []tmpl.UserViewModel
	var userViewCoordinate tmpl.UserViewCoordinateModel

	if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and userid=?", req.WsId, req.TmplId, userid).First(&userViewCoordinate).Error; err == nil {
		coordinate := strings.Split(userViewCoordinate.Coordinate, ",")

		for _, viewIdStr := range coordinate {
			viewId, err := strconv.Atoi(viewIdStr)
			if err != nil {
				continue
			}
			for index, view := range views {
				if view.Id == viewId {
					sortViews = append(sortViews, view)
					views = append(views[:index], views[index+1:]...)
					break
				}
			}
		}
	}
	sortViews = append(sortViews, views...)

	var isPinSign string = "yes"
	for _, v := range sortViews {
		if v.Pin == "yes" {
			isPinSign = "no"
			break
		}
	}

	//获取主视图
	mainView := tmpl.UserViewModel{
		Id:        -1,
		WsId:      req.WsId,
		TmplId:    req.WsId,
		Userid:    userid,
		Name:      "主列表",
		Mode:      "list",
		Filter:    "",
		Lor:       "",
		SortOrder: "",
		Columns:   "",
		Pin:       isPinSign,
	}
	sortViews = append([]tmpl.UserViewModel{mainView}, sortViews...)
	return sortViews, nil
}

func filter2FormatNew(filter string, fieldsMap map[string]model.FieldModel) (formatFilter string) {
	type FilterAll struct {
		FieldType string        `json:"filter_type"`
		Lor       string        `json:"lor"`
		Data      []interface{} `json:"data"`
	}
	var filterAllList []FilterAll
	if err := json.Unmarshal([]byte(filter), &filterAllList); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	var formatFilterResList []FilterAll
	for _, all := range filterAllList {
		var formatFilterRes FilterAll
		filterStr, err := json.Marshal(all.Data)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		filterStrOne := filter2Format(string(filterStr), fieldsMap)
		var filterJson []interface{}
		if err := json.Unmarshal([]byte(filterStrOne), &filterJson); err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		formatFilterRes.FieldType = all.FieldType
		formatFilterRes.Lor = all.Lor
		formatFilterRes.Data = filterJson
		formatFilterResList = append(formatFilterResList, formatFilterRes)
	}
	res, _ := json.Marshal(formatFilterResList)
	return string(res)
}

func filter2Format(filter string, fieldsMap map[string]model.FieldModel) (formatFilter string) {
	type FilterAtom struct {
		FieldKey string      `json:"field_key"`
		Op       string      `json:"op"`
		Value    interface{} `json:"value"`
	}
	var filterAtomList []FilterAtom
	if err := json.Unmarshal([]byte(filter), &filterAtomList); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var formatFilterAtomList []FilterAtom
	for _, filterAtom := range filterAtomList {
		field, ok := fieldsMap[filterAtom.FieldKey]
		if !ok {
			continue
		}
		var formatFilterAtom FilterAtom
		switch field.Mode {
		case enum.FieldPersonCom:
			userid := int(filterAtom.Value.(float64))
			userList := GetUserList([]int{userid})

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = userList
		case enum.FieldStatusCom:
			statusId := int(filterAtom.Value.(float64))
			status := GetStatusInfo(statusId)

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = status
		case enum.FieldRelatedCom:
			issueIdStr := filterAtom.Value.(string)
			var relatedData []interface{}
			relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
			if err1 == nil {
				relatedData = append(relatedData, relatedDocument)
			}

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = relatedData
		default:
			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = filterAtom.Value
		}
		formatFilterAtomList = append(formatFilterAtomList, formatFilterAtom)
	}
	res, _ := json.Marshal(formatFilterAtomList)
	return string(res)
}

func sortOrder2Format(sortOrder string, fieldsMap map[string]model.FieldModel) (formatSortOrder string) {
	type SortOrderAtom struct {
		FieldKey  string `json:"field_key"`
		SortOrder int    `json:"sort_order"`
	}
	var sortOrderAtomList []SortOrderAtom
	if err := json.Unmarshal([]byte(sortOrder), &sortOrderAtomList); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var formatSortOrderAtomList []SortOrderAtom
	for _, sortOrderAtom := range sortOrderAtomList {
		_, ok := fieldsMap[sortOrderAtom.FieldKey]
		if !ok {
			continue
		}
		var formatSortOrderAtom SortOrderAtom
		formatSortOrderAtom.FieldKey = sortOrderAtom.FieldKey
		formatSortOrderAtom.SortOrder = sortOrderAtom.SortOrder
		formatSortOrderAtomList = append(formatSortOrderAtomList, formatSortOrderAtom)
	}
	res, _ := json.Marshal(formatSortOrderAtomList)
	return string(res)
}

func columns2Format(columns string, fieldsMap map[string]model.FieldModel) (formatColumns string) {
	type ColumnsAtom struct {
		FieldKey      string `json:"field_key"`
		Name          string `json:"name"`
		Mode          string `json:"mode"`
		Expr          string `json:"expr"`
		Show          string `json:"show"`
		RelatedTmplId int    `json:"related_tmpl_id"`
	}
	var columnsAtomList []ColumnsAtom
	if err := json.Unmarshal([]byte(columns), &columnsAtomList); err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var formatColumnsAtomList []ColumnsAtom
	for _, columnsAtom := range columnsAtomList {
		field, ok := fieldsMap[columnsAtom.FieldKey]
		if !ok {
			continue
		}
		var formatColumnsAtom ColumnsAtom
		formatColumnsAtom.FieldKey = field.FieldKey
		formatColumnsAtom.Name = field.Name
		formatColumnsAtom.Mode = field.Mode
		formatColumnsAtom.Expr = field.Expr
		formatColumnsAtom.Show = columnsAtom.Show
		formatColumnsAtom.RelatedTmplId = field.RelatedTmplId
		formatColumnsAtomList = append(formatColumnsAtomList, formatColumnsAtom)
		delete(fieldsMap, columnsAtom.FieldKey)
	}

	if len(fieldsMap) > 0 {
		for _, field := range fieldsMap {
			var formatColumnsAtom ColumnsAtom
			formatColumnsAtom.FieldKey = field.FieldKey
			formatColumnsAtom.FieldKey = field.FieldKey
			formatColumnsAtom.Name = field.Name
			formatColumnsAtom.Mode = field.Mode
			formatColumnsAtom.Expr = field.Expr
			formatColumnsAtom.Show = "yes"
			formatColumnsAtom.RelatedTmplId = field.RelatedTmplId
			formatColumnsAtomList = append(formatColumnsAtomList, formatColumnsAtom)
		}
	}

	res, _ := json.Marshal(formatColumnsAtomList)
	return string(res)
}

func (s *ViewSrv) GetViewInfo(userid int, req types.GetViewInfoReq) (resp interface{}, err error) {
	type View struct {
		tmpl.UserViewModel
		FormatFilter string `json:"format_filter"`
	}

	//获取当前所有字段
	tmplCaller := new(tmpl2.Caller)
	fields, _ := tmplCaller.GetFields(req.TmplId)

	var fieldsMap = make(map[string]model.FieldModel)
	for _, field := range fields {
		fieldsMap[field.FieldKey] = field
	}

	//todo filter,sort,columns要过滤一下已废弃的字段
	//除了filter以外sort、columns都不需要格式化

	var view View
	if req.Id == -1 {
		//获取主视图
		view = View{
			UserViewModel: tmpl.UserViewModel{
				Id:        -1,
				WsId:      req.WsId,
				TmplId:    req.WsId,
				Userid:    userid,
				Name:      "主列表",
				Mode:      "list",
				Filter:    "",
				Lor:       "",
				SortOrder: "",
				Columns:   "",
				Pin:       "",
			},
			FormatFilter: "",
		}
		//获取列表字段, 取还是全量字段
		type ColumnsAtom struct {
			FieldKey      string `json:"field_key"`
			Name          string `json:"name"`
			Mode          string `json:"mode"`
			Expr          string `json:"expr"`
			Show          string `json:"show"`
			RelatedTmplId int    `json:"related_tmpl_id"`
		}
		var columnsAtomList []ColumnsAtom
		for _, field := range fields {
			columns := ColumnsAtom{
				FieldKey:      field.FieldKey,
				Name:          field.Name,
				Mode:          field.Mode,
				Expr:          field.Expr,
				Show:          "yes",
				RelatedTmplId: field.RelatedTmplId,
			}
			columnsAtomList = append(columnsAtomList, columns)
		}
		res, _ := json.Marshal(columnsAtomList)
		view.Columns = string(res)
		return view, nil
	} else {
		if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&view).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, err
		}

		if view.Mode == "board" {
			//组装filter
			if view.Filter != "" {
				view.FormatFilter = filter2FormatNew(view.Filter, fieldsMap)
			}
			//获取列表字段, 取还是全量字段
			type ColumnsAtom struct {
				FieldKey      string `json:"field_key"`
				Name          string `json:"name"`
				Mode          string `json:"mode"`
				Expr          string `json:"expr"`
				Show          string `json:"show"`
				RelatedTmplId int    `json:"related_tmpl_id"`
			}
			var columnsAtomList []ColumnsAtom
			for _, field := range fields {
				columns := ColumnsAtom{
					FieldKey:      field.FieldKey,
					Name:          field.Name,
					Mode:          field.Mode,
					Expr:          field.Expr,
					Show:          "yes",
					RelatedTmplId: field.RelatedTmplId,
				}
				columnsAtomList = append(columnsAtomList, columns)
			}
			res, _ := json.Marshal(columnsAtomList)
			view.Columns = string(res)
		} else {
			//组装filter
			if view.Filter != "" {
				view.FormatFilter = filter2FormatNew(view.Filter, fieldsMap)
			}
			//组装sort_order
			if view.SortOrder != "" {
				view.SortOrder = sortOrder2Format(view.SortOrder, fieldsMap)
			}
			//组装columns
			if view.Columns != "" {
				view.Columns = columns2Format(view.Columns, fieldsMap)
			} else {
				//获取列表字段, 取还是全量字段
				type ColumnsAtom struct {
					FieldKey      string `json:"field_key"`
					Name          string `json:"name"`
					Mode          string `json:"mode"`
					Expr          string `json:"expr"`
					Show          string `json:"show"`
					RelatedTmplId int    `json:"related_tmpl_id"`
				}
				var columnsAtomList []ColumnsAtom
				for _, field := range fields {
					columns := ColumnsAtom{
						FieldKey:      field.FieldKey,
						Name:          field.Name,
						Mode:          field.Mode,
						Expr:          field.Expr,
						Show:          "yes",
						RelatedTmplId: field.RelatedTmplId,
					}
					columnsAtomList = append(columnsAtomList, columns)
				}
				res, _ := json.Marshal(columnsAtomList)
				view.Columns = string(res)
			}
		}
		return view, nil
	}
}

func (s *ViewSrv) CreateView(userid int, req types.CreateViewReq) (resp interface{}, err error) {
	var userView = tmpl.UserViewModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		Userid:    userid,
		Name:      req.Name,
		Mode:      req.Mode,
		Filter:    req.Filter,
		Lor:       req.Lor,
		SortOrder: req.SortOrder,
		Columns:   req.Columns,
		Pin:       "no",
		CardGroup: req.CardGroup,
	}
	if err = global.GVA_DB.Create(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) RenameView(userid int, req types.RenameViewReq) (resp interface{}, err error) {
	var userView tmpl.UserViewModel
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	userView.Name = req.Name
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).Updates(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) UpdateViewFilter(userid int, req types.UpdateViewFilterReq) (resp interface{}, err error) {
	var userView tmpl.UserViewModel
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	userView.Filter = req.Filter
	userView.Lor = req.Lor
	if err = global.GVA_DB.Save(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) UpdateViewSortOrder(userid int, req types.UpdateViewSortOrderReq) (resp interface{}, err error) {
	var userView tmpl.UserViewModel
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	userView.SortOrder = req.SortOrder
	if err = global.GVA_DB.Save(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) UpdateViewColumns(userid int, req types.UpdateViewColumnsReq) (resp interface{}, err error) {
	var userView tmpl.UserViewModel
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	userView.Columns = req.Columns
	if err = global.GVA_DB.Save(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) UpdateViewCardGroup(userid int, req types.UpdateViewCardGroupReq) (resp interface{}, err error) {
	var userView tmpl.UserViewModel
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).First(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	userView.CardGroup = req.CardGroup
	if err = global.GVA_DB.Save(&userView).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) DeleteView(userid int, req types.DeleteViewReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).Delete(&tmpl.UserViewModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	return
}

func (s *ViewSrv) PinView(userid int, req types.PinViewReq) (resp interface{}, err error) {
	tx := global.GVA_DB.Begin()
	if err = tx.Model(&tmpl.UserViewModel{}).Where("ws_id=? and tmpl_id=? and userid=?", req.WsId, req.TmplId, userid).Update("pin", "no").Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return
	}

	if err = tx.Model(&tmpl.UserViewModel{}).Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).Update("pin", "yes").Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (s *ViewSrv) UnPinView(userid int, req types.UnPinViewReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Model(&tmpl.UserViewModel{}).Where("id=? and ws_id=? and tmpl_id=? and userid=?", req.Id, req.WsId, req.TmplId, userid).Update("pin", "no").Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	return
}

func (s *ViewSrv) SetViewCoordinate(userid int, req types.SetViewCoordinateReq) (resp interface{}, err error) {
	var userViewCoordinate tmpl.UserViewCoordinateModel
	err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and userid=?", req.WsId, req.TmplId, userid).First(&userViewCoordinate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userViewCoordinate.WsId = req.WsId
			userViewCoordinate.TmplId = req.TmplId
			userViewCoordinate.Coordinate = req.Coordinate
			userViewCoordinate.Userid = userid
			if err = global.GVA_DB.Create(&userViewCoordinate).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
				return nil, err
			}
			return
		}
		global.GVA_LOG.Error(err.Error())
	} else {
		userViewCoordinate.Coordinate = req.Coordinate
		if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and userid=?", req.WsId, req.TmplId, userid).Updates(&userViewCoordinate).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, err
		}
	}
	return
}
