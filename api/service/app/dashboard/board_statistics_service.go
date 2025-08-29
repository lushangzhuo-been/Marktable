package dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"log"
	dashboardConst "mark3/const/dashboard"
	"mark3/global"
	model "mark3/repository/db/model/tmpl"
	userModel "mark3/repository/db/model/user"
	"math"
	"sort"
	"strconv"
	"time"
)

type BoardStatisticsService struct{}

func (s *BoardStatisticsService) GetBoardStatisticsByID(boardID int, filter, filterDown, lor string, queryDB *gorm.DB) (map[string]interface{}, error) {
	var board model.BoardModel
	if err := queryDB.Where("id = ?", boardID).First(&board).Error; err != nil {
		return nil, err
	}

	boardDict := s.QueryResultToMap(board)
	boardDict["filter"] = filter
	boardDict["lor"] = lor
	boardDict["filter_down"] = filterDown

	return s.GetBoardStatisticsPreview(boardDict, queryDB)
}

// GetMgFilterNew 组装MongoDB查询条件
func (s *BoardStatisticsService) GetMgFilterNew(mgFilter bson.M, filterStr string, lor string, fieldDict map[string]model.FieldModel) bson.M {
	if filterStr == "" {
		return mgFilter
	}

	var dataList []map[string]interface{}
	if err := json.Unmarshal([]byte(filterStr), &dataList); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return mgFilter
	}

	var filterList []bson.M
	for _, one := range dataList {
		data, ok := one["data"].([]interface{})
		if !ok {
			continue
		}

		var dataItems []map[string]interface{}
		for _, item := range data {
			if m, ok := item.(map[string]interface{}); ok {
				dataItems = append(dataItems, m)
			}
		}

		filterListOne := s.GetOnlyFilter(dataItems, fieldDict)
		if len(filterListOne) == 0 {
			continue
		}

		filterType, _ := one["filter_type"].(string)
		switch filterType {
		case "filter":
			filterList = append(filterList, filterListOne[0])
		case "group":
			if len(filterListOne) == 1 {
				filterList = append(filterList, filterListOne[0])
			} else if len(filterListOne) > 1 {
				groupLor, _ := one["lor"].(string)
				if groupLor == "filter_or" {
					filterList = append(filterList, bson.M{"$or": filterListOne})
				} else {
					filterList = append(filterList, bson.M{"$and": filterListOne})
				}
			}
		}
	}

	if len(filterList) > 0 {
		if lor == "filter_or" {
			mgFilter["$or"] = filterList
		} else {
			mgFilter["$and"] = filterList
		}
	}

	return mgFilter
}

// GetOnlyFilter 处理单组过滤条件
func (s *BoardStatisticsService) GetOnlyFilter(dataList []map[string]interface{}, fieldDict map[string]model.FieldModel) []bson.M {
	var filterList []bson.M

	for _, one := range dataList {
		fieldKey, _ := one["field_key"].(string)
		fieldInfo, exists := fieldDict[fieldKey]
		if !exists {
			continue
		}

		op, _ := one["op"].(string)
		value := one["value"]

		var cond bson.M
		switch fieldInfo.Mode {
		case "text_com", "textarea_com":
			if strValue, ok := value.(string); ok {
				cond = bson.M{fieldKey: bson.M{"$regex": strValue, "$options": "i"}}
			}
		case "link_com":
			key := fmt.Sprintf("%s.%sname", fieldKey, fieldKey)
			cond = bson.M{key: bson.M{op: value}}
		default:
			cond = bson.M{fieldKey: bson.M{op: value}}
		}

		if cond != nil {
			filterList = append(filterList, cond)
		}
	}

	return filterList
}

// GetMgFilterDown 组装MongoDB下钻过滤条件
func (s *BoardStatisticsService) GetMgFilterDown(mgFilter bson.M, filterStr string, boardDict map[string]interface{}) bson.M {
	// 初始化值
	boardDict["x_value"] = nil
	boardDict["group_value"] = nil

	if filterStr == "" {
		return mgFilter
	}

	var filterDict map[string]interface{}
	if err := json.Unmarshal([]byte(filterStr), &filterDict); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return mgFilter
	}

	var filterList []bson.M

	// 处理x_axis条件
	if xAxis, ok := filterDict["x_axis"].(string); ok {
		xValueStr := fmt.Sprintf("%v", filterDict["x_value"])

		xValueStrNum, err := strconv.Atoi(xValueStr)
		var xValue interface{}
		if err != nil {
			fmt.Println("转换失败:", err)
			xValue = xValueStr
		} else {
			fmt.Println("转换成功:", xValueStrNum)
			xValue = xValueStrNum
		}

		boardDict["x_value"] = xValueStr

		if xValue == "空值" {
			cond := bson.M{"$or": []bson.M{
				{xAxis: bson.M{"$eq": nil}},
				{xAxis: bson.M{"$eq": ""}},
				{xAxis: bson.M{"$eq": bson.A{}}},
			}}
			filterList = append(filterList, cond)
		} else if dateMode, ok := filterDict["axis_date_mode"].(string); ok &&
			contains([]string{"day", "month", "year"}, dateMode) {
			s.AddDateFilter(&filterList, xAxis, xValue.(string), dateMode)
		} else {
			cond := bson.M{xAxis: bson.M{"$eq": xValue}}
			filterList = append(filterList, cond)
		}
	}

	// 处理group_axis条件
	if groupAxis, ok := filterDict["group_axis"].(string); ok {
		groupValueStr := fmt.Sprintf("%v", filterDict["group_value"])
		groupValueNum, err := strconv.Atoi(groupValueStr)
		var groupValue interface{}
		if err != nil {
			fmt.Println("转换失败:", err)
			groupValue = groupValueStr
		} else {
			fmt.Println("转换成功:", groupValueNum)
			groupValue = groupValueNum
		}

		boardDict["group_value"] = groupValue

		if groupValue == "空值" {
			cond := bson.M{"$or": []bson.M{
				{groupAxis: bson.M{"$eq": nil}},
				{groupAxis: bson.M{"$eq": ""}},
				{groupAxis: bson.M{"$eq": bson.A{}}},
			}}
			filterList = append(filterList, cond)
		} else if dateMode, ok := filterDict["group_date_mode"].(string); ok &&
			contains([]string{"day", "month", "year"}, dateMode) {
			s.AddDateFilter(&filterList, groupAxis, groupValue.(string), dateMode)
		} else {
			cond := bson.M{groupAxis: bson.M{"$eq": groupValue}}
			filterList = append(filterList, cond)
		}
	}

	// 合并过滤条件
	if len(filterList) > 0 {
		if existingAnd, ok := mgFilter["$and"].([]bson.M); ok {
			mgFilter["$and"] = append(existingAnd, filterList...)
		} else {
			mgFilter["$and"] = filterList
		}
	}

	return mgFilter
}

// AddDateFilter 添加日期过滤条件
func (s *BoardStatisticsService) AddDateFilter(filterList *[]bson.M, key, dateValue, dateMode string) {
	var startTime, endTime time.Time
	var err error

	switch dateMode {
	case "day":
		startTime, err = time.Parse("2006-01-02", dateValue)
		if err != nil {
			return
		}
		endTime = startTime.AddDate(0, 0, 1)
	case "month":
		startTime, err = time.Parse("2006-01", dateValue)
		if err != nil {
			return
		}
		endTime = startTime.AddDate(0, 1, 0)
	case "year":
		startTime, err = time.Parse("2006", dateValue)
		if err != nil {
			return
		}
		endTime = startTime.AddDate(1, 0, 0)
	default:
		return
	}

	*filterList = append(*filterList, bson.M{
		key: bson.M{"$gte": startTime.Format("2006-01-02")},
	}, bson.M{
		key: bson.M{"$lt": endTime.Format("2006-01-02")},
	})
}

// contains 检查字符串是否在切片中
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// TwoDimension 二维分组
func (s *BoardStatisticsService) TwoDimension(
	boardDict map[string]interface{},
	fieldDict map[string]model.FieldModel,
	dataList []map[string]interface{},
	db *gorm.DB,
) []map[string]interface{} {

	// 获取X轴字段和日期模式
	fieldName, _ := boardDict["x_axis"].(string)
	dateMode, _ := boardDict["axis_date_mode"].(string)
	xValue, _ := boardDict["x_value"].(string)

	// 第一维度分组
	xAxisDimension := s.OneDimension(
		boardDict,
		fieldDict,
		dataList,
		fieldName,
		xValue,
		dateMode,
		db,
	)

	if len(xAxisDimension) == 0 {
		return xAxisDimension
	}

	// 获取分组字段和日期模式
	groupFieldName, _ := boardDict["group_axis"].(string)
	groupDateMode, _ := boardDict["group_date_mode"].(string)
	groupValue, _ := boardDict["group_value"].(string)

	// 第二维度分组
	for i := range xAxisDimension {
		groupDataList, ok := xAxisDimension[i]["data_list"].([]map[string]interface{})
		if !ok {
			continue
		}

		groupDimension := s.OneDimension(
			boardDict,
			fieldDict,
			groupDataList,
			groupFieldName,
			groupValue,
			groupDateMode,
			db,
		)

		xAxisDimension[i]["group_dimension"] = groupDimension
	}

	return xAxisDimension
}

func (s *BoardStatisticsService) GetBoardStatisticsPreview(boardDict map[string]interface{}, queryDB *gorm.DB) (map[string]interface{}, error) {
	wsID := int(boardDict["ws_id"].(float64))
	tmplID := int(boardDict["tmpl_id"].(float64))

	fieldDict, err := s.GetTmplFieldDict(wsID, tmplID)
	if err != nil {
		return nil, err
	}

	mgField := bson.M{"_id": 0}
	if xAxis, ok := boardDict["x_axis"].(string); ok && xAxis != "" {
		mgField[xAxis] = 1
	}
	if groupAxis, ok := boardDict["group_axis"].(string); ok && groupAxis != "" {
		mgField[groupAxis] = 1
	}

	mgFilter := bson.M{
		"ws_id":   boardDict["ws_id"],
		"tmpl_id": boardDict["tmpl_id"],
	}
	filter, ok := boardDict["filter"].(string)
	if !ok {
		filter = ""
	}
	lor, ok := boardDict["lor"].(string)
	if !ok {
		lor = ""
	}
	mgFilter = s.GetMgFilterNew(mgFilter, filter, lor, fieldDict)
	filterDown, ok := boardDict["filter_down"]
	if ok {
		mgFilter = s.GetMgFilterDown(mgFilter, filterDown.(string), boardDict)
	}
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	cursor, err := collection.Find(context.Background(), mgFilter, options.Find().SetProjection(mgField))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	dataList := make([]map[string]interface{}, 0)
	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		dataList = append(dataList, s.BsonToMap(document))
	}

	res := make([]map[string]interface{}, 0)
	mode := boardDict["mode"].(string)

	switch mode {
	case string(dashboardConst.BaseHistogram), string(dashboardConst.BaseLine):
		fieldName, ok := boardDict["x_axis"].(string)
		if !ok {
			fieldName = ""
		}
		fieldValue, ok := boardDict["x_value"].(string)
		if !ok {
			fieldValue = ""
		}
		dateMode, ok := boardDict["axis_date_mode"].(string)
		if !ok {
			dateMode = ""
		}
		res = s.OneDimension(boardDict, fieldDict, dataList, fieldName, fieldValue, dateMode, queryDB)
	case string(dashboardConst.BasePie), string(dashboardConst.DonutPie):
		fieldName, ok := boardDict["group_axis"].(string)
		if !ok {
			fieldName = ""
		}
		fieldValue, ok := boardDict["group_value"].(string)
		if !ok {
			fieldValue = ""
		}
		dateMode, ok := boardDict["group_date_mode"].(string)
		if !ok {
			dateMode = ""
		}
		res = s.OneDimension(boardDict, fieldDict, dataList, fieldName, fieldValue, dateMode, queryDB)
		s.CalPropData(res, "value", "rate")
	case string(dashboardConst.GroupHistogram), string(dashboardConst.StackHistogram), string(dashboardConst.GroupLine):
		res = s.TwoDimension(boardDict, fieldDict, dataList, queryDB)
	default:
		res = make([]map[string]interface{}, 0)
	}

	xCN := ""
	if xAxis, ok := boardDict["x_axis"].(string); ok {
		if field, ok := fieldDict[xAxis]; ok {
			xCN = field.Name
		}
	}

	groupCN := ""
	if groupAxis, ok := boardDict["group_axis"].(string); ok {
		if field, ok := fieldDict[groupAxis]; ok {
			groupCN = field.Name
		}
	}

	return map[string]interface{}{
		"board_statistics": res,
		"x_cn":             xCN,
		"group_cn":         groupCN,
	}, nil
}

func (s *BoardStatisticsService) GetTmplFieldDict(wsID, tmplID int) (map[string]model.FieldModel, error) {
	fieldList, err := s.GetTmplFieldDictAll(wsID, tmplID)
	if err != nil {
		return nil, err
	}

	fieldDict := make(map[string]model.FieldModel)
	for _, field := range fieldList {
		fieldDict[field.FieldKey] = field
	}

	return fieldDict, nil
}

func (s *BoardStatisticsService) GetUserNameByIDs(ids []int, queryDB *gorm.DB) (map[string]map[string]interface{}, error) {
	var users []userModel.UserModel
	if err := queryDB.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}

	userDict := make(map[string]map[string]interface{})
	for _, user := range users {
		userDict[strconv.Itoa(user.Id)] = map[string]interface{}{
			"id":        user.Id,
			"username":  user.Username,
			"full_name": user.FullName,
		}
	}

	return userDict, nil
}

func (s *BoardStatisticsService) GetRelatedDocuments(wsID, tmplID int, issueIDs []string) (map[string]map[string]interface{}, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range issueIDs {
		if id == "空值" {
			continue
		}
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objID)
	}

	mgFilter := bson.M{
		"ws_id":   wsID,
		"tmpl_id": tmplID,
		"_id":     bson.M{"$in": objectIDs},
	}

	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	cursor, err := collection.Find(context.Background(), mgFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	dataDict := make(map[string]map[string]interface{})
	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		id := document["_id"].(primitive.ObjectID).Hex()
		dataDict[id] = s.BsonToMap(document)
	}

	return dataDict, nil
}

// TimeComData 时间分组数据处理
func (s *BoardStatisticsService) TimeComData(
	dataList []map[string]interface{},
	fieldInfo model.FieldModel,
	fieldValue string,
	boardDict map[string]interface{},
	dateMode string,
) []map[string]interface{} {

	resDict := make(map[string]int)
	dataDict := make(map[string][]map[string]interface{})
	fieldName := fieldInfo.FieldKey
	showEmpty := boardDict["show_empty"].(string) == "yes"

	for _, one := range dataList {
		dateValue, _ := one[fieldName].(string)
		value := s.GetDateValue(dateValue, dateMode)

		if value == "" && showEmpty {
			value = "空值"
		}

		if value != "" {
			if fieldValue != "" && fieldValue != value {
				continue
			}

			if _, exists := resDict[value]; !exists {
				resDict[value] = 0
				dataDict[value] = []map[string]interface{}{}
			}
			resDict[value]++
			dataDict[value] = append(dataDict[value], one)
		}
	}

	dataGroupList := make([]map[string]interface{}, 0)
	var emptyDict map[string]interface{}

	for k, v := range resDict {
		item := map[string]interface{}{
			"name":       k,
			"value":      v,
			"filed_mode": fieldInfo.Mode,
			"data_list":  dataDict[k],
		}

		if k == "空值" {
			emptyDict = item
		} else {
			dataGroupList = append(dataGroupList, item)
		}
	}

	// 排序处理
	if len(dataGroupList) > 0 {
		order, _ := boardDict["order"].(string)
		reverse := order != "asc"

		sort.Slice(dataGroupList, func(i, j int) bool {
			nameI := dataGroupList[i]["name"].(string)
			nameJ := dataGroupList[j]["name"].(string)
			if reverse {
				return nameI > nameJ
			}
			return nameI < nameJ
		})
	}

	if emptyDict != nil {
		dataGroupList = append(dataGroupList, emptyDict)
	}

	return dataGroupList
}

// StatusComData 获取状态分组数据
func (s *BoardStatisticsService) StatusComData(
	dataList []map[string]interface{},
	fieldInfo model.FieldModel,
	fieldValue string,
	boardDict map[string]interface{},
	db *gorm.DB,
) []map[string]interface{} {

	fieldName := fieldInfo.FieldKey
	tmplID := int(boardDict["tmpl_id"].(float64))
	statusDict, err := s.GetTmplStatus(tmplID, db)
	if err != nil {
		return nil
	}

	resDict := make(map[string]int)
	dataDict := make(map[string][]map[string]interface{})
	showEmpty := boardDict["show_empty"].(string) == "yes"

	for _, one := range dataList {
		valueInt, _ := one[fieldName]
		value := fmt.Sprintf("%d", valueInt)
		if (value == "" || value == "null") && showEmpty {
			value = "空值"
		}

		if value != "" && value != "null" {
			if fieldValue != "" && fieldValue != value {
				continue
			}

			if _, exists := resDict[value]; !exists {
				resDict[value] = 0
				dataDict[value] = []map[string]interface{}{}
			}
			resDict[value]++
			dataDict[value] = append(dataDict[value], one)
		}
	}

	dataGroupList := make([]map[string]interface{}, 0)
	var emptyDict map[string]interface{}

	for k, v := range resDict {
		item := map[string]interface{}{
			"name":       k,
			"value":      v,
			"filed_mode": fieldInfo.Mode,
			"data_list":  dataDict[k],
		}

		if k == "空值" {
			emptyDict = item
		} else {
			dataGroupList = append(dataGroupList, item)
		}
	}

	// 更新状态名称
	for i := range dataGroupList {
		dataGroupList[i]["status_id"] = dataGroupList[i]["name"]
		if statusID, err := strconv.Atoi(dataGroupList[i]["status_id"].(string)); err == nil {
			if status, exists := statusDict[statusID]; exists {
				dataGroupList[i]["name"] = status["name"]
			}
		}
	}

	// 排序处理
	if len(dataGroupList) > 0 {
		order, _ := boardDict["order"].(string)
		reverse := order != "asc"

		sort.Slice(dataGroupList, func(i, j int) bool {
			valI := dataGroupList[i]["value"].(int)
			valJ := dataGroupList[j]["value"].(int)
			if valI == valJ {
				return dataGroupList[i]["name"].(string) < dataGroupList[j]["name"].(string)
			}
			if reverse {
				return valI > valJ
			}
			return valI < valJ
		})
	}

	if emptyDict != nil {
		dataGroupList = append(dataGroupList, emptyDict)
	}

	return dataGroupList
}

// GetTmplStatus 获取模板状态
func (s *BoardStatisticsService) GetTmplStatus(tmplID int, db *gorm.DB) (map[int]map[string]string, error) {
	var statusList []map[string]interface{}
	result := db.Table("tmpl_status").
		Where("tmpl_id = ?", tmplID).
		Find(&statusList)
	if result.Error != nil {
		return nil, result.Error
	}

	statusDict := make(map[int]map[string]string)
	for _, status := range statusList {
		id := int(status["id"].(int64))
		statusDict[id] = map[string]string{
			"id":   strconv.Itoa(id),
			"name": status["name"].(string),
		}
	}

	return statusDict, nil
}

func (s *BoardStatisticsService) OneDimension(boardDict map[string]interface{}, fieldDict map[string]model.FieldModel,
	dataList []map[string]interface{}, fieldName, fieldValue, dateMode string, queryDB *gorm.DB) []map[string]interface{} {

	fieldInfo, ok := fieldDict[fieldName]
	if !ok {
		return []map[string]interface{}{}
	}

	switch fieldInfo.Mode {
	case "text_com":
		return s.TextComData(dataList, fieldInfo, fieldValue, boardDict)
	case "select_com", "multi_select_com", "person_com", "related_com":
		return s.ChoiceData(dataList, fieldInfo, fieldValue, boardDict, queryDB)
	case "time_com", "date_com":
		return s.TimeComData(dataList, fieldInfo, fieldValue, boardDict, dateMode)
	default:
		if fieldName == "status" {
			return s.StatusComData(dataList, fieldInfo, fieldValue, boardDict, queryDB)
		} else if fieldName == "creator" {
			return s.ChoiceData(dataList, fieldInfo, fieldValue, boardDict, queryDB)
		}
	}

	return make([]map[string]interface{}, 0)
}

func (s *BoardStatisticsService) TextComData(dataList []map[string]interface{}, fieldInfo model.FieldModel,
	fieldValue string, boardDict map[string]interface{}) []map[string]interface{} {

	resDict := make(map[string]int)
	dataDict := make(map[string][]map[string]interface{})
	fieldName := fieldInfo.FieldKey
	showEmpty := boardDict["show_empty"].(string) == "yes"

	for _, one := range dataList {
		value, ok := one[fieldName].(string)
		if (!ok || value == "") && showEmpty {
			value = "空值"
		}
		if ok && value != "" {
			if fieldValue != "" && fieldValue != value {
				continue
			}
			resDict[value]++
			dataDict[value] = append(dataDict[value], one)
		}
	}

	return s.prepareDataGroupList(resDict, dataDict, fieldInfo, boardDict)
}

func (s *BoardStatisticsService) ChoiceData(dataList []map[string]interface{}, fieldInfo model.FieldModel,
	fieldValue string, boardDict map[string]interface{}, queryDB *gorm.DB) []map[string]interface{} {

	resDict := make(map[string]int)
	dataDict := make(map[string][]map[string]interface{})
	fieldName := fieldInfo.FieldKey
	showEmpty := boardDict["show_empty"].(string) == "yes"

	// 处理枚举值排序
	enumValuesSortDict := make(map[string]int)
	if fieldInfo.Mode == "select_com" || fieldInfo.Mode == "multi_select_com" {
		var enumValues []map[string]interface{}
		if err := json.Unmarshal([]byte(fieldInfo.EnumValues), &enumValues); err == nil {
			for i, enumValue := range enumValues {
				enumValuesSortDict[enumValue["value"].(string)] = i + 1
			}
		}
	}

	for _, one := range dataList {
		values, ok := one[fieldName].(primitive.A)
		if ok && len(values) > 0 {
			for _, val := range values {
				value := fmt.Sprintf("%v", val)
				if fieldValue != "" && fieldValue != value {
					continue
				}
				resDict[value]++
				dataDict[value] = append(dataDict[value], one)
			}
		} else if showEmpty {
			value := "空值"
			if fieldValue != "" && fieldValue != value {
				continue
			}
			resDict[value]++
			dataDict[value] = append(dataDict[value], one)
		}
	}

	dataGroupList := s.prepareDataGroupList(resDict, dataDict, fieldInfo, boardDict)

	// 处理人员信息
	if fieldInfo.Mode == "person_com" || fieldName == "creator" {
		var nameList []string
		for k := range resDict {
			if k != "空值" {
				nameList = append(nameList, k)
			}
		}

		userDict := make(map[string]map[string]interface{})
		if len(nameList) > 0 {
			var ids []int
			for _, name := range nameList {
				if id, err := strconv.Atoi(name); err == nil {
					ids = append(ids, id)
				}
			}
			if len(ids) > 0 {
				var err error
				userDict, err = s.GetUserNameByIDs(ids, queryDB)
				if err != nil {
					log.Printf("GetUserNameByIDs error: %v", err)
				}
			}
		}

		for i, item := range dataGroupList {
			if item["name"] == "空值" {
				dataGroupList[i]["user_id"] = "空值"
				dataGroupList[i]["user_fullname"] = "空值"
				continue
			}

			userID, ok := item["name"].(string)
			if !ok {
				userID = ""
			}
			dataGroupList[i]["user_id"] = userID
			if user, ok := userDict[userID]; ok {
				dataGroupList[i]["name"] = user["full_name"]
				dataGroupList[i]["user_name"] = user["username"]
			} else {
				dataGroupList[i]["user_name"] = userID
			}
		}
	}

	// 处理关联文档
	if fieldInfo.Mode == "related_com" {
		var nameList []string
		for k := range resDict {
			if k != "空值" {
				nameList = append(nameList, k)
			}
		}

		documentsDict := make(map[string]map[string]interface{})
		if len(nameList) > 0 {
			var err error
			documentsDict, err = s.GetRelatedDocuments(
				fieldInfo.WsId,
				fieldInfo.RelatedTmplId,
				nameList,
			)
			if err != nil {
				log.Printf("GetRelatedDocuments error: %v", err)
			}
		}

		for i, item := range dataGroupList {
			if item["name"] == "空值" {
				dataGroupList[i]["document_id"] = "空值"
				dataGroupList[i]["document_name"] = "空值"
				continue
			}

			docID := item["name"].(string)
			dataGroupList[i]["document_id"] = docID
			if doc, ok := documentsDict[docID]; ok {
				dataGroupList[i]["name"] = doc["title"]
			}
		}
	}

	// 处理排序
	if fieldInfo.Mode == "select_com" || fieldInfo.Mode == "multi_select_com" {
		sort.Slice(dataGroupList, func(i, j int) bool {
			sortI := 0
			if val, ok := dataGroupList[i]["select_sort"].(int); ok {
				sortI = val
			}
			sortJ := 0
			if val, ok := dataGroupList[j]["select_sort"].(int); ok {
				sortJ = val
			}
			return sortI < sortJ
		})
	}

	return dataGroupList
}

func (s *BoardStatisticsService) prepareDataGroupList(resDict map[string]int, dataDict map[string][]map[string]interface{},
	fieldInfo model.FieldModel, boardDict map[string]interface{}) []map[string]interface{} {

	dataGroupList := make([]map[string]interface{}, 0)
	var emptyItem map[string]interface{}

	for k, v := range resDict {
		item := map[string]interface{}{
			"name":       k,
			"value":      v,
			"filed_mode": fieldInfo.Mode,
			"data_list":  dataDict[k],
		}

		if k == "空值" {
			emptyItem = item
		} else {
			if fieldInfo.Mode == "select_com" || fieldInfo.Mode == "multi_select_com" {
				item["select_sort"] = 0 // 这里应该从enumValuesSortDict获取
			}
			dataGroupList = append(dataGroupList, item)
		}
	}

	// 排序
	if len(dataGroupList) > 0 {
		order := boardDict["order"].(string)
		reverse := order != "asc"
		sort.Slice(dataGroupList, func(i, j int) bool {
			if dataGroupList[i]["value"].(int) == dataGroupList[j]["value"].(int) {
				return dataGroupList[i]["name"].(string) < dataGroupList[j]["name"].(string)
			}
			if reverse {
				return dataGroupList[i]["value"].(int) > dataGroupList[j]["value"].(int)
			}
			return dataGroupList[i]["value"].(int) < dataGroupList[j]["value"].(int)
		})
	}

	if emptyItem != nil {
		dataGroupList = append(dataGroupList, emptyItem)
	}

	return dataGroupList
}

func (s *BoardStatisticsService) GetDateValue(dateValue string, dateMode string) string {
	if dateValue == "" {
		return ""
	}

	switch dateMode {
	case "day":
		if len(dateValue) >= 10 {
			return dateValue[:10]
		}
	case "month":
		if len(dateValue) >= 7 {
			return dateValue[:7]
		}
	case "year":
		if len(dateValue) >= 4 {
			return dateValue[:4]
		}
	}

	return ""
}

func (s *BoardStatisticsService) CalPropData(dataList []map[string]interface{}, keyName, rateName string) {
	if keyName == "" {
		keyName = "value"
	}
	if rateName == "" {
		rateName = "rate"
	}

	total := 0
	for _, data := range dataList {
		if val, ok := data[keyName].(int); ok {
			total += val
		}
	}

	if total > 0 {
		for _, data := range dataList {
			if val, ok := data[keyName].(int); ok {
				rate := float64(val) / float64(total) * 100
				data[rateName] = math.Round(rate*100) / 100
			}
		}
	}
}

func (s *BoardStatisticsService) BsonToMap(doc bson.M) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range doc {
		if oid, ok := v.(primitive.ObjectID); ok {
			result[k] = oid.Hex()
		} else {
			result[k] = v
		}
	}
	return result
}

func (s *BoardStatisticsService) QueryResultToMap(model interface{}) map[string]interface{} {
	data, _ := json.Marshal(model)
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}
