package rule_action_log

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/pkg/mq"
	enum "mark3/repository/db/enum/rule"
	enumTmpl "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/rule"
	tmplModel "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/service/app"
	"mark3/service/tmpl"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckRule(userid int, wsId int, tmplId int, issueId string, oldData bson.M, fieldsMap map[string]tmplModel.FieldModel,
	changeData map[string]struct {
		OldItem string
		NewItem string
	}, action string) {
	var ruleList []model.RuleModel
	if action == "create" {
		global.GVA_DB.Model(model.RuleModel{}).Select("*").Where("rule_type='create' and is_enable='yes' and is_deleted='no' and ws_id=? and tmpl_id=? ", wsId, tmplId).Find(&ruleList)
		for _, rule := range ruleList {
			document, err := GetOneDocument(rule, issueId)
			if err != nil {
				continue
			}
			AddRuleActionLog(rule, issueId, document)
		}
	} else if action == "update" || action == "click" || action == "switch" {
		global.GVA_DB.Model(model.RuleModel{}).Select("*").Where("rule_type='update' and is_enable='yes' and is_deleted='no' and ws_id=? and tmpl_id=? ", wsId, tmplId).Find(&ruleList)
		for _, rule := range ruleList {
			if _, ok := changeData[rule.FieldKey]; ok {
				document, err := GetOneDocument(rule, issueId)
				// todo 非任意值判断
				//if rule.OldValue != "" && rule.OldValue != "任意值" {
				//	CheckFieldUpdateRule(rule.FieldKey, rule.OldValue, oldData, fieldsMap)
				//}
				//if rule.NewValue != "" && rule.NewValue != "任意值" {
				//	CheckFieldUpdateRule(rule.FieldKey, rule.NewValue, document, fieldsMap)
				//}

				if err != nil {
					continue
				}
				AddRuleActionLog(rule, issueId, document)
			} else {
				continue
			}
		}
	}
}

func DelRuleAction(wsId int, tmplId int, documents []bson.M) {
	var ruleList []model.RuleModel
	global.GVA_DB.Model(model.RuleModel{}).Select("*").Where("rule_type='delete' and is_enable='yes' and is_deleted='no' and ws_id=? and tmpl_id=? ", wsId, tmplId).Find(&ruleList)
	for _, rule := range ruleList {
		for _, document := range documents {
			if id, ok := document["_id"]; !ok {
				continue
			} else {
				objID := id.(primitive.ObjectID)
				issueId := objID.Hex()
				AddRuleActionLog(rule, issueId, document)
			}
		}
	}
}

func AddRuleActionLog(rule model.RuleModel, issueId string, document bson.M) {
	var actionList []model.RuleActionModel
	global.GVA_DB.Where("is_deleted='no' and rule_id=?", rule.Id).Find(&actionList)
	actionMap := make(map[int]model.RuleActionModel, len(actionList))
	uuidStr := uuid.New().String()
	if len(actionList) > 0 {
		var ruleLog model.RuleLogModel
		ruleLog.WsId = rule.WsId
		ruleLog.TmplId = rule.TmplId
		ruleLog.RuleId = rule.Id
		ruleLog.RuleExecuteUuid = uuidStr
		ruleLog.RuleName = rule.Name
		ruleLog.RuleType = rule.RuleType
		ruleLog.ExecuteAt = common.GetCurrentTime()
		ruleLog.DataId = issueId
		ruleLog.DataName = fmt.Sprintf("%s", document["title"])

		ruleData, rErr := json.Marshal(rule)
		if rErr != nil {
			ruleLog.RuleData = ""
		} else {
			ruleLog.RuleData = string(ruleData)
		}

		ruleActionData, aErr := json.Marshal(actionList)
		if aErr != nil {
			ruleLog.RuleActionData = ""
		} else {
			ruleLog.RuleActionData = string(ruleActionData)
		}

		global.GVA_DB.Model(&model.RuleLogModel{}).Create(&ruleLog)
	}
	for _, action := range actionList {
		actionMap[action.Id] = action
		if action.ActionType == enum.RuleActionTypeEmail {
			// 插入动作执行日志表
			var actionLog model.RuleActionLogModel
			actionLog.WsId = rule.WsId
			actionLog.TmplId = rule.TmplId
			actionLog.RuleId = rule.Id
			actionLog.RuleExecuteUuid = uuidStr
			actionLog.RuleActionId = action.Id
			actionLog.Status = "no"
			actionLog.ActionType = action.ActionType
			actionLog.DataId = issueId
			actionLog.EmailTitle = action.EmailTitle
			// 解析邮件内容，获取字段数据
			contents := GetEmailContents(rule, action.EmailContents, document, issueId)
			actionLog.EmailContents = contents
			// 查询人员字段相关的邮箱
			emails := GetUserEmail(action, document)
			//actionLog.CreatedAt = common.GetCurrentTime()
			if emails != "" {
				actionLog.EmailAddress = emails
				global.GVA_DB.Model(&model.RuleActionLogModel{}).Create(&actionLog)
				// 生产消息
				mq.PublishSimple("mark_email", strconv.Itoa(actionLog.Id))
			}
		} else if action.ActionType == enum.RuleActionTypeUpdate {
			// 插入动作执行日志表
			var field tmplModel.FieldModel
			if err := global.GVA_DB.Where("tmpl_id=? and field_key=?", rule.TmplId, action.FieldKey).First(&field).Error; err != nil {
				continue
			}

			var actionLog model.RuleActionLogModel
			actionLog.WsId = rule.WsId
			actionLog.TmplId = rule.TmplId
			actionLog.RuleId = rule.Id
			actionLog.RuleExecuteUuid = uuidStr
			actionLog.RuleActionId = action.Id
			actionLog.Status = "no"
			actionLog.ActionType = action.ActionType
			actionLog.DataId = issueId
			actionLog.FieldKey = action.FieldKey
			actionLog.FieldValue = action.FieldValue
			//actionLog.CreatedAt = common.GetCurrentTime()
			global.GVA_DB.Model(&model.RuleActionLogModel{}).Create(&actionLog)
			// 生产消息
			mq.PublishSimple("mark_update", strconv.Itoa(actionLog.Id))
		}
	}
}

func GetOneDocument(rule model.RuleModel, issueId string) (bson.M, error) {
	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMap(rule.TmplId)

	filter, err := GetFilter(rule, fieldsMap)
	if err != nil {
		return nil, err
	}
	objectID, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		return nil, err
	}
	filter["_id"] = objectID

	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func GetDocuments(rule model.RuleModel) ([]bson.M, error) {
	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMap(rule.TmplId)

	filter, err := GetFilter(rule, fieldsMap)
	if err != nil {
		return nil, err
	}
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue")
	// mgField := bson.D{{"_id", 1}}
	// cursor, err := collection.Find(context.TODO(), filter, options.Find().SetProjection(mgField))
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}
	return documents, nil
}

func GetFilter(rule model.RuleModel, fieldsMap map[string]tmplModel.FieldModel) (filter bson.M, err error) {
	// 设置时间过滤条件
	filter = bson.M{
		"ws_id":   rule.WsId,
		"tmpl_id": rule.TmplId,
	}

	if rule.RuleType == "schedule" {
		now := time.Now()
		var fieldDay string
		if rule.TriggerDay > 0 {
			fieldDay = now.AddDate(0, 0, -rule.TriggerDay).Format("2006-01-02")
		} else if rule.TriggerDay < 0 {
			fieldDay = now.AddDate(0, 0, -rule.TriggerDay).Format("2006-01-02")
		} else {
			fieldDay = now.Format("2006-01-02")
		}
		filter[rule.FieldKey] = bson.M{"$regex": fieldDay, "$options": "i"}
		//if fieldsMap[rule.FieldKey].Expr == "Ymd" {
		//	filter[rule.FieldKey] = fieldDay
		//} else {
		//	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", fieldDay), time.Local)
		//	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", fieldDay), time.Local)
		//
		//	filter[rule.FieldKey] = bson.M{
		//		"$gte": startTime,
		//		"$lte": endTime,
		//	}
		//}
	}

	type filterAtom struct {
		FieldKey string      `json:"field_key"`
		Op       string      `json:"op"`
		Value    interface{} `json:"value"`
	}

	var filterList []filterAtom
	if err = json.Unmarshal([]byte(rule.Filter), &filterList); err != nil {
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
				if field.Mode == enumTmpl.FieldTextCom || field.Mode == enumTmpl.FieldTextareaCom {
					cond = bson.M{atom.FieldKey: bson.M{"$regex": atom.Value, "$options": "i"}}
				} else if field.Mode == enumTmpl.FieldLinkCom {
					key := fmt.Sprintf("%s.%s", atom.FieldKey, "name")
					cond = bson.M{key: bson.M{atom.Op: atom.Value}}
				} else if field.Mode == enumTmpl.FieldNumberCom || field.Mode == enumTmpl.FieldProgressCom ||
					field.Mode == enumTmpl.FieldPersonCom || field.Mode == enumTmpl.FieldStatusCom {
					// 数值类型字段
					if intValStr, oks := atom.Value.(string); oks {
						intVal, okInt := strconv.Atoi(intValStr)
						if okInt == nil {
							cond = bson.M{atom.FieldKey: bson.M{atom.Op: intVal}}
						} else {
							return filter, errors.New("参数异常")
						}
					} else if statusIdF, okf := atom.Value.(float64); okf {
						intVal := int(statusIdF)
						cond = bson.M{atom.FieldKey: bson.M{atom.Op: intVal}}
					} else {
						return filter, errors.New("参数异常")
					}
				} else {
					cond = bson.M{atom.FieldKey: bson.M{atom.Op: atom.Value}}
				}
			}
			filterInner = append(filterInner, cond)
		}
	}

	filter["$and"] = filterInner
	return filter, nil
}

func GetUserEmail(action model.RuleActionModel, document bson.M) string {
	var userIds []int

	if action.UserList != "" {
		userLists := strings.Split(action.UserList, ",")
		for _, s := range userLists {
			uid, err := strconv.Atoi(s)
			if err == nil {
				userIds = append(userIds, uid)
			}
		}
	}

	if action.IssueRoleList != "" {
		userFields := strings.Split(action.IssueRoleList, ",")
		for _, userField := range userFields {
			fieldUsers := document[userField]
			if fieldUsers == nil {
				continue
			}
			val := reflect.ValueOf(fieldUsers)
			for i := 0; i < val.Len(); i++ {
				v := val.Index(i)
				vInt32, ok := v.Interface().(int32)
				if ok {
					userIds = append(userIds, int(vInt32))
				}
			}
		}
	}

	if len(userIds) > 0 {
		var emails []string
		var users []user.UserModel
		global.GVA_DB.Where("id in ?", userIds).Find(&users)
		for _, user := range users {
			if user.Email == "" {
				continue
			} else {
				emails = append(emails, user.Email)
			}
		}
		return strings.Join(emails, ",")
	} else {
		return ""
	}
}

func GetEmailContents(rule model.RuleModel, contents string, document bson.M, issueId string) string {
	if contents == "" {
		return ""
	}
	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMapName(rule.TmplId)
	// 定义span标签的正则表达式
	spanRegexp := regexp.MustCompile(`<span[^>]*>(.*?)</span>`)
	// 使用正则表达式找到所有span标签及其内容
	matches := spanRegexp.FindAllStringSubmatch(contents, -1)

	// 遍历所有匹配项，并替换内容
	for _, match := range matches {
		fieldName := match[1]
		if field, ok := fieldsMap[fieldName]; ok {
			if document[field.FieldKey] == nil {
				contents = strings.Replace(contents, match[0], "", 1)
			} else if field.FieldKey == "title" {
				issueUrl := fmt.Sprintf(global.GVA_CONFIG.URL.Issue, field.WsId, field.TmplId, issueId)
				newStr := fmt.Sprintf("<a href=\"%s\">%s</a>", issueUrl, document[field.FieldKey])
				contents = strings.Replace(contents, match[0], newStr, 1)
			} else if field.Mode == enumTmpl.FieldSelectCom || field.Mode == enumTmpl.FieldMultiSelectCom {
				newStr := fmt.Sprintf("%s", document[field.FieldKey])
				contents = strings.Replace(contents, match[0], newStr, 1)
			} else if field.Mode == enumTmpl.FieldPersonCom {
				// 查询人员名称并替换
				fieldUsers := document[field.FieldKey]
				if fieldUsers == nil {
					contents = strings.Replace(contents, match[0], "", 1)
					continue
				}
				val := reflect.ValueOf(fieldUsers)
				var userIds []int
				for i := 0; i < val.Len(); i++ {
					v := val.Index(i)
					vInt32, ok := v.Interface().(int32)
					if ok {
						userIds = append(userIds, int(vInt32))
					}
				}
				if len(userIds) > 0 {
					var userNames []string
					var users []user.UserModel
					global.GVA_DB.Where("id in ?", userIds).Find(&users)
					for _, user := range users {
						if user.FullName == "" {
							continue
						} else {
							userNames = append(userNames, user.FullName)
						}
					}
					contents = strings.Replace(contents, match[0], strings.Join(userNames, "、"), 1)
				} else {
					contents = strings.Replace(contents, match[0], "", 1)
				}
			} else if field.Mode == enumTmpl.FieldNumberCom || field.Mode == enumTmpl.FieldProgressCom {
				numValue, ok := document[field.FieldKey].(int32)
				var newStr string
				if ok {
					newStr = fmt.Sprintf("%d", numValue)
				}
				contents = strings.Replace(contents, match[0], newStr, 1)
			} else if field.Mode == enumTmpl.FieldStatusCom {
				statusId := document[field.FieldKey].(int32)
				var newStr string
				status, err := tmplCaller.GetStatus(rule.TmplId, int(statusId))
				if err != nil {
					newStr = ""
				} else {
					newStr = status.Name
				}
				contents = strings.Replace(contents, match[0], newStr, 1)
			} else if field.Mode == enumTmpl.FieldRelatedCom {
				// 自动化规则适配-关联字段
				fieldDocuments := document[field.FieldKey]
				if fieldDocuments == nil {
					contents = strings.Replace(contents, match[0], "", 1)
					continue
				}
				valueAssert := fieldDocuments.(primitive.A)
				var issueIds []string
				for _, issueId := range valueAssert {
					issueIdStr, okr := issueId.(string)
					if okr {
						issueIds = append(issueIds, issueIdStr)
					}
				}
				if len(issueIds) > 0 {
					var issueTitles []string
					issueIdsStr := strings.Join(issueIds, ",")
					documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, issueIdsStr)
					if documents != nil {
						for _, relatedDocument := range documents {
							title := fmt.Sprintf("%s", relatedDocument["title"])
							issueTitles = append(issueTitles, title)
						}
						contents = strings.Replace(contents, match[0], strings.Join(issueTitles, "、"), 1)
					} else {
						contents = strings.Replace(contents, match[0], "", 1)
					}
				} else {
					contents = strings.Replace(contents, match[0], "", 1)
				}
			} else {
				newStr := fmt.Sprintf("%s", document[field.FieldKey])
				contents = strings.Replace(contents, match[0], newStr, 1)
			}
		} else {
			contents = strings.Replace(contents, match[0], "", 1)
		}
	}

	return contents
}

func CheckFieldUpdateRule(fieldKey string, compareValue string, document bson.M, fieldsMap map[string]tmplModel.FieldModel) bool {
	if field, ok := fieldsMap[fieldKey]; !ok {
		return false
	} else if document[fieldKey] == nil {
		return false
	} else {
		// 自动化规则适配-
		if field.Mode == enumTmpl.FieldSelectCom || field.Mode == enumTmpl.FieldMultiSelectCom ||
			field.Mode == enumTmpl.FieldPersonCom || field.Mode == enumTmpl.FieldRelatedCom {
			compareValueList := strings.Split(compareValue, ",")
			fmt.Println(compareValueList)
			// 比较两个list
			if selectValues, ok := document[fieldKey].([]interface{}); ok {
				for _, s := range selectValues {
					vStr := fmt.Sprintf("%d", s)
					if !common.InArray(vStr, compareValueList) {
						return false
					}
				}
				return true
			}
		} else if field.Mode == enumTmpl.FieldStatusCom {
			statusLists := strings.Split(compareValue, ",")
			// 判断 documentData[fieldKey] 是否在 数组中
			if statusId, ok := document[fieldKey].(int); ok {
				statusStr := fmt.Sprintf("%d", statusId)
				if common.InArray(statusStr, statusLists) {
					return true
				}
			}
		} else {
			// 其余类型字段不做比较
			return true
		}
	}
	return false
}
