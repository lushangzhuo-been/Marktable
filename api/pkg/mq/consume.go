package mq

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/pkg/email"
	enum "mark3/repository/db/enum/rule"
	enumTmpl "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/rule"
	tmplModel "mark3/repository/db/model/tmpl"
	userModel "mark3/repository/db/model/user"
	"mark3/service/tmpl"
	"strconv"
	"strings"
)

type Link struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func ConsumeSimple() {
	//声明队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := global.GVA_MQ.QueueDeclare(
		//队列名称
		"mark_update",
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := global.GVA_MQ.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	//forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理
			actionId := string(d.Body)
			log.Printf("Received a message: %s", d.Body)
			UpdateDocumentData(actionId)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}

func ConsumeSimpleEmail() {
	//声明队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := global.GVA_MQ.QueueDeclare(
		//队列名称
		"mark_email",
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := global.GVA_MQ.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	//forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理
			actionId := string(d.Body)
			log.Printf("Received a message: %s", actionId)
			SendEmail(actionId)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}

func SendEmail(ruleActionLogId string) {
	var action model.RuleActionLogModel
	global.GVA_DB.Where("id=?", ruleActionLogId).First(&action)
	if action.Id != 0 {
		subject := action.EmailTitle
		body := action.EmailContents
		to := strings.Split(action.EmailAddress, ",")
		if err := email.SendEmail(to, nil, nil, subject, body, ""); err != nil {
			global.GVA_LOG.Error(err.Error())
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusFail, "execute_at": common.GetCurrentTime()})
		} else {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusSuccess, "execute_at": common.GetCurrentTime()})
		}
	}

}

func UpdateDocumentData(ruleActionLogId string) {
	var action model.RuleActionLogModel
	global.GVA_DB.Where("id=?", ruleActionLogId).First(&action)
	if action.Id != 0 {
		collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
		objectID, err := primitive.ObjectIDFromHex(action.DataId)
		if err != nil {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusFail, "execute_at": common.GetCurrentTime()})
		}
		filter := bson.M{
			"$and": []bson.M{
				{"ws_id": action.WsId},
				{"tmpl_id": action.TmplId},
				{"_id": objectID},
			},
		}

		var oldDocument bson.M
		err = collection.FindOne(context.TODO(), filter).Decode(&oldDocument)
		if err != nil {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusFail, "execute_at": common.GetCurrentTime()})
			return
		}

		document := bson.M{}
		document["updated_at"] = common.GetCurrentTime()
		tmplCaller := new(tmpl.Caller)
		fieldsMap := tmplCaller.GetFieldsMap(action.TmplId)
		if field, ok := fieldsMap[action.FieldKey]; !ok {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusFail, "execute_at": common.GetCurrentTime()})
			return
		} else {
			// 自动化规则适配-关联字段
			if field.Mode == enumTmpl.FieldSelectCom || field.Mode == enumTmpl.FieldMultiSelectCom ||
				field.Mode == enumTmpl.FieldRelatedCom {
				var valStrList []string
				if action.FieldValue != "" {
					valStrList = strings.Split(action.FieldValue, ",")
				}
				document[action.FieldKey] = valStrList
			} else if field.Mode == enumTmpl.FieldPersonCom {
				var userIds []int
				userLists := strings.Split(action.FieldValue, ",")
				for _, s := range userLists {
					uid, err := strconv.Atoi(s)
					if err == nil {
						userIds = append(userIds, uid)
					}
				}
				document[action.FieldKey] = userIds
			} else if field.Mode == enumTmpl.FieldNumberCom || field.Mode == enumTmpl.FieldProgressCom || field.Mode == enumTmpl.FieldStatusCom {
				var numValue int
				intVal, err := strconv.Atoi(action.FieldValue)
				if err == nil {
					numValue = intVal
				}
				document[action.FieldKey] = numValue
			} else {
				document[action.FieldKey] = action.FieldValue
			}
		}

		update := bson.M{
			"$set": document,
		}
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusFail, "execute_at": common.GetCurrentTime()})
		} else {
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Where("id=? ", ruleActionLogId).Updates(map[string]interface{}{"status": enum.RuleActionLogStatusSuccess, "execute_at": common.GetCurrentTime()})
			// 添加编辑日志
			var user userModel.UserModel
			global.GVA_DB.Where("username=?", "MarkAdmin").First(&user)
			if user.Id != 0 {
				issueId := objectID.Hex()
				go documentLog(user.Id, action.WsId, action.TmplId, issueId, document, oldDocument, tmplCaller)
			}
		}
	}
}

func documentLog(userid int, wsId int, tmplId int, issueId string, newData bson.M, oldData bson.M, tmplCaller *tmpl.Caller) {
	action := "update"
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
	}

	document := bson.M{}
	document["issue_id"] = issueId
	document["ws_id"] = wsId
	document["tmpl_id"] = tmplId
	document["creator"] = userid
	document["action"] = action
	document["content"] = logContent
	document["created_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue_log")
	_, err := collection.InsertOne(context.TODO(), &document)

	if err != nil {
		return
	}
}

func formatCreateData(tmplId int, newData bson.M, fieldsMap map[string]tmplModel.FieldModel, tmplCaller *tmpl.Caller) (newDataFormat map[string]string) {
	newDataFormat = make(map[string]string)
	for key, value := range newData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}
		var str string
		switch field.Mode {
		case enumTmpl.FieldMultiSelectCom, enumTmpl.FieldSelectCom:
			str = strings.Join(value.([]string), ",")
			if str == "" {
				continue
			}
		case enumTmpl.FieldNumberCom, enumTmpl.FieldProgressCom:
			if valueAssert, ok1 := value.(int32); ok1 {
				str = fmt.Sprintf("%d", valueAssert)
			} else if valueInt, ok2 := value.(int); ok2 {
				str = fmt.Sprintf("%d", valueInt)
			}
			if str == "" {
				continue
			}
		case enumTmpl.FieldPersonCom:
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
		case enumTmpl.FieldLinkCom:
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
		case enumTmpl.FieldStatusCom:
			statusId := value.(int)
			status, err := tmplCaller.GetStatus(tmplId, statusId)
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
		case enumTmpl.FieldRelatedCom:
			valueAssert := value.([]string)
			for _, issueIdStr := range valueAssert {
				relatedDocument, _ := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if relatedDocument != nil {
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

func formatNewData(tmplId int, newData bson.M, fieldsMap map[string]tmplModel.FieldModel, tmplCaller *tmpl.Caller) (newDataFormat map[string]string) {
	newDataFormat = make(map[string]string)
	for key, value := range newData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}
		var str string
		switch field.Mode {
		case enumTmpl.FieldMultiSelectCom, enumTmpl.FieldSelectCom:
			str = strings.Join(value.([]string), ",")
		case enumTmpl.FieldNumberCom, enumTmpl.FieldProgressCom:
			if valueAssert, ok1 := value.(int32); ok1 {
				str = fmt.Sprintf("%d", valueAssert)
			} else if valueInt, ok2 := value.(int); ok2 {
				str = fmt.Sprintf("%d", valueInt)
			}
		case enumTmpl.FieldPersonCom:
			var useridList []int
			useridList, okInt := value.([]int)
			if !okInt {
				valueAssert, okInt1 := value.([]int32)
				if okInt1 {
					for _, userid := range valueAssert {
						useridList = append(useridList, int(userid))
					}
				}
			}

			users := GetUserList(useridList)
			var fullNameArr []string
			for _, user := range users {
				fullNameArr = append(fullNameArr, user.FullName)
			}
			str = strings.Join(fullNameArr, ",")
		case enumTmpl.FieldLinkCom:
			valueAssert := value.(Link)
			var strArr []string
			if valueAssert.Name != "" {
				strArr = append(strArr, fmt.Sprintf("网址:%s", valueAssert.URL))
			}
			if valueAssert.URL != "" {
				strArr = append(strArr, fmt.Sprintf("名称:%s", valueAssert.Name))
			}
			str = strings.Join(strArr, ",")
		case enumTmpl.FieldStatusCom:
			statusId := value.(int)
			status, err := tmplCaller.GetStatus(tmplId, statusId)
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
		case enumTmpl.FieldRelatedCom:
			valueAssert := value.([]string)
			for _, issueIdStr := range valueAssert {
				relatedDocument, _ := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if relatedDocument != nil {
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

func formatOldData(tmplId int, oldData bson.M, fieldsMap map[string]tmplModel.FieldModel, tmplCaller *tmpl.Caller) (oldDataFormat map[string]string) {
	oldDataFormat = make(map[string]string)
	for key, value := range oldData {
		field, ok := fieldsMap[key]
		if !ok {
			continue
		}

		switch field.Mode {
		case enumTmpl.FieldMultiSelectCom, enumTmpl.FieldSelectCom:
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
		case enumTmpl.FieldNumberCom, enumTmpl.FieldProgressCom:
			var str string
			if valueAssert, ok1 := value.(int32); ok1 {
				str = fmt.Sprintf("%d", valueAssert)
			} else if valueInt, ok2 := value.(int); ok2 {
				str = fmt.Sprintf("%d", valueInt)
			}
			oldDataFormat[key] = str
		case enumTmpl.FieldPersonCom:
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
		case enumTmpl.FieldLinkCom:
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
		case enumTmpl.FieldStatusCom:
			statusId := value.(int32)
			var str string
			status, err := tmplCaller.GetStatus(tmplId, int(statusId))
			if err != nil {
				str = ""
			} else {
				str = status.Name
			}
			oldDataFormat[key] = str
		case enumTmpl.FieldRelatedCom:
			valueAssert := value.(primitive.A)
			for _, v := range valueAssert {
				issueIdStr, ok1 := v.(string)
				if !ok1 {
					continue
				}
				relatedDocument, _ := GetDocument(field.WsId, field.RelatedTmplId, issueIdStr)
				if relatedDocument != nil {
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

func GetUserList(userIds []int) []userModel.UserModel {
	if len(userIds) == 0 {
		return nil
	}
	var users []userModel.UserModel
	global.GVA_DB.Where("id in ?", userIds).Find(&users)
	return users
}

func GetDocument(wsId int, tmplId int, issueId string) (resp bson.M, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": wsId},
			{"tmpl_id": tmplId},
			{"_id": objectID},
		},
	}

	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return nil, nil
	}
	document["issue_url"] = fmt.Sprintf(global.GVA_CONFIG.URL.Issue, wsId, tmplId, issueId)
	return document, nil
}
