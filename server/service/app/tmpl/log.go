package tmpl

import (
	"context"
	"fmt"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/service/app"
	"mark3/service/rule_action_log"
	"mark3/service/tmpl"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LogCreateAction     string = "create"
	LogUpdateAction     string = "update"
	LogSwitchAction     string = "switch"
	LogUploadFileAction string = "upload_file"
	LogDeleteFileAction string = "delete_file"
)

var LogActionMap = map[string]string{
	LogCreateAction:     "创建了任务",
	LogUpdateAction:     "修改了任务",
	LogSwitchAction:     "转化了任务",
	LogUploadFileAction: "上传了附件",
	LogDeleteFileAction: "删除了附件",
}

func log(userid int, wsId int, tmplId int, issueId string, newData bson.M, oldData bson.M, action string, tmplCaller *tmpl.Caller) {
	fieldsMap := tmplCaller.GetFieldsMap(tmplId)

	var newDataFormat = make(map[string]string)
	var oldDataFormat = make(map[string]string)
	if action == "create" {
		newDataFormat = formatCreateData(tmplId, oldData, fieldsMap, tmplCaller)
	} else {
		newDataFormat = formatNewData(tmplId, newData, fieldsMap, tmplCaller)
		oldDataFormat = formatOldData(tmplId, oldData, fieldsMap, tmplCaller)
	}
	var changeData = make(map[string]struct {
		OldItem string
		NewItem string
	})

	for key, newItem := range newDataFormat {
		if key == "updated_at" || key == "created_at" {
			continue
		}

		oldItem := oldDataFormat[key]
		if oldItem == "" {
			oldItem = "--"
		}
		if newItem == "" {
			newItem = "--"
		}

		if newItem == oldItem {
			continue
		}
		changeData[key] = struct {
			OldItem string
			NewItem string
		}{OldItem: oldItem, NewItem: newItem}
	}

	if len(changeData) == 0 {
		return
	}

	var logContent []string
	for key, item := range changeData {
		logContent = append(logContent, fmt.Sprintf("设置【%s】：旧值：%s，新值：%s", fieldsMap[key].Name, item.OldItem, item.NewItem))

		//Todo 监控handler字段是否变化，如果变化则：1，向指派人发送通知 2，计算亲密用户
		if key == "handler" {
			var handlerUseridListInt32 []int32
			if action == "create" {
				handlerUseridListInt32 = oldData[key].([]int32)
			} else {
				handlerUseridListInt32 = newData[key].([]int32)
			}

			if len(handlerUseridListInt32) == 0 {
				return
			}
			var handlerUseridList []int
			for _, handlerUseridInt32 := range handlerUseridListInt32 {
				handlerUserid := int(handlerUseridInt32)
				handlerUseridList = append(handlerUseridList, handlerUserid)
			}
			var title string
			if newDataFormat["title"] == "" {
				title = oldDataFormat["title"]
			} else {
				title = newDataFormat["title"]
			}

			monitorHandlerField(wsId, handlerUseridList, userid)
			reassignNotice(wsId, tmplId, issueId, userid, handlerUseridList, title)
		}
	}

	document := bson.M{}
	document["issue_id"] = issueId
	document["ws_id"] = wsId
	document["tmpl_id"] = tmplId
	document["creator"] = userid
	document["action"] = action
	document["content"] = logContent
	document["created_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_log")
	_, err := collection.InsertOne(context.TODO(), &document)
	go rule_action_log.CheckRule(userid, wsId, tmplId, issueId, oldData, fieldsMap, changeData, action)
	if err != nil {
		return
	}
}

func monitorHandlerField(wsId int, useridList []int, userid int) {
	appSrv := new(app.CommonSrv)
	appSrv.LogUserOfCloseUser(wsId, useridList, userid)
}

func logForFile(userid int, wsId int, tmplId int, issueId string, fileName string, action string) {
	logContent := fileName
	document := bson.M{}
	document["issue_id"] = issueId
	document["ws_id"] = wsId
	document["tmpl_id"] = tmplId
	document["creator"] = userid
	document["action"] = action
	document["content"] = []string{logContent}
	document["created_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database("mark3").Collection("issue_log")
	_, err := collection.InsertOne(context.TODO(), &document)
	if err != nil {
		return
	}
}

func formatCreateData(tmplId int, newData bson.M, fieldsMap map[string]model.FieldModel, tmplCaller *tmpl.Caller) (newDataFormat map[string]string) {
	newDataFormat = make(map[string]string)
	for key, value := range newData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}
		var str string
		switch field.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom:
			str = strings.Join(value.([]string), ",")
			if str == "" {
				continue
			}
		case enum.FieldNumberCom, enum.FieldProgressCom:
			valueAssert, ok := value.(int32)
			if ok {
				str = fmt.Sprintf("%d", valueAssert)
			}
			if str == "" {
				continue
			}
		case enum.FieldPersonCom:
			valueAssert := value.([]int32)
			var useridList []int
			for _, userid := range valueAssert {
				useridList = append(useridList, int(userid))
			}
			users := GetUserList(useridList)
			var fullNameArr []string
			for _, user := range users {
				fullNameArr = append(fullNameArr, user.FullName)
			}
			str = strings.Join(fullNameArr, ",")
			if str == "" {
				continue
			}
		case enum.FieldLinkCom:
			valueAssert := value.(Link)
			var strArr []string
			if valueAssert.Name != "" {
				strArr = append(strArr, fmt.Sprintf("网址:%s", valueAssert.URL))
			}
			if valueAssert.URL != "" {
				strArr = append(strArr, fmt.Sprintf("名称:%s", valueAssert.Name))
			}
			str = strings.Join(strArr, ",")
			if str == "" {
				continue
			}
		case enum.FieldStatusCom:
			statusId := value.(int)
			status, err := tmplCaller.GetStatus(tmplId, statusId)
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
		case enum.FieldRelatedCom:
			valueAssert := value.([]string)
			for _, issueIdStr := range valueAssert {
				relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if err1 == nil {
					str = fmt.Sprintf("%s", relatedDocument["title"])
					break
				}
			}
		default:
			str = value.(string)
			if str == "" {
				continue
			}
		}
		newDataFormat[key] = str
	}
	return
}

func formatNewData(tmplId int, newData bson.M, fieldsMap map[string]model.FieldModel, tmplCaller *tmpl.Caller) (newDataFormat map[string]string) {
	newDataFormat = make(map[string]string)
	for key, value := range newData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}
		var str string
		switch field.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom:
			str = strings.Join(value.([]string), ",")
		case enum.FieldNumberCom, enum.FieldProgressCom:
			vInt32, ok := value.(int32)
			if ok {
				str = fmt.Sprintf("%d", vInt32)
			}
		case enum.FieldPersonCom:
			valueAssert := value.([]int32)
			var useridList []int
			for _, userid := range valueAssert {
				useridList = append(useridList, int(userid))
			}
			users := GetUserList(useridList)
			var fullNameArr []string
			for _, user := range users {
				fullNameArr = append(fullNameArr, user.FullName)
			}
			str = strings.Join(fullNameArr, ",")
		case enum.FieldLinkCom:
			valueAssert := value.(Link)
			var strArr []string
			if valueAssert.Name != "" {
				strArr = append(strArr, fmt.Sprintf("网址:%s", valueAssert.URL))
			}
			if valueAssert.URL != "" {
				strArr = append(strArr, fmt.Sprintf("名称:%s", valueAssert.Name))
			}
			str = strings.Join(strArr, ",")
		case enum.FieldStatusCom:
			statusId := value.(int)
			status, err := tmplCaller.GetStatus(tmplId, statusId)
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
		case enum.FieldRelatedCom:
			valueAssert := value.([]string)
			for _, issueIdStr := range valueAssert {
				relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if err1 == nil {
					str = fmt.Sprintf("%s", relatedDocument["title"])
					break
				}
			}
		default:
			str = value.(string)
		}
		newDataFormat[key] = str
	}
	return
}

func formatOldData(tmplId int, oldData bson.M, fieldsMap map[string]model.FieldModel, tmplCaller *tmpl.Caller) (oldDataFormat map[string]string) {
	oldDataFormat = make(map[string]string)
	for key, value := range oldData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}

		switch field.Mode {
		case enum.FieldMultiSelectCom, enum.FieldSelectCom:
			valueAssert := value.(primitive.A)
			var strArr []string
			for _, v := range valueAssert {
				vStr, ok := v.(string)
				if !ok {
					continue
				}
				strArr = append(strArr, vStr)
			}
			oldDataFormat[key] = strings.Join(strArr, ",")
		case enum.FieldNumberCom, enum.FieldProgressCom:
			vInt32, ok := value.(int32)
			var str string
			if ok {
				str = fmt.Sprintf("%d", vInt32)
			}
			oldDataFormat[key] = str
		case enum.FieldPersonCom:
			valueAssert := value.(primitive.A)
			var intArr []int
			for _, v := range valueAssert {
				vInt32, ok := v.(int32)
				if !ok {
					continue
				}
				intArr = append(intArr, int(vInt32))
			}
			users := GetUserList(intArr)
			var fullNameArr []string
			for _, user := range users {
				fullNameArr = append(fullNameArr, user.FullName)
			}
			oldDataFormat[key] = strings.Join(fullNameArr, ",")
		case enum.FieldLinkCom:
			valueAssert, ok := value.(primitive.M)
			if !ok {
				continue
			}
			var strArr []string
			for k, v := range valueAssert {
				vStr := v.(string)
				if k == "url" && vStr != "" {
					strArr = append(strArr, fmt.Sprintf("网址:%s", vStr))
				}
				if k == "name" && vStr != "" {
					strArr = append(strArr, fmt.Sprintf("名称:%s", vStr))
				}
			}
			oldDataFormat[key] = strings.Join(strArr, ",")
		case enum.FieldStatusCom:
			statusId := value.(int32)
			var str string
			status, err := tmplCaller.GetStatus(tmplId, int(statusId))
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
			oldDataFormat[key] = str
		case enum.FieldRelatedCom:
			valueAssert := value.(primitive.A)
			for _, v := range valueAssert {
				issueIdStr, ok := v.(string)
				if !ok {
					continue
				}
				relatedDocument, err1 := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if err1 == nil {
					oldDataFormat[key] = fmt.Sprintf("%s", relatedDocument["title"])
					break
				}
			}
		default:
			oldDataFormat[key] = value.(string)
		}
	}
	return
}
