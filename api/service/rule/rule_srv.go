package rule

import (
	"encoding/json"
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/rule"
	enumTmpl "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/rule"
	tmplModel "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/schedule"
	"mark3/service/app"
	"mark3/service/tmpl"
	commonTypes "mark3/types/common"
	"mark3/types/rule"
	"strconv"
	"strings"
)

type RuleSrv struct{}

func (s *RuleSrv) Create(userid int, req rule.CreateReq) (resp interface{}, err error) {
	// 动作数据解析
	var actionList []model.RuleActionModel
	if req.ActionList == "" {
		return nil, errors.New("请输入执行动作")
	}
	if err = json.Unmarshal([]byte(req.ActionList), &actionList); err != nil {
		return nil, errors.New("操作失败")
	}

	// 规则
	var ruleInfo model.RuleModel
	ruleInfo.WsId = req.WsId
	ruleInfo.TmplId = req.TmplId
	ruleInfo.Name = req.Name
	ruleInfo.RuleType = req.RuleType
	ruleInfo.IsEnable = enum.RuleEnableYes
	ruleInfo.IsDeleted = enum.RuleDeletedNo
	ruleInfo.Description = req.Description
	ruleInfo.Filter = req.Filter
	ruleInfo.TriggerDay = req.TriggerDay
	ruleInfo.TriggerTime = req.TriggerTime
	ruleInfo.FieldKey = req.FieldKey
	ruleInfo.OldValue = req.OldValue
	ruleInfo.NewValue = req.NewValue
	ruleInfo.CreatedBy = userid
	ruleInfo.CreatedAt = common.GetCurrentTime()
	ruleInfo.UpdatedAt = common.GetCurrentTime()

	// 保存规则
	if err = global.GVA_DB.Create(&ruleInfo).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 保存动作
	err = s.createOrUpdateAction(ruleInfo.Id, actionList)
	if err != nil {
		return nil, err
	}
	// 定时任务
	_, _ = s.AddOrUpdateSchedule(ruleInfo.Id)
	return
}

func (s *RuleSrv) Delete(req rule.DeleteReq) (resp interface{}, err error) {
	// 规则软删除
	if err = global.GVA_DB.Where("id=? and ws_id=? and tmpl_id=?", req.Id, req.WsId, req.TmplId).Updates(&model.RuleModel{
		IsDeleted: enum.RuleDeletedYes,
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	// 定时任务
	_, _ = s.AddOrUpdateSchedule(req.Id)

	return
}

func (s *RuleSrv) Update(req rule.UpdateReq) (resp interface{}, err error) {
	// 动作数据解析
	var actionList []model.RuleActionModel
	if req.ActionList == "" {
		return nil, errors.New("请输入执行动作")
	}
	if err = json.Unmarshal([]byte(req.ActionList), &actionList); err != nil {
		return nil, errors.New("操作失败")
	}

	//判断是否存在
	var detailVo model.RuleModel
	global.GVA_DB.Model(&model.RuleModel{}).Where("is_deleted='no' and id=?", req.Id).First(&detailVo)
	if detailVo.Id == 0 {
		return nil, errors.New("数据不存在")
	}

	// 更新规则
	detailVo.Name = req.Name
	detailVo.RuleType = req.RuleType
	detailVo.Description = req.Description
	detailVo.Filter = req.Filter
	detailVo.TriggerDay = req.TriggerDay
	detailVo.TriggerTime = req.TriggerTime
	detailVo.FieldKey = req.FieldKey
	detailVo.OldValue = req.OldValue
	detailVo.NewValue = req.NewValue
	detailVo.UpdatedAt = common.GetCurrentTime()

	if err = global.GVA_DB.Save(&detailVo).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	// 更新动作
	err = s.createOrUpdateAction(req.Id, actionList)
	if err != nil {
		return nil, err
	}
	// 定时任务
	_, _ = s.AddOrUpdateSchedule(req.Id)
	return
}

// 创建、修改动作
func (s *RuleSrv) createOrUpdateAction(ruleId int, actionList []model.RuleActionModel) (err error) {
	var notDelID []int
	for _, action := range actionList {
		if action.Id > 0 {
			if err = global.GVA_DB.Where("id=?", action.Id).Updates(&model.RuleActionModel{
				ActionType:    action.ActionType,
				FieldKey:      action.FieldKey,
				FieldValue:    action.FieldValue,
				EmailTitle:    action.EmailTitle,
				EmailContents: action.EmailContents,
				UserList:      action.UserList,
				IssueRoleList: action.IssueRoleList,
				IsDeleted:     enum.RuleActionDeletedNo,
			}).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
				return errors.New("操作失败")
			}
			notDelID = append(notDelID, action.Id)
		} else {
			action.IsDeleted = enum.RuleActionDeletedNo
			action.RuleId = ruleId
			global.GVA_DB.Model(&model.RuleActionModel{}).Create(&action)
			notDelID = append(notDelID, action.Id)
		}
	}
	global.GVA_DB.Model(&model.RuleActionModel{}).Where("rule_id=? and id not in ?", ruleId, notDelID).Updates(map[string]interface{}{"is_deleted": enum.RuleActionDeletedYes})
	return nil
}

func (s *RuleSrv) Detail(req rule.DetailReq) (resp interface{}, err error) {
	var detailVo rule.DetailVo
	global.GVA_DB.Model(&model.RuleModel{}).Where("is_deleted='no' and id=?", req.Id).First(&detailVo)

	if detailVo.Id != 0 {
		var tmplDetail tmplModel.TmplModel
		if err = global.GVA_DB.Where("id=? and ws_id=?", req.TmplId, req.WsId).First(&tmplDetail).Error; err != nil {
			return nil, errors.New("参数错误，流程不存在")
		}
		detailVo.TmplName = tmplDetail.Name
		// 查询流程字段
		tmplCaller := new(tmpl.Caller)
		fieldsMap := tmplCaller.GetFieldsMap(detailVo.TmplId)
		if detailVo.RuleType == "update" {
			if field, ok := fieldsMap[detailVo.FieldKey]; !ok {
				return nil, errors.New("字段不存在")
			} else {
				if field.Mode == enumTmpl.FieldPersonCom {
					if detailVo.OldValue != "" && detailVo.OldValue != "任意值" {
						users := app.GetUsersByUserIdStr(detailVo.OldValue)
						detailVo.OldValueUser = users
					}
					if detailVo.NewValue != "" && detailVo.NewValue != "任意值" {
						users := app.GetUsersByUserIdStr(detailVo.NewValue)
						detailVo.NewValueUser = users
					}
				}
				// 自动化规则适配-关联字段
				if field.Mode == enumTmpl.FieldRelatedCom {
					if detailVo.OldValue != "" && detailVo.OldValue != "任意值" {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, detailVo.OldValue)
						detailVo.OldValueDocuments = documents
					}
					if detailVo.NewValue != "" && detailVo.NewValue != "任意值" {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, detailVo.NewValue)
						detailVo.NewValueDocuments = documents
					}
				}
			}
		}

		if detailVo.Filter != "" {
			detailVo.FormatFilter = filter2Format(detailVo.Filter, fieldsMap)
		}

		var users []user.UserModel
		global.GVA_DB.Where("id = ?", detailVo.CreatedBy).Find(&users)
		detailVo.Creator = users
		var actionList []rule.ActionListVo
		global.GVA_DB.Raw("select * from rule_action where is_deleted='no' and rule_id=?", detailVo.Id).Scan(&actionList)

		for i, action := range actionList {
			if field, ok := fieldsMap[action.FieldKey]; ok {
				if field.Mode == enumTmpl.FieldPersonCom {
					users := app.GetUsersByUserIdStr(action.FieldValue)
					actionList[i].FieldValueUser = users
				}
				// 自动化规则适配-关联字段
				if field.Mode == enumTmpl.FieldRelatedCom {
					if len(action.FieldValue) > 0 {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, action.FieldValue)
						actionList[i].FieldValueDocuments = documents
					}
				}
			}
			// 查询邮件用户
			users := app.GetUsersByUserIdStr(action.UserList)
			actionList[i].UserListInfo = users
		}
		detailVo.ActionList = actionList
		return detailVo, nil
	} else {
		return nil, errors.New("数据不存在")
	}
}

func (s *RuleSrv) Page(req rule.PageReq) (resp interface{}, err error) {
	var countSql = "select count(id) from rule where is_deleted='no' and ws_id=? and tmpl_id=?"
	var sql = "select * from rule where is_deleted='no' and ws_id=? and tmpl_id=?"
	sqlParams := []interface{}{req.WsId, req.TmplId}
	if req.RuleTypeSelect != "" {
		ruleTypes := strings.Split(req.RuleTypeSelect, ",")
		countSql += " and rule_type in ? "
		sql += " and rule_type in ? "
		sqlParams = append(sqlParams, ruleTypes)
	}
	if req.IsEnableSelect != "" {
		isEnables := strings.Split(req.IsEnableSelect, ",")
		countSql += " and is_enable in ? "
		sql += " and is_enable in ? "
		sqlParams = append(sqlParams, isEnables)
	}
	var createdByUserIds []int
	if req.CreatedBySelect != "" {
		createdBys := strings.Split(req.CreatedBySelect, ",")
		for _, s := range createdBys {
			uid, err := strconv.Atoi(s)
			if err == nil {
				createdByUserIds = append(createdByUserIds, uid)
			}
		}
		countSql += " and created_by in ? "
		sql += " and created_by in ? "
		sqlParams = append(sqlParams, createdByUserIds)
	}
	var cnt int64
	global.GVA_DB.Raw(countSql, sqlParams...).Scan(&cnt)

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	var rules []rule.RulePageVo
	sql += " ORDER BY created_at desc limit ?, ?"
	sqlParams = append(sqlParams, (pageNum-1)*pageSize)
	sqlParams = append(sqlParams, pageSize)
	global.GVA_DB.Raw(sql, sqlParams...).Scan(&rules)
	for i, ruleDetail := range rules {
		var users []user.UserModel
		global.GVA_DB.Where("id = ?", ruleDetail.CreatedBy).Find(&users)
		rules[i].Creator = users
	}

	resp = commonTypes.BasePageResp{
		List: rules,
		Cnt:  cnt,
	}
	return
}

func (s *RuleSrv) Enable(req rule.EnableReq) (resp interface{}, err error) {
	if req.IsEnable == enum.RuleEnableYes || req.IsEnable == enum.RuleEnableNo {
		if err = global.GVA_DB.Where("id=?", req.Id).Updates(&model.RuleModel{
			IsEnable:  req.IsEnable,
			UpdatedAt: common.GetCurrentTime(),
		}).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("操作失败")
		}
		// todo 如果规则类型为定时任务，则启动或停止对应任务

	} else {
		return nil, nil
	}
	return
}

func (s *RuleSrv) RuleLog(req rule.RuleLogReq) (resp interface{}, err error) {
	var params []interface{}
	var sqlCount strings.Builder
	var sql strings.Builder

	sqlCount.WriteString("SELECT count(id) from rule_log WHERE ws_id=? and tmpl_id=? ")
	sql.WriteString("select * from rule_log where ws_id=? and tmpl_id=? ")
	params = append(params, req.WsId)
	params = append(params, req.TmplId)
	if req.RuleId > 0 {
		sqlCount.WriteString(" AND rule_id=? ")
		sql.WriteString(" AND rule_id=? ")
		params = append(params, req.RuleId)
	}
	var cnt int64
	global.GVA_DB.Raw(sqlCount.String(), params...).Scan(&cnt)

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}

	var ruleLog []model.RuleLogModel
	sql.WriteString(" ORDER BY execute_at desc limit ?, ? ")
	params = append(params, (pageNum-1)*pageSize)
	params = append(params, pageSize)
	global.GVA_DB.Raw(sql.String(), params...).Scan(&ruleLog)

	resp = commonTypes.BasePageResp{
		List: ruleLog,
		Cnt:  cnt,
	}
	return
}

func (s *RuleSrv) ActionLog(req rule.ActionLogReq) (resp interface{}, err error) {
	// 动作
	var params []interface{}
	var sql strings.Builder

	sql.WriteString("select * from rule_action_log where rule_execute_uuid=? ")
	params = append(params, req.RuleExecuteUuid)
	var actionLogList []rule.ActionLogListVo
	global.GVA_DB.Raw(sql.String(), params...).Scan(&actionLogList)
	tmplCaller := new(tmpl.Caller)
	fieldsMap := tmplCaller.GetFieldsMap(req.TmplId)
	for i, actionLog := range actionLogList {
		if actionLog.ActionType == enum.RuleActionTypeEmail {
			emailLists := strings.Split(actionLog.EmailAddress, ",")
			var users []user.UserModel
			global.GVA_DB.Where(" email in ?", emailLists).Find(&users)
			actionLogList[i].EmailUser = users
		}
		if actionLog.ActionType == enum.RuleActionTypeUpdate {
			if field, ok := fieldsMap[actionLog.FieldKey]; ok {
				if field.Mode == enumTmpl.FieldPersonCom {
					users := app.GetUsersByUserIdStr(actionLog.FieldValue)
					actionLogList[i].FieldValueUser = users
				}
				// 自动化规则适配-关联字段
				if field.Mode == enumTmpl.FieldRelatedCom {
					if len(actionLog.FieldValue) > 0 {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, actionLog.FieldValue)
						actionLogList[i].FieldValueDocuments = documents
					}
				}
			}
		}
	}

	// 规则
	var ruleLog model.RuleLogModel
	global.GVA_DB.Model(&model.RuleLogModel{}).Where(" rule_execute_uuid=?", req.RuleExecuteUuid).First(&ruleLog)
	var detailVo rule.DetailVo
	if err = json.Unmarshal([]byte(ruleLog.RuleData), &detailVo); err == nil {
		if detailVo.RuleType == "update" {
			if field, ok := fieldsMap[detailVo.FieldKey]; !ok {
				return nil, errors.New("字段不存在")
			} else {
				if field.Mode == enumTmpl.FieldPersonCom {
					if detailVo.OldValue != "" && detailVo.OldValue != "任意值" {
						users := app.GetUsersByUserIdStr(detailVo.OldValue)
						detailVo.OldValueUser = users
					}
					if detailVo.NewValue != "" && detailVo.NewValue != "任意值" {
						users := app.GetUsersByUserIdStr(detailVo.NewValue)
						detailVo.NewValueUser = users
					}
				}
				// 自动化规则适配-关联字段
				if field.Mode == enumTmpl.FieldRelatedCom {
					if detailVo.OldValue != "" && detailVo.OldValue != "任意值" {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, detailVo.OldValue)
						detailVo.OldValueDocuments = documents
					}
					if detailVo.NewValue != "" && detailVo.NewValue != "任意值" {
						documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, detailVo.NewValue)
						detailVo.NewValueDocuments = documents
					}
				}
			}
		}

		var users []user.UserModel
		global.GVA_DB.Where("id = ?", detailVo.CreatedBy).Find(&users)
		detailVo.Creator = users
		if detailVo.Filter != "" {
			detailVo.FormatFilter = filter2Format(detailVo.Filter, fieldsMap)
		}
	}

	resp = rule.RuleActionLogResp{
		RuleDetail: detailVo,
		List:       actionLogList,
		Cnt:        int64(len(actionLogList)),
	}
	return
}

func (s *RuleSrv) Config(req rule.ConfigReq) (resp interface{}, err error) {
	var config rule.ConfigResp
	config.RuleType = common.FormatConfig(enum.RuleTypeMap)
	config.RuleEnable = common.FormatConfig(enum.RuleEnableMap)
	config.RuleDeleted = common.FormatConfig(enum.RuleDeletedMap)
	config.RuleActionType = common.FormatConfig(enum.RuleActionTypeMap)
	config.RuleActionDeleted = common.FormatConfig(enum.RuleActionDeletedMap)
	config.ActionLogStatus = common.FormatConfig(enum.RuleActionLogStatusMap)

	return config, nil
}

func (s *RuleSrv) getActionList(ruleID int) {

}

func (s *RuleSrv) AddOrUpdateSchedule(ruleId int) (resp interface{}, err error) {
	var ruleDetail model.RuleModel
	global.GVA_DB.Model(&model.RuleModel{}).Where("id=?", ruleId).First(&ruleDetail)

	if ruleDetail.Id != 0 && ruleDetail.RuleType == "schedule" {
		schedule.UpdateTaskCron(ruleDetail)
		return ruleDetail, nil
	} else if ruleDetail.Id != 0 {
		return nil, errors.New("非定时规则")
	} else {
		return nil, errors.New("数据不存在")
	}
}

// DoRuleCron 手动执行单个任务
func (s *RuleSrv) DoRuleCron(req rule.DetailReq) (resp interface{}, err error) {
	var ruleDetail model.RuleModel
	global.GVA_DB.Model(&model.RuleModel{}).Where("id=?", req.Id).First(&ruleDetail)
	if ruleDetail.Id != 0 {
		schedule.ExecuteTaskCron(ruleDetail)
		return ruleDetail, nil
	} else {
		return nil, errors.New("数据不存在")
	}
}

func filter2Format(filter string, fieldsMap map[string]tmplModel.FieldModel) (formatFilter string) {
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
		case enumTmpl.FieldPersonCom:
			var userList []user.UserModel
			if useridStr, oks := filterAtom.Value.(string); oks {
				userList = app.GetUsersByUserIdStr(useridStr)
			} else if useridF, okf := filterAtom.Value.(float64); okf {
				useridS := strconv.Itoa(int(useridF))
				userList = app.GetUsersByUserIdStr(useridS)
			} else {
				continue
			}

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = userList
		case enumTmpl.FieldStatusCom:
			var status tmplModel.StatusModel
			if statusIdStr, oks := filterAtom.Value.(string); oks {
				statusId, _ := strconv.Atoi(statusIdStr)
				status = app.GetStatusInfo(statusId)
			} else if statusIdF, okf := filterAtom.Value.(float64); okf {
				statusId := int(statusIdF)
				status = app.GetStatusInfo(statusId)
			} else {
				continue
			}

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = status
		case enumTmpl.FieldRelatedCom:
			issueIdStr := filterAtom.Value.(string)
			documents := app.GetDocumentsByIdStr(field.WsId, field.RelatedTmplId, issueIdStr)

			formatFilterAtom.FieldKey = filterAtom.FieldKey
			formatFilterAtom.Op = filterAtom.Op
			formatFilterAtom.Value = documents
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
