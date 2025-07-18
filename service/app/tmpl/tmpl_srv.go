package tmpl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/pkg/right"
	enum "mark3/repository/db/enum/tmpl"
	wsEnum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/tmpl"
	wsModel "mark3/repository/db/model/ws"
	"mark3/service/rule_action_log"
	"mark3/service/tmpl"
	"mark3/types/app"
	types "mark3/types/app/tmpl"
	commonTypes "mark3/types/common"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TmplSrv struct{}

func (s *TmplSrv) GetTmplInfo(userid int, req types.GetTmplInfoReq) (resp interface{}, err error) {
	type Res struct {
		Tmpl       model.TmplModel   `json:"tmpl"`
		AdminList  []app.MemberVo    `json:"admin_list"`
		Permission map[string]string `json:"permission"`
	}

	var tmpl model.TmplModel
	if err = global.GVA_DB.Where("id=? and ws_id=?", req.TmplId, req.WsId).First(&tmpl).Error; err != nil {
		return nil, errors.New("参数错误，流程不存在")
	}

	//获取用户权限 todo
	var permission = make(map[string]string)

	if tmpl.Mode == enum.TmplModePublic {
		permission = map[string]string{
			"view_permission": enum.RoleViewPermissionAll,
			"edit_permission": "create,export_list",
		}
		return Res{Tmpl: tmpl, AdminList: nil, Permission: permission}, nil
	} else {
		type Role struct {
			Userid         int    `json:"userid"`
			Sign           string `json:"sign"`
			Name           string `json:"name"`
			ViewPermission string `json:"view_permission"`
			EditPermission string `json:"edit_permission"`
		}
		var role Role
		if err = global.GVA_DB.Raw("select a.userid, b.sign, b.name, b.view_permission, b.edit_permission from tmpl_member as a left join tmpl_role as b on a.role_id=b.id where a.tmpl_id=? and a.userid=?", req.TmplId, userid).First(&role).Error; err != nil {
			permission = map[string]string{
				"view_permission": "",
				"edit_permission": "",
			}
		} else {
			permission = map[string]string{
				"view_permission": role.ViewPermission,
				"edit_permission": role.EditPermission,
			}
		}

		var adminRole model.RoleModel
		if err = global.GVA_DB.Where("tmpl_id=? and sign=?", req.TmplId, enum.RoleSignAdmin).First(&adminRole).Error; err != nil {
			return Res{Tmpl: tmpl, AdminList: nil, Permission: permission}, nil
		}

		var adminList []app.MemberVo
		global.GVA_DB.Raw("select a.userid, b.username, b.full_name, b.avatar, b.email, b.phone from tmpl_member as a inner join user as b on a.userid=b.id where a.tmpl_id=? and a.role_id=? and a.status=?", req.TmplId, adminRole.Id, enum.TmplStatusOk).Find(&adminList)
		return Res{Tmpl: tmpl, AdminList: adminList, Permission: permission}, nil
	}
}

func (s *TmplSrv) GetTmplList(userid int, req types.GetTmplListReq) (resp interface{}, err error) {
	type TmplVo struct {
		Id       int    `json:"id"`
		WsId     int    `json:"ws_id"`
		Name     string `json:"name"`
		Mode     string `json:"mode"`
		IsMember string `json:"is_member"`
		IsAdmin  string `json:"is_admin"`
	}

	var tmplList []TmplVo
	global.GVA_DB.Model(model.TmplModel{}).Select("id, ws_id, name, mode").Where("ws_id=?", req.WsId).Find(&tmplList)

	var member wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=? and status=? and role=?", req.WsId, userid, wsEnum.WsStatusOk, wsEnum.WsRoleAdmin).First(&member).Error; err == nil {
		for i, vo := range tmplList {
			vo.IsMember = "yes"
			vo.IsAdmin = "yes"
			tmplList[i] = vo
		}
		return tmplList, nil
	}

	for i, vo := range tmplList {
		if vo.Mode == enum.TmplModePublic {
			vo.IsMember = "yes"
			vo.IsAdmin = "no"
		} else {
			var member model.MemberModel
			if err := global.GVA_DB.Where("tmpl_id=? and userid=? and status=?", vo.Id, userid, enum.TmplStatusOk).First(&member).Error; err != nil {
				vo.IsMember = "no"
			} else {
				vo.IsMember = "yes"
			}

			var role model.RoleModel
			if err := global.GVA_DB.Where("tmpl_id=? and id=? and sign=?", vo.Id, member.RoleId, enum.RoleSignAdmin).First(&role).Error; err != nil {
				vo.IsAdmin = "no"
			} else {
				vo.IsAdmin = "yes"
			}
		}
		tmplList[i] = vo
	}
	return tmplList, nil
}

func (s *TmplSrv) GetTmplOtherList(req types.GetTmplOtherListReq) (resp interface{}, err error) {
	type TmplVo struct {
		Id   int    `json:"id"`
		WsId int    `json:"ws_id"`
		Name string `json:"name"`
		Mode string `json:"mode"`
	}

	var tmplList []TmplVo
	global.GVA_DB.Model(model.TmplModel{}).Select("id, ws_id, name, mode").Where("ws_id=? and id !=?", req.WsId, req.TmplId).Find(&tmplList)

	return tmplList, nil
}

func (s *TmplSrv) GetConfig(req types.GetConfigReq) (resp interface{}, err error) {
	tmplCaller := new(tmpl.Caller)
	fields, err := tmplCaller.GetFields(req.TmplId)
	if err != nil {
		return
	}

	type Item struct {
		FieldKey      string                   `json:"field_key"`
		Name          string                   `json:"name"`
		Mode          string                   `json:"mode"`
		Expr          string                   `json:"expr"`
		EnumValues    string                   `json:"enum_values"`
		RelatedTmplId int                      `json:"related_tmpl_id"`
		Op            []commonTypes.BaseConfig `json:"op"`
	}
	opMap := enum.FieldModeOpMap
	var items []Item
	for _, field := range fields {
		op, ok := opMap[field.Mode]
		if !ok {
			continue
		}
		item := Item{
			FieldKey:      field.FieldKey,
			Name:          field.Name,
			Mode:          field.Mode,
			Expr:          field.Expr,
			EnumValues:    field.EnumValues,
			RelatedTmplId: field.RelatedTmplId,
			Op:            op,
		}
		if field.Mode == enum.FieldStatusCom {
			statusList, _ := tmplCaller.GetStatusList(req.TmplId)
			if statusList != nil {
				statusStr, sErr := json.Marshal(statusList)
				if sErr == nil {
					item.EnumValues = string(statusStr)
				}
			}
		}
		items = append(items, item)
	}

	for _, field := range enum.FieldLevel1List {
		mode := ""
		enumValues := ""
		op, ok := opMap[mode]
		if !ok {
			continue
		}
		item := Item{
			FieldKey:      field.FieldKey,
			Name:          field.Name,
			Mode:          mode,
			Expr:          field.Expr,
			EnumValues:    enumValues,
			RelatedTmplId: field.RelatedTmplId,
			Op:            op,
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *TmplSrv) GetStatusList(req types.GetStatusListReq) (resp interface{}, err error) {
	tmplCaller := new(tmpl.Caller)
	statusList, err := tmplCaller.GetStatusList(req.TmplId)
	return statusList, err
}

func (s *TmplSrv) GetScreen(req types.GetScreenReq) (resp interface{}, err error) {
	tmplCaller := new(tmpl.Caller)
	var module string
	if req.Module == "list" {
		module = "detail"
	} else {
		module = req.Module
	}
	screenList, _ := tmplCaller.GetScreenByModule(req.TmplId, module)

	if req.Module == "list" {
		columns := req.Columns
		if columns == "" {
			return screenList, nil
		}
		type ColumnsAtom struct {
			FieldKey      string `json:"field_key"`
			Name          string `json:"name"`
			Mode          string `json:"mode"`
			Expr          string `json:"expr"`
			RelatedTmplId int    `json:"related_tmpl_id"`
			Show          string `json:"show"`
		}
		var columnsAtomList []ColumnsAtom
		if err = json.Unmarshal([]byte(columns), &columnsAtomList); err != nil {
			global.GVA_LOG.Error(err.Error())
			return screenList, nil
		}
		var res []types.ScreenVo
		for _, columnsAtom := range columnsAtomList {
			for _, s := range screenList {
				if columnsAtom.FieldKey == s.FieldKey && columnsAtom.Show == "yes" {
					res = append(res, s)
				}
			}
		}
		return res, nil
	}
	return screenList, nil
}

func (s *TmplSrv) getFilter(userid int, req types.GetListDataReq, fieldsMap map[string]model.FieldModel, userTmplRightInfo right.UserTmplRight) (filter bson.M, err error) {
	filter = bson.M{
		"ws_id":   req.WsId,
		"tmpl_id": req.TmplId,
	}

	//获取数据权限
	if userTmplRightInfo.TmplMode == enum.TmplModePrivate {
		if userTmplRightInfo.WsRoleSign == wsEnum.WsRoleGeneral {
			if userTmplRightInfo.TmplViewPermission == enum.RoleViewPermissionOnlyMe {
				filter["$or"] = []bson.M{
					{"creator": userid},
					{"handler": userid},
				}
			}
		}
	}

	type filterAtom struct {
		FieldKey string      `json:"field_key"`
		Op       string      `json:"op"`
		Value    interface{} `json:"value"`
	}

	//增加下钻过滤
	filterDownInner := s.getFilterDown(req.FilterDown)
	if len(filterDownInner) > 0 {
		filter["$and"] = filterDownInner
	}

	var filterList []filterAtom
	if err = json.Unmarshal([]byte(req.Filter), &filterList); err != nil {
		return filter, nil
	}
	if len(filterList) == 0 {
		return filter, nil
	}

	var filterInner bson.A
	var cond bson.M
	for _, atom := range filterList {
		if atom.FieldKey != "" && atom.Op != "" {
			if field, ok := fieldsMap[atom.FieldKey]; !ok {
				continue
			} else {
				//如果字段类型为文本类型的，需要特殊处理一下
				if field.Mode == enum.FieldTextCom || field.Mode == enum.FieldTextareaCom {
					cond = bson.M{atom.FieldKey: bson.M{"$regex": atom.Value, "$options": "i"}}
				} else if field.Mode == enum.FieldLinkCom {
					key := fmt.Sprintf("%s.%s", atom.FieldKey, "name")
					cond = bson.M{key: bson.M{atom.Op: atom.Value}}
				} else {
					cond = bson.M{atom.FieldKey: bson.M{atom.Op: atom.Value}}
				}
			}
			filterInner = append(filterInner, cond)
		}
	}

	if req.Lor == "filter_or" {
		filter["$or"] = filterInner
	} else if len(filterDownInner) > 0 {
		//过滤和图表下钻过滤合并
		filter["$and"] = append(filterDownInner, filterInner...)
	} else {
		filter["$and"] = filterInner
	}
	return filter, nil
}

func (s *TmplSrv) getFilterNew(userid int, req types.GetListDataReq, fieldsMap map[string]model.FieldModel, userTmplRightInfo right.UserTmplRight) (filter bson.M, err error, subTmplEmpty bool) {
	filter = bson.M{
		"ws_id":   req.WsId,
		"tmpl_id": req.TmplId,
	}
	subTmplEmpty = false

	//获取数据权限
	if userTmplRightInfo.TmplMode == enum.TmplModePrivate {
		if userTmplRightInfo.WsRoleSign == wsEnum.WsRoleGeneral {
			if userTmplRightInfo.TmplViewPermission == enum.RoleViewPermissionOnlyMe {
				filter["$or"] = []bson.M{
					{"creator": userid},
					{"handler": userid},
				}
			}
		}
	}

	type filterAll struct {
		FieldType string        `json:"filter_type"`
		Lor       string        `json:"lor"`
		Data      []interface{} `json:"data"`
	}

	// 看板过滤，增加下钻过滤
	filterDownInner := s.getFilterDown(req.FilterDown)
	// 看板过滤，是否显示空值
	if req.BoardId > 0 {
		var board model.BoardModel
		if err := global.GVA_DB.Where("id=?", req.BoardId).First(&board).Error; err != nil {
			return nil, errors.New("参数错误，看板不存在"), subTmplEmpty
		}
		if board.ShowEmpty == "no" {
			if board.XAxis != "" {
				cond := bson.M{
					"$and": []bson.M{
						{board.XAxis: bson.M{"$exists": true}},
						{board.XAxis: bson.M{"$ne": nil}},
						{board.XAxis: bson.M{"$ne": ""}},
						{board.XAxis: bson.M{"$ne": bson.A{}}},
					},
				}
				filterDownInner = append(filterDownInner, cond)
			}
			if board.GroupAxis != "" {
				cond := bson.M{
					"$and": []bson.M{
						{board.GroupAxis: bson.M{"$exists": true}},
						{board.GroupAxis: bson.M{"$ne": nil}},
						{board.GroupAxis: bson.M{"$ne": ""}},
						{board.GroupAxis: bson.M{"$ne": bson.A{}}},
					},
				}
				filterDownInner = append(filterDownInner, cond)
			}
		}
	}
	// 子任务过滤
	if req.ParentTmplId > 0 && req.ObjId != "" {
		var subObjs []model.TmplSubObjModel
		if err = global.GVA_DB.Where("ws_id=? and tmpl_id=? and sub_tmpl_id=? and obj_id=?",
			req.WsId, req.ParentTmplId, req.TmplId, req.ObjId).Find(&subObjs).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误"), subTmplEmpty
		}
		if len(subObjs) > 0 {
			var objectIds []primitive.ObjectID
			for _, obj := range subObjs {
				objectId, err := primitive.ObjectIDFromHex(obj.SubObjId)
				if err != nil {
					continue
				}
				objectIds = append(objectIds, objectId)
			}
			filterDownInner = append(filterDownInner, bson.M{
				"_id": bson.M{
					"$in": objectIds,
				},
			})
		} else {
			subTmplEmpty = true
			return nil, nil, subTmplEmpty
		}
	}
	if len(filterDownInner) > 0 {
		filter["$and"] = filterDownInner
	}

	var filterAllList []filterAll
	if err = json.Unmarshal([]byte(req.Filter), &filterAllList); err != nil {
		return filter, nil, subTmplEmpty
	}

	var filterInner bson.A
	for _, all := range filterAllList {
		filterListOne := s.getOnlyFilter(all.Data, fieldsMap)
		if all.FieldType == "filter" && len(filterListOne) > 0 {
			filterInner = append(filterInner, filterListOne[0])
		} else if all.FieldType == "group" {
			if len(filterListOne) == 1 {
				filterInner = append(filterInner, filterListOne[0])
			} else if len(filterListOne) > 1 {
				if all.Lor == "filter_or" {
					filterInner = append(filterInner, bson.M{"$or": filterListOne})
				} else {
					filterInner = append(filterInner, bson.M{"$and": filterListOne})
				}
			}
		}
	}

	if len(filterInner) == 0 {
		return filter, nil, subTmplEmpty
	}
	if req.Lor == "filter_or" {
		filter["$or"] = filterInner
	} else if len(filterDownInner) > 0 {
		//过滤和图表下钻过滤合并
		filter["$and"] = append(filterDownInner, filterInner...)
	} else {
		filter["$and"] = filterInner
	}
	return filter, nil, subTmplEmpty
}

func (s *TmplSrv) getOnlyFilter(filterList []interface{}, fieldsMap map[string]model.FieldModel) bson.A {
	type filterAtom struct {
		FieldKey string      `json:"field_key"`
		Op       string      `json:"op"`
		Value    interface{} `json:"value"`
	}

	var filterInner bson.A
	var cond bson.M
	for _, atomMap := range filterList {
		atom := filterAtom{}
		arr, err := json.Marshal(atomMap)
		if err != nil {
			continue
		}
		err = json.Unmarshal(arr, &atom)
		if err != nil {
			continue
		}

		if atom.FieldKey != "" && atom.Op != "" {
			if field, ok := fieldsMap[atom.FieldKey]; !ok {
				continue
			} else {
				//如果字段类型为文本类型的，需要特殊处理一下
				if field.Mode == enum.FieldTextCom || field.Mode == enum.FieldTextareaCom {
					cond = bson.M{atom.FieldKey: bson.M{"$regex": atom.Value, "$options": "i"}}
				} else if field.Mode == enum.FieldLinkCom {
					key := fmt.Sprintf("%s.%s", atom.FieldKey, "name")
					cond = bson.M{key: bson.M{atom.Op: atom.Value}}
				} else {
					cond = bson.M{atom.FieldKey: bson.M{atom.Op: atom.Value}}
				}
			}
			filterInner = append(filterInner, cond)
		}
	}
	return filterInner
}

func (s *TmplSrv) getFilterDown(filterDownStr string) bson.A {
	var filterInner bson.A
	var cond bson.M
	type filterDownAtom struct {
		XAxis         string      `json:"x_axis"`
		XValue        interface{} `json:"x_value"`
		XDateMode     string      `json:"axis_date_mode"`
		GroupAxis     string      `json:"group_axis"`
		GroupValue    interface{} `json:"group_value"`
		GroupDateMode string      `json:"group_date_mode"`
	}
	var filterDownMap filterDownAtom
	if err := json.Unmarshal([]byte(filterDownStr), &filterDownMap); err != nil {
		return filterInner
	}

	dateModeMap := map[string]string{
		"day":   "day",
		"month": "month",
		"year":  "year",
	}
	_, xOk := dateModeMap[filterDownMap.XDateMode]
	_, groupOk := dateModeMap[filterDownMap.GroupDateMode]
	if filterDownMap.XAxis != "" {
		if filterDownMap.XValue == "空值" {
			cond = bson.M{
				"$or": []bson.M{
					{filterDownMap.XAxis: bson.M{"$exists": false}},
					{filterDownMap.XAxis: bson.M{"$eq": nil}},
					{filterDownMap.XAxis: bson.M{"$eq": ""}},
					{filterDownMap.XAxis: bson.M{"$eq": bson.A{}}},
				},
			}
			filterInner = append(filterInner, cond)
		} else if xOk {
			startDay, endDay := s.getStartEndDay(filterDownMap.XValue, filterDownMap.XDateMode)
			if startDay != "" && endDay != "" {
				filterInner = append(filterInner, bson.M{filterDownMap.XAxis: bson.M{"$gte": startDay}})
				filterInner = append(filterInner, bson.M{filterDownMap.XAxis: bson.M{"$lt": endDay}})
			}
		} else {
			xValueStr := fmt.Sprintf("%v", filterDownMap.XValue)
			xValueStrNum, err := strconv.Atoi(xValueStr)
			var xValue interface{}
			if err != nil {
				fmt.Println("转换失败:", err)
				xValue = xValueStr
			} else {
				fmt.Println("转换成功:", xValueStrNum)
				xValue = xValueStrNum
			}
			cond = bson.M{filterDownMap.XAxis: bson.M{"$eq": xValue}}
			filterInner = append(filterInner, cond)
		}
	}

	if filterDownMap.GroupAxis != "" {
		if filterDownMap.GroupValue == "空值" {
			cond = bson.M{
				"$or": []bson.M{
					{filterDownMap.GroupAxis: bson.M{"$exists": false}},
					{filterDownMap.GroupAxis: bson.M{"$eq": nil}},
					{filterDownMap.GroupAxis: bson.M{"$eq": ""}},
					{filterDownMap.GroupAxis: bson.M{"$eq": bson.A{}}},
				},
			}
			filterInner = append(filterInner, cond)
		} else if groupOk {
			startDay, endDay := s.getStartEndDay(filterDownMap.GroupValue, filterDownMap.GroupDateMode)
			if startDay != "" && endDay != "" {
				filterInner = append(filterInner, bson.M{filterDownMap.GroupAxis: bson.M{"$gte": startDay}})
				filterInner = append(filterInner, bson.M{filterDownMap.GroupAxis: bson.M{"$lt": endDay}})
			}
		} else {
			groupValueStr := fmt.Sprintf("%v", filterDownMap.GroupValue)

			groupValueNum, err := strconv.Atoi(groupValueStr)
			var groupValue interface{}
			if err != nil {
				fmt.Println("转换失败:", err)
				groupValue = groupValueStr
			} else {
				fmt.Println("转换成功:", groupValueNum)
				groupValue = groupValueNum
			}
			cond = bson.M{filterDownMap.GroupAxis: bson.M{"$eq": groupValue}}
			filterInner = append(filterInner, cond)
		}
	}

	return filterInner
}

func (s *TmplSrv) getSortOrder(req types.GetListDataReq) (sortOrder bson.D, err error) {
	type SortOrderAtom struct {
		FieldKey  string `json:"field_key"`
		SortOrder int    `json:"sort_order"`
	}

	var sortOrderList []SortOrderAtom
	if err = json.Unmarshal([]byte(req.SortOrder), &sortOrderList); err != nil {
		return bson.D{{Key: "created_at", Value: -1}}, nil
	}
	if len(sortOrderList) == 0 {
		return bson.D{{Key: "created_at", Value: -1}}, nil
	}
	for _, atom := range sortOrderList {
		sortOrder = append(sortOrder, bson.E{Key: atom.FieldKey, Value: atom.SortOrder})
	}
	return
}

func (s *TmplSrv) getStartEndDay(dateStr interface{}, dateMode string) (string, string) {
	startStr := dateStr.(string)
	endDay := ""
	if dateMode == "day" {
		startDay, _ := time.ParseInLocation("2006-01-02", startStr, time.Local)
		endDay = startDay.AddDate(0, 0, 1).Format("2006-01-02")
	} else if dateMode == "month" {
		startStr += "-01"
		startDay, _ := time.ParseInLocation("2006-01-02", startStr, time.Local)
		endDay = startDay.AddDate(0, 1, 0).Format("2006-01-02")
	} else if dateMode == "year" {
		startStr += "-01-01"
		startDay, _ := time.ParseInLocation("2006-01-02", startStr, time.Local)
		endDay = startDay.AddDate(1, 0, 0).Format("2006-01-02")
	}
	return startStr, endDay
}

func getPaginationList(filter bson.M, sortOrder bson.D, pageSize int, pageNum int) ([]bson.M, int64, error) {
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	skip := (pageNum - 1) * pageSize

	findOptions := options.Find().
		SetSort(sortOrder).
		SetLimit(int64(pageSize)).
		SetSkip(int64(skip))

	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	cnt, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, 0, err
	}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, 0, err
	}
	return documents, cnt, nil
}

func getAllList(filter bson.M, sortOrder bson.D) ([]bson.M, error) {
	findOptions := options.Find().
		SetSort(sortOrder)

	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}
	return documents, nil
}

func (s *TmplSrv) GetListData(ctx *gin.Context, userid int, req types.GetListDataReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMap(req.TmplId)

	filter, err, subTmplEmpty := s.getFilterNew(userid, req, fieldsMap, userTmplRightInfo)
	if err != nil {
		return nil, err
	}
	if subTmplEmpty {
		resp = commonTypes.BasePageResp{
			List: nil,
			Cnt:  0,
		}
		return
	}

	sortOrder, err := s.getSortOrder(req)
	if err != nil {
		return nil, err
	}

	var documents []bson.M
	var cnt int64

	if req.Export == "yes" {
		documents, err = getAllList(filter, sortOrder)
		if err != nil {
			return nil, err
		}

		var tmpl model.TmplModel
		if err := global.GVA_DB.Where("id=?", req.TmplId).First(&tmpl).Error; err != nil {
			return nil, errors.New("参数错误，流程不存在")
		}
		exportExcel(ctx, tmpl, documents, tmplCaller)
		return
	} else {
		documents, cnt, err = getPaginationList(filter, sortOrder, req.PageSize, req.PageNum)
		if err != nil {
			return nil, err
		}

		relatedTmplMember := make(map[int]string)

		for _, document := range documents {
			for key, value := range document {
				field, ok := fieldsMap[key]
				if !ok {
					continue
				}
				if field.Mode == enum.FieldPersonCom {
					userIds := convertPersonCom(value)
					document[key] = GetUserList(userIds)
				}
				if field.Mode == enum.FieldStatusCom {
					statusId, _ := value.(int32)
					status, _ := tmplCaller.GetStatus(req.TmplId, int(statusId))
					document[key] = status
				}
				if field.Mode == enum.FieldRelatedCom {
					var relatedData []interface{}
					valueAssert := value.(primitive.A)
					for _, issueId := range valueAssert {
						issueIdStr, okr := issueId.(string)
						if okr {
							if _, ok := relatedTmplMember[field.RelatedTmplId]; !ok {
								isMember := s.GetTmplAuth(field.WsId, field.RelatedTmplId, userid)
								relatedTmplMember[field.RelatedTmplId] = isMember
							}

							relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
							if err1 == nil {
								relatedDocument["is_member"] = relatedTmplMember[field.RelatedTmplId]
								relatedData = append(relatedData, relatedDocument)
								break
							}
						}
					}
					document[key] = relatedData
				}
			}
			document["progress_num"] = s.getProgressCount(req.WsId, req.TmplId, document["_id"].(primitive.ObjectID).Hex())
		}

		resp = commonTypes.BasePageResp{
			List: documents,
			Cnt:  cnt,
		}
		return
	}
}

func (s *TmplSrv) GetTmplAuth(wsId int, tmplId int, userid int) (resp string) {
	var member wsModel.MemberModel
	if err := global.GVA_DB.Where("ws_id=? and userid=? and status=? and role=?", wsId, userid, wsEnum.WsStatusOk, wsEnum.WsRoleAdmin).First(&member).Error; err == nil {
		return "yes"
	}
	var tmplInfo model.TmplModel
	if err := global.GVA_DB.Where("id=? and ws_id=?", tmplId, wsId).First(&tmplInfo).Error; err != nil {
		return "no"
	}
	if tmplInfo.Mode == enum.TmplModePublic {
		return "yes"
	} else {
		var member model.MemberModel
		if err := global.GVA_DB.Where("tmpl_id=? and userid=? and status=?", tmplId, userid, enum.TmplStatusOk).First(&member).Error; err != nil {
			return "no"
		} else {
			return "yes"
		}
	}
}

func (s *TmplSrv) GetListDataSelect(req types.GetListDataSelectReq, userid int) (resp interface{}, err error) {
	filter := bson.M{
		"ws_id":   req.WsId,
		"tmpl_id": req.TmplId,
	}
	var filterInner bson.A
	if req.Key != "" {
		cond := bson.M{"title": bson.M{"$regex": req.Key, "$options": "i"}}
		filterInner = append(filterInner, cond)
	}
	if req.Ex != "" {
		ex := strings.Split(req.Ex, ",")
		var objectIds []primitive.ObjectID
		for _, id := range ex {
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				continue
			}
			objectIds = append(objectIds, objectId)
		}
		if len(objectIds) > 0 {
			cond := bson.M{"_id": bson.M{"$nin": objectIds}}
			filterInner = append(filterInner, cond)
		}
	}
	if len(filterInner) > 0 {
		filter["$and"] = filterInner
	}

	sortOrder := bson.D{{Key: "updated_at", Value: -1}}
	var documents []bson.M
	var cnt int64
	documents, cnt, err = getPaginationList(filter, sortOrder, req.PageSize, req.PageNum)
	if err != nil {
		return nil, err
	}

	isMember := s.GetTmplAuth(req.WsId, req.TmplId, userid)
	for _, document := range documents {
		document["is_member"] = isMember
	}

	resp = commonTypes.BasePageResp{
		List: documents,
		Cnt:  cnt,
	}
	return
}

type Permission struct {
	Edit   string `json:"edit"`
	Delete string `json:"delete"`
}

func convertPersonCom(value interface{}) []int {
	val := reflect.ValueOf(value)
	var userIds []int
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i)
		vInt32, ok := v.Interface().(int32)
		if ok {
			userIds = append(userIds, int(vInt32))
		}
	}
	return userIds
}

func getDataEditRight(userid int, document bson.M, userTmplRightInfo right.UserTmplRight) (permission Permission) {
	permission = Permission{Edit: "no", Delete: "no"}

	creatorIds := convertPersonCom(document["creator"])
	handlerIds := convertPersonCom(document["handler"])

	for _, handlerId := range handlerIds {
		if handlerId == userid {
			permission.Edit = "yes"
			break
		}
	}

	for _, creatorId := range creatorIds {
		if creatorId == userid {
			permission.Edit = "yes"
			permission.Delete = "yes"
		}
	}

	if userTmplRightInfo.TmplRoleSign == enum.RoleSignAdmin {
		permission.Edit = "yes"
		permission.Delete = "yes"
	}

	if userTmplRightInfo.WsRoleSign == wsEnum.WsRoleAdmin {
		permission.Edit = "yes"
		permission.Delete = "yes"
	}
	return
}

func (s *TmplSrv) GetFileRight(userid int, req types.GetDataReq, userTmplRightInfo right.UserTmplRight) (permission Permission, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
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

	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return Permission{Edit: "no", Delete: "no"}, errors.New("数据不存在")
	}

	return getDataEditRight(userid, document, userTmplRightInfo), nil
}

func (s *TmplSrv) GetData(userid int, req types.GetDataReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
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

	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMap(req.TmplId)
	for key, value := range document {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}
		if field.Mode == enum.FieldPersonCom {
			userIds := convertPersonCom(value)
			document[key] = GetUserList(userIds)
		}
		if field.Mode == enum.FieldStatusCom {
			statusId, _ := value.(int32)
			status, _ := tmplCaller.GetStatus(req.TmplId, int(statusId))
			document[key] = status
		}
		if field.Mode == enum.FieldRelatedCom {
			var relatedData []interface{}
			valueAssert := value.(primitive.A)
			for _, issueId := range valueAssert {
				issueIdStr, okr := issueId.(string)
				if okr {
					relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
					isMember := s.GetTmplAuth(req.WsId, req.TmplId, userid)
					if err1 == nil {
						relatedDocument["is_member"] = isMember
						relatedData = append(relatedData, relatedDocument)
						break
					}
				}
			}
			document[key] = relatedData
		}
	}
	return document, nil
}

type Link struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func (s *TmplSrv) CreateSubAction(ctx *gin.Context, req types.CreateSubActionReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	var createReq types.CommonReq
	createReq.WsId = req.WsId
	createReq.TmplId = req.TmplId
	document, _ := s.Create(ctx, createReq, userTmplRightInfo)
	if document != nil {
		if err = global.GVA_DB.Create(&model.TmplSubObjModel{
			WsId:      req.WsId,
			TmplId:    req.ParentTmplId,
			SubTmplId: req.TmplId,
			ObjId:     req.ObjId,
			SubObjId:  document["_id"].(string),
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
	} else {
		return nil, errors.New("操作失败")
	}

	return
}

func (s *TmplSrv) Create(ctx *gin.Context, req types.CommonReq, userTmplRightInfo right.UserTmplRight) (resp bson.M, err error) {
	//判断是否有添加权限
	isPermission := false
	if userTmplRightInfo.TmplMode == enum.TmplModePrivate {
		for _, p := range strings.Split(userTmplRightInfo.TmplEditPermission, ",") {
			if p == enum.RoleEditPermissionCreate {
				isPermission = true
			}
		}
	} else {
		isPermission = true
	}

	if !isPermission {
		return nil, errors.New("无创建任务权限")
	}

	tmplCaller := new(tmpl.Caller)
	screens, _ := tmplCaller.GetScreenByModule(req.TmplId, "create")

	var errMsg []string
	document := bson.M{}
	for _, screen := range screens {
		val := ctx.PostForm(screen.FieldKey)
		switch screen.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom, enum.FieldRelatedCom:
			var enumValues []string
			err = json.Unmarshal([]byte(val), &enumValues)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = enumValues
		case enum.FieldNumberCom, enum.FieldProgressCom:
			if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			if len(val) == 0 {
				continue
			}
			var number int64
			number, err = strconv.ParseInt(val, 10, 32)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			document[screen.FieldKey] = int32(number)
		case enum.FieldPersonCom:
			var enumValues []int32
			err = json.Unmarshal([]byte(val), &enumValues)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			res, err := common.SliceRemoveDuplicate(enumValues)
			if err != nil {
				errMsg = append(errMsg, err.Error())
			}
			document[screen.FieldKey] = res.([]int32)
		case enum.FieldLinkCom:
			var link Link
			err = json.Unmarshal([]byte(val), &link)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			if screen.Required == enum.ScreenRequiredYes && (link.Name == "" || link.URL == "") {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = link
		default:
			if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = val
		}
	}
	if len(errMsg) > 0 {
		return nil, errors.New(strings.Join(errMsg, "\n"))
	}
	//获取流程中的初始状态
	status, err := tmplCaller.GetFirstStatus(req.TmplId)
	if err != nil {
		return nil, err
	}
	userid, _ := ctx.Get("userid")

	document["ws_id"] = req.WsId
	document["tmpl_id"] = req.TmplId
	document["status"] = status.Id
	document["creator"] = []int32{int32(userid.(int))}
	document["created_at"] = common.GetCurrentTime()
	document["updated_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	result, err := collection.InsertOne(context.TODO(), &document)
	if err != nil {
		return nil, err
	}
	objID := result.InsertedID.(primitive.ObjectID)
	go log(userid.(int), req.WsId, req.TmplId, objID.Hex(), nil, document, "create", tmplCaller)
	document["_id"] = objID.Hex()
	return document, nil
}

func (s *TmplSrv) Update(ctx *gin.Context, req types.UpdateReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	userid, _ := ctx.Get("userid")

	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
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

	tmplCaller := new(tmpl.Caller)
	screens, _ := tmplCaller.GetScreenByModule(req.TmplId, "detail")
	fieldKey := ctx.PostForm("field_key")
	val := ctx.PostForm(fieldKey)

	conflictRight := false //权限冲突标记
	document := bson.M{}
	for _, screen := range screens {
		if screen.FieldKey == fieldKey {
			if screen.Level == "level1" {
				return nil, errors.New("参数非法，level1级别字段不支持修改")
			}

			if len(screen.ReadOnlyRule) != 0 {
				//判断编辑字段的读写规则
				var readOnlyRule struct {
					StatusList    []int    `json:"status_list"`
					UserList      []int    `json:"user_list"`
					IssueRoleList []string `json:"issue_role_list"`
				}
				err := json.Unmarshal([]byte(screen.ReadOnlyRule), &readOnlyRule)
				if err != nil {
					return nil, err
				}
				var (
					isSign = false //判断状态是否命中
					isGo   = false //判断是有权限编辑
				)

				for _, s := range readOnlyRule.StatusList {
					if s == int(oldDocument["status"].(int32)) {
						isSign = true
						conflictRight = true
					}
				}
				if isSign {
					for _, u := range readOnlyRule.UserList {
						if u == userid {
							isGo = true
							break
						}
					}
					if !isGo {
						for _, f := range readOnlyRule.IssueRoleList {
							fieldMap := tmplCaller.GetFieldsMap(req.TmplId)
							_, ok := fieldMap[f]
							if !ok {
								continue
							}
							if oldDocument[f] == nil {
								continue
							}
							val := reflect.ValueOf(oldDocument[f])
							var userIds []int
							for i := 0; i < val.Len(); i++ {
								v := val.Index(i)
								vInt32, ok := v.Interface().(int32)
								if ok {
									userIds = append(userIds, int(vInt32))
								}
							}
							for _, u := range userIds {
								if u == userid {
									isGo = true
									break
								}
							}
							if isGo {
								break
							}
						}
					}
					if !isGo {
						return nil, errors.New("当前状态下该字段被设置为只读，不支持编辑")
					}
				}
			}

			switch screen.Mode {
			case enum.FieldMultiSelectCom, enum.FieldSelectCom, enum.FieldRelatedCom:
				var enumValues []string
				err = json.Unmarshal([]byte(val), &enumValues)
				if err != nil {
					return nil, errors.New(screen.Name + "格式错误")
				}
				if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
					return nil, errors.New(screen.Name + "必填项")
				}
				document[screen.FieldKey] = enumValues
			case enum.FieldNumberCom, enum.FieldProgressCom:
				if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
					return nil, errors.New(screen.Name + "必填项")
				}
				if len(val) == 0 {
					document[screen.FieldKey] = ""
				} else {
					var number int64
					number, err = strconv.ParseInt(val, 10, 32)
					if err != nil {
						return nil, errors.New(screen.Name + "格式错误")
					}
					document[screen.FieldKey] = int32(number)
				}
			case enum.FieldPersonCom:
				var enumValues []int32
				err = json.Unmarshal([]byte(val), &enumValues)
				if err != nil {
					return nil, errors.New(screen.Name + "格式错误")
				}
				if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
					return nil, errors.New(screen.Name + "必填项")
				}
				res, err := common.SliceRemoveDuplicate(enumValues)
				if err != nil {
					return nil, errors.New(screen.Name + "必填项")
				}
				document[screen.FieldKey] = res.([]int32)
			case enum.FieldLinkCom:
				var link Link
				err = json.Unmarshal([]byte(val), &link)
				if err != nil {
					return nil, errors.New(screen.Name + "格式错误")
				}
				if screen.Required == enum.ScreenRequiredYes && (link.Name == "" || link.URL == "") {
					return nil, errors.New(screen.Name + "必填项")
				}
				document[screen.FieldKey] = link
			default:
				if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
					return nil, errors.New(screen.Name + "必填项")
				}
				document[screen.FieldKey] = val
			}
		}
	}
	//如果没有配置读写规则(优先级最高)，则判断任务的基础权限（修改权限：创建人、处理人）
	if !conflictRight && getDataEditRight(userid.(int), oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	document["updated_at"] = common.GetCurrentTime()

	update := bson.M{
		"$set": document,
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	go log(userid.(int), req.WsId, req.TmplId, req.Id, document, oldDocument, "update", tmplCaller)
	return
}

func (s *TmplSrv) Duplicate(userid int) {
}

func (s *TmplSrv) DeleteSubAction(userid int, req types.DeleteSubActionReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	// 删除数据
	var delReq types.DeleteReq
	delReq.WsId = req.WsId
	delReq.TmplId = req.TmplId
	delReq.Ids = req.Ids
	_, err = s.Delete(userid, delReq, userTmplRightInfo)
	if err != nil {
		return nil, err
	}

	// 删除关联关系
	ids := strings.Split(req.Ids, ",")
	if err = global.GVA_DB.Where(
		"ws_id=? and tmpl_id=? and sub_tmpl_id=? and obj_id =? and sub_obj_id in ?",
		req.WsId, req.ParentTmplId, req.TmplId, req.ObjId, ids).Delete(
		&model.TmplSubObjModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *TmplSrv) Delete(userid int, req types.DeleteReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	ids := strings.Split(req.Ids, ",")
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		objectIds = append(objectIds, objectId)
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": objectIds,
		},
	}
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}
	var delObjectIds []primitive.ObjectID
	var delDocuments []bson.M
	for _, document := range documents {
		if getDataEditRight(userid, document, userTmplRightInfo).Delete == "yes" {
			delObjectIds = append(delObjectIds, document["_id"].(primitive.ObjectID))
			delDocuments = append(delDocuments, document)
		}
	}
	filter = bson.M{
		"_id": bson.M{
			"$in": delObjectIds,
		},
	}
	_, err = collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// 执行规则
	go rule_action_log.DelRuleAction(req.WsId, req.TmplId, delDocuments)
	return
}

func (s *TmplSrv) GetStepList(req types.GetStepListReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
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

	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return nil, errors.New("数据不存在")
	}
	statusId, ok := document["status"].(int32)
	if !ok {
		return nil, errors.New("状态数据异常")
	}

	tmplCaller := new(tmpl.Caller)
	stepList, err := tmplCaller.GetStepList(req.TmplId, int(statusId))
	if err != nil {
		return nil, err
	}
	type stepVo struct {
		Id               int    `json:"id"`
		StartStatusId    int    `json:"start_status_id"`
		StartStatusName  string `json:"start_status_name"`
		StartStatusColor string `json:"start_status_color"`
		EndStatusId      int    `json:"end_status_id"`
		EndStatusName    string `json:"end_status_name"`
		EndStatusColor   string `json:"end_status_color"`
	}
	var stepVoList []stepVo
	for _, step := range stepList {
		var (
			startStatusName  string
			startStatusColor string
			endStatusName    string
			endStatusColor   string
		)
		startStatus, err := tmplCaller.GetStatus(req.TmplId, step.StartStatusId)
		if err == nil {
			startStatusName = startStatus.Name
			startStatusColor = startStatus.Color
		}
		endStatus, err := tmplCaller.GetStatus(req.TmplId, step.EndStatusId)
		if err == nil {
			endStatusName = endStatus.Name
			endStatusColor = endStatus.Color
		}
		var stepVo = stepVo{
			Id:               step.Id,
			StartStatusId:    step.StartStatusId,
			StartStatusName:  startStatusName,
			StartStatusColor: startStatusColor,
			EndStatusId:      step.EndStatusId,
			EndStatusName:    endStatusName,
			EndStatusColor:   endStatusColor,
		}
		stepVoList = append(stepVoList, stepVo)
	}
	return stepVoList, nil
}

func (s *TmplSrv) GetStepScreen(ctx *gin.Context, req types.GetStepScreenReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"_id": objectID},
		},
	}

	var oldDocument = bson.M{}
	err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	userid, _ := ctx.Get("userid")
	tmplCaller := new(tmpl.Caller)
	//判断是否限制操作
	limits, err := tmplCaller.GetStepLimiter(req.TmplId, req.StepId)
	if err != nil {
		return nil, err
	}
	conflictRight := false //权限冲突标记
	for _, l := range limits {
		if l.Mode == "permission" {
			conflictRight = true
			var rule struct {
				UserList      []int    `json:"user_list"`
				IssueRoleList []string `json:"issue_role_list"`
			}
			err := json.Unmarshal([]byte(l.Rule), &rule)
			if err != nil {
				return nil, err
			}
			var isGo = false //判断是有权限编辑
			for _, u := range rule.UserList {
				if u == userid {
					isGo = true
					break
				}
			}
			if !isGo {
				for _, f := range rule.IssueRoleList {
					fieldMap := tmplCaller.GetFieldsMap(req.TmplId)
					_, ok := fieldMap[f]
					if !ok {
						continue
					}
					if oldDocument[f] == nil {
						continue
					}
					userIds := convertPersonCom(oldDocument[f])
					for _, u := range userIds {
						if u == userid {
							isGo = true
							break
						}
					}
					if isGo {
						break
					}
				}
			}
			if !isGo {
				return nil, errors.New("你没有流转当前状态的权限，请查看工作流中的流转规则")
			}
		} else if l.Mode == "sub" {
			//todo
			continue
		} else {
			continue
		}
	}

	//如果没有配置权限类型的限制器(优先级最高)，则判断任务的基础权限（转化状态权限：创建人、处理人）
	if !conflictRight && getDataEditRight(userid.(int), oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	_, err = tmplCaller.GetStep(req.TmplId, req.StepId)
	if err != nil {
		return nil, err
	}

	resp, err = tmplCaller.GetStepScreen(req.TmplId, req.StepId)
	return
}

func (s *TmplSrv) SwitchStep(ctx *gin.Context, req types.SwitchStepReq, userTmplRightInfo right.UserTmplRight) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"_id": objectID},
		},
	}

	var oldDocument = bson.M{}
	err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	userid, _ := ctx.Get("userid")

	tmplCaller := new(tmpl.Caller)
	step, err := tmplCaller.GetStep(req.TmplId, req.StepId)
	if err != nil {
		return nil, err
	}

	//判断是否限制操作
	limits, err := tmplCaller.GetStepLimiter(req.TmplId, req.StepId)
	if err != nil {
		return nil, err
	}
	conflictRight := false //权限冲突标记
	for _, l := range limits {
		if l.Mode == "permission" {
			conflictRight = true
			var rule struct {
				UserList      []int    `json:"user_list"`
				IssueRoleList []string `json:"issue_role_list"`
			}
			err := json.Unmarshal([]byte(l.Rule), &rule)
			if err != nil {
				return nil, err
			}
			var isGo = false //判断是有权限编辑
			for _, u := range rule.UserList {
				if u == userid {
					isGo = true
					break
				}
			}
			if !isGo {
				for _, f := range rule.IssueRoleList {
					fieldMap := tmplCaller.GetFieldsMap(req.TmplId)
					_, ok := fieldMap[f]
					if !ok {
						continue
					}
					if oldDocument[f] == nil {
						continue
					}
					userIds := convertPersonCom(oldDocument[f])
					for _, u := range userIds {
						if u == userid {
							isGo = true
							break
						}
					}
					if isGo {
						break
					}
				}
			}
			if !isGo {
				return nil, errors.New("你没有流转当前状态的权限，请查看工作流中的流转规则")
			}
		} else if l.Mode == "sub" {
			//todo
			continue
		} else {
			continue
		}
	}

	//如果没有配置权限类型的限制器(优先级最高)，则判断任务的基础权限（转化状态权限：创建人、处理人）
	if !conflictRight && getDataEditRight(userid.(int), oldDocument, userTmplRightInfo).Edit == "no" {
		return nil, errors.New("无操作数据权限")
	}

	screens, err := tmplCaller.GetStepScreen(req.TmplId, req.StepId)
	if err != nil {
		return nil, err
	}

	var document = bson.M{}
	var errMsg []string
	for _, screen := range screens {
		if screen.Level == "level1" {
			errMsg = append(errMsg, screen.Name+"，参数非法，level1级别字段不支持修改")
			continue
		}

		val := ctx.PostForm(screen.FieldKey)
		switch screen.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom, enum.FieldRelatedCom:
			var enumValues []string
			err = json.Unmarshal([]byte(val), &enumValues)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = enumValues
		case enum.FieldNumberCom, enum.FieldProgressCom:
			if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			if len(val) == 0 {
				document[screen.FieldKey] = ""
			} else {
				var number int64
				number, err = strconv.ParseInt(val, 10, 32)
				if err != nil {
					errMsg = append(errMsg, screen.Name+"参数格式错误")
					continue
				}
				document[screen.FieldKey] = int32(number)
			}
		case enum.FieldPersonCom:
			var enumValues []int32
			err = json.Unmarshal([]byte(val), &enumValues)
			if err != nil {
				errMsg = append(errMsg, screen.Name+"参数格式错误")
				continue
			}
			if screen.Required == enum.ScreenRequiredYes && len(enumValues) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			res, err := common.SliceRemoveDuplicate(enumValues)
			if err != nil {
				return nil, errors.New(screen.Name + "必填项")
			}
			document[screen.FieldKey] = res.([]int32)
		case enum.FieldLinkCom:
			var link Link
			err = json.Unmarshal([]byte(val), &link)
			if err != nil {
				return nil, errors.New(screen.Name + "格式错误")
			}
			if screen.Required == enum.ScreenRequiredYes && (link.Name == "" || link.URL == "") {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = link
		default:
			if screen.Required == enum.ScreenRequiredYes && len(val) == 0 {
				errMsg = append(errMsg, screen.Name+"必填项")
				continue
			}
			document[screen.FieldKey] = val
		}
	}
	if len(errMsg) > 0 {
		return nil, errors.New(strings.Join(errMsg, "\n"))
	}

	document["status"] = step.EndStatusId
	document["updated_at"] = common.GetCurrentTime()
	update := bson.M{
		"$set": document,
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	go log(userid.(int), req.WsId, req.TmplId, req.Id, document, oldDocument, "switch", tmplCaller)

	return
}

func (s *TmplSrv) GetProgressList(userid int, req types.GetProgressList) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_progress")

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
		},
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	skip := (pageNum - 1) * pageSize

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).
		SetLimit(int64(pageSize)).
		SetSkip(int64(skip))

	cnt, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	for _, document := range documents {
		userId := int(document["creator"].(int32))
		var userIds = []int{userId}
		users := GetUserList(userIds)
		if len(users) >= 1 {
			document["creator"] = users[0]
		} else {
			document["creator"] = nil
		}

		document["permission"] = "no"
		if userId == userid {
			document["permission"] = "yes"
		}
	}

	resp = commonTypes.BasePageResp{
		List: documents,
		Cnt:  cnt,
	}
	return
}

func (s *TmplSrv) getProgressCount(wsId int, tmplId int, issueId string) int64 {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_progress")
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": wsId},
			{"tmpl_id": tmplId},
			{"issue_id": issueId},
		},
	}
	cnt, _ := collection.CountDocuments(context.TODO(), filter)
	return cnt
}

func (s *TmplSrv) AddProgress(userid int, req types.AddProgress) (resp interface{}, err error) {
	document := bson.M{}
	document["issue_id"] = req.IssueId
	document["ws_id"] = req.WsId
	document["tmpl_id"] = req.TmplId
	document["content"] = req.Content
	document["creator"] = userid
	document["created_at"] = common.GetCurrentTime()
	document["updated_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_progress")
	_, err = collection.InsertOne(context.TODO(), &document)
	if err != nil {
		return
	}
	go AtPersonNotice(userid, req.WsId, req.TmplId, document)
	return
}

func (s *TmplSrv) UpdateProgress(userid int, req types.UpdateProgress) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_progress")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
			{"_id": objectID},
		},
	}
	var progress bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	creator, ok := progress["creator"].(int32)
	if !ok {
		return nil, err
	}
	if int(creator) != userid {
		return nil, errors.New("无权限修改")
	}

	progress["content"] = req.Content
	update := bson.M{
		"$set": progress,
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	go AtPersonNotice(userid, req.WsId, req.TmplId, progress)
	return
}

func (s *TmplSrv) DeleteProgress(userid int, req types.DeleteProgress) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_progress")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
			{"_id": objectID},
		},
	}
	var progress bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		return nil, errors.New("数据不存在")
	}
	creator, ok := progress["creator"].(int32)
	if !ok {
		return nil, err
	}
	if int(creator) != userid {
		return nil, errors.New("无权限修改")
	}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return
}

func (s *TmplSrv) GetLogList(req types.GetLogList) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_log")

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
		},
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	skip := (pageNum - 1) * pageSize

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).
		SetLimit(int64(pageSize)).
		SetSkip(int64(skip))

	cnt, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	for _, document := range documents {
		userId := int(document["creator"].(int32))
		var userIds = []int{userId}
		users := GetUserList(userIds)
		if len(users) >= 1 {
			document["creator"] = users[0]
		} else {
			document["creator"] = ""
		}
		logAction := document["action"].(string)
		document["action_desc"] = LogActionMap[logAction]
	}

	resp = commonTypes.BasePageResp{
		List: documents,
		Cnt:  cnt,
	}
	return
}

func (s *TmplSrv) GetSubListCount(req types.GetSubListCountReq) (resp interface{}, err error) {
	type TmplSubVo struct {
		SubObjId string `json:"sub_obj_id"`
	}
	var tmplSubList []TmplSubVo
	sql := "select b.sub_obj_id from tmpl_sub_config AS a " +
		"LEFT JOIN tmpl_sub_obj AS b " +
		"ON a.ws_id = b.ws_id and a.tmpl_id = b.tmpl_id and a.sub_tmpl_id = b.sub_tmpl_id " +
		"where b.id is not null and a.ws_id=? and a.tmpl_id=? and b.obj_id=?"
	global.GVA_DB.Raw(sql, req.WsId, req.TmplId, req.ObjId).Scan(&tmplSubList)

	if len(tmplSubList) == 0 {
		return 0, nil
	} else {
		var objectIds []primitive.ObjectID
		for _, obj := range tmplSubList {
			objectId, err := primitive.ObjectIDFromHex(obj.SubObjId)
			if err != nil {
				continue
			}
			objectIds = append(objectIds, objectId)
		}
		filter := bson.M{
			"_id": bson.M{
				"$in": objectIds,
			},
		}
		collection := global.GVA_MONGO.Database("mark3").Collection("issue")
		cnt, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			return 0, err
		} else {
			return cnt, nil
		}
	}
}
