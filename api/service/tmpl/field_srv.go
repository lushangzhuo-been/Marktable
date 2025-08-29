package tmpl

import (
	"encoding/json"
	"errors"
	"fmt"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	types "mark3/types/tmpl"
	"strings"
)

type FieldSrv struct{}

func (s *FieldSrv) Config() (resp interface{}, err error) {
	var config types.FieldConfigResp
	config.ModeConfig = common.FormatConfig(enum.FieldModeMap)
	config.NumberComConfig = common.FormatConfig(enum.FieldNumberComMap)
	config.TimeComConfig = common.FormatConfig(enum.FieldTimeComMap)
	config.PersonComConfig = common.FormatConfig(enum.FieldPersonComMap)
	return config, nil
}

func (s *FieldSrv) List(req types.FieldListReq) (resp interface{}, err error) {
	var fields []model.FieldModel
	if req.HasStatus == "yes" {
		global.GVA_DB.Where("tmpl_id=? and (level in ? or field_key = 'status')", req.TmplId, []string{"level2", "level3"}).Find(&fields)
	} else {
		global.GVA_DB.Where("tmpl_id=? and level in ?", req.TmplId, []string{"level2", "level3"}).Find(&fields)
	}
	for i, field := range fields {
		if field.Mode == enum.FieldStatusCom {
			var statusList []model.StatusModel
			global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&statusList)
			if statusList != nil {
				statusStr, sErr := json.Marshal(statusList)
				if sErr == nil {
					fields[i].EnumValues = string(statusStr)
				}
			}
		}
	}

	// 字段排序
	var coordinate model.ScreenCoordinateModel
	global.GVA_DB.Where("tmpl_id=? and module=?", req.TmplId, enum.ScreenCoordinateModuleField).First(&coordinate)
	if coordinate.Id != 0 {
		coordinateSlice := strings.Split(coordinate.Coordinate, ",")
		var dataMap = make(map[string]model.FieldModel)
		for _, vo := range fields {
			dataMap[vo.FieldKey] = vo
		}

		var sort []model.FieldModel
		for _, key := range coordinateSlice {
			if data, ok := dataMap[key]; ok {
				sort = append(sort, data)
				delete(dataMap, key)
			}
		}

		for _, vo := range dataMap {
			sort = append(sort, vo)
		}
		return sort, nil
	}

	return fields, nil
}

func (s *FieldSrv) Info(req types.FieldInfoReq) (resp interface{}, err error) {
	var field model.FieldModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&field).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	return field, nil
}

func (s *FieldSrv) Create(req types.FieldCreateReq) (resp interface{}, err error) {
	if _, ok := enum.FieldModeMap[req.Mode]; !ok {
		return nil, errors.New("参数错误")
	}

	var field model.FieldModel
	if err = global.GVA_DB.Where("tmpl_id=? and name=?", req.TmplId, req.Name).First(&field).Error; err == nil {
		return nil, errors.New("字段命名重复，请重新命名")
	}

	field = model.FieldModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		FieldKey:  fmt.Sprintf("custom_field_%d_%s", req.TmplId, common.GetUUID()),
		Name:      req.Name,
		Desc:      req.Desc,
		Mode:      req.Mode,
		Level:     "level3",
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}

	switch req.Mode {
	case enum.FieldTextCom, enum.FieldTextareaCom, enum.FieldHtmlTextCom, enum.FieldLinkCom:
	case enum.FieldSelectCom, enum.FieldMultiSelectCom:
		if req.EnumValues == "" {
			return nil, errors.New("请填写枚举值")
		}
		var enumValues []types.FieldEnumValue
		err = json.Unmarshal([]byte(req.EnumValues), &enumValues)
		if err != nil {
			return nil, errors.New("枚举值有误")
		}
		for _, enumValue := range enumValues {
			if enumValue.Value == "" || enumValue.Color == "" {
				return nil, errors.New("枚举值有误")
			}
		}
		enumValuesByte, err := json.Marshal(enumValues)
		if err != nil {
			return nil, errors.New("枚举值有误")
		}
		field.EnumValues = string(enumValuesByte)
	case enum.FieldNumberCom:
		field.Expr = req.Expr
	case enum.FieldPersonCom:
		if _, ok := enum.FieldPersonComMap[req.Expr]; !ok {
			return nil, errors.New("参数错误")
		}
		field.Expr = req.Expr
	case enum.FieldRelatedCom:
		field.RelatedTmplId = req.RelatedTmplId
	default:
	}
	if err = global.GVA_DB.Create(&field).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *FieldSrv) Update(req types.FieldUpdateReq) (resp interface{}, err error) {
	var field model.FieldModel
	//判断是否重名
	if err = global.GVA_DB.Where("tmpl_id=? and id !=? and name=?", req.TmplId, req.Id, req.Name).First(&field).Error; err == nil {
		return nil, fmt.Errorf("[%s]字段命名重复，请重新命名", req.Name)
	}

	//判断是否存在
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&field).Error; err != nil {
		return nil, errors.New("参数错误")
	}

	switch field.Mode {
	case enum.FieldTextCom, enum.FieldTextareaCom, enum.FieldHtmlTextCom, enum.FieldLinkCom:
	case enum.FieldSelectCom, enum.FieldMultiSelectCom:
		if req.EnumValues == "" {
			return nil, errors.New("请填写枚举值")
		}
		var enumValues []types.FieldEnumValue
		err = json.Unmarshal([]byte(req.EnumValues), &enumValues)
		if err != nil {
			return nil, errors.New("枚举值有误")
		}
		for _, enumValue := range enumValues {
			if enumValue.Value == "" || enumValue.Color == "" {
				return nil, errors.New("枚举值有误")
			}
		}
		enumValuesByte, err := json.Marshal(enumValues)
		if err != nil {
			return nil, errors.New("枚举值有误")
		}
		field.EnumValues = string(enumValuesByte)
	case enum.FieldNumberCom:
		field.Expr = req.Expr
	case enum.FieldPersonCom:
		if _, ok := enum.FieldPersonComMap[req.Expr]; !ok {
			return nil, errors.New("参数错误")
		}
		field.Expr = req.Expr
	case enum.FieldRelatedCom:
		if field.RelatedTmplId != req.RelatedTmplId {
			return nil, errors.New("关联字段禁止编辑")
		}
	default:
	}
	field.Name = req.Name
	field.Desc = req.Desc
	field.UpdatedAt = common.GetCurrentTime()
	if err = global.GVA_DB.Save(&field).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}

	return
}

func (s *FieldSrv) GetAllPersonCom(req types.FieldGetAllPersonComReq) (resp interface{}, err error) {
	var fields []model.FieldModel
	global.GVA_DB.Where("tmpl_id=? and mode=?", req.TmplId, enum.FieldPersonCom).Find(&fields)
	return fields, nil
}

func (s *FieldSrv) GetReadOnlyRule(req types.FieldGetReadOnlyRuleReq) (resp interface{}, err error) {
	var field model.FieldModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&field).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}

	if len(field.ReadOnlyRule) == 0 {
		return
	}

	type ReadOnlyRule struct {
		StatusList       []int               `json:"status_list"`
		UserList         []int               `json:"user_list"`
		IssueRoleList    []string            `json:"issue_role_list"`
		UserListFormat   []user.UserModel    `json:"user_list_format"`
		StatusListFormat []model.StatusModel `json:"status_list_format"`
	}
	var readOnlyRule ReadOnlyRule
	if err = json.Unmarshal([]byte(field.ReadOnlyRule), &readOnlyRule); err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("json解析错误")
	}

	if len(readOnlyRule.StatusList) > 0 {
		var statusList []model.StatusModel
		if err = global.GVA_DB.Where("id in (?)", readOnlyRule.StatusList).Find(&statusList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误")
		}
		readOnlyRule.StatusListFormat = statusList
	}

	if len(readOnlyRule.UserList) > 0 {
		var userList []user.UserModel
		if err = global.GVA_DB.Where("id in (?)", readOnlyRule.UserList).Find(&userList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误")
		}
		readOnlyRule.UserListFormat = userList
	}
	return readOnlyRule, nil
}

func (s *FieldSrv) UpdateReadOnlyRule(req types.FieldUpdateReadOnlyRuleReq) (resp interface{}, err error) {
	var statusList = []int{}
	var userList = []int{}
	var issueRoleList = []string{}
	if req.StatusList != "" {
		if err = json.Unmarshal([]byte(req.StatusList), &statusList); err != nil {
			return nil, err
		}
	}
	if req.UserList != "" {
		if err = json.Unmarshal([]byte(req.UserList), &userList); err != nil {
			return nil, err
		}
	}
	if req.IssueRoleList != "" {
		if err = json.Unmarshal([]byte(req.IssueRoleList), &issueRoleList); err != nil {
			return nil, err
		}
	}

	if len(userList) > 0 || len(issueRoleList) > 0 {
		if len(statusList) == 0 {
			return nil, errors.New("只读状态不能为空，必须要填写")
		}
	}

	var readOnlyRule = struct {
		StatusList    []int    `json:"status_list"`
		UserList      []int    `json:"user_list"`
		IssueRoleList []string `json:"issue_role_list"`
	}{
		StatusList:    statusList,
		UserList:      userList,
		IssueRoleList: issueRoleList,
	}

	res, err := json.Marshal(readOnlyRule)
	if err != nil {
		return nil, err
	}

	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Updates(&model.FieldModel{
		ReadOnlyRule: string(res),
		UpdatedAt:    common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *FieldSrv) Delete(req types.FieldDeleteReq) (resp interface{}, err error) {
	var field model.FieldModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&field).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("数据不存在")
	}
	if field.Level != "level3" {
		return nil, errors.New("level1、leve2级别字段不支持删除")
	}
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).Delete(&model.FieldModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *FieldSrv) Enumeration(req types.FieldListReq) (resp interface{}, err error) {
	var fields []model.FieldModel
	//状态、人员、单选、多选
	global.GVA_DB.Where("tmpl_id=? and mode in ('status_com', 'person_com', 'select_com', 'multi_select_com')", req.TmplId).Find(&fields)
	for i, field := range fields {
		if field.Mode == enum.FieldStatusCom {
			var statusList []model.StatusModel
			global.GVA_DB.Where("tmpl_id=?", req.TmplId).Find(&statusList)
			if statusList != nil {
				statusStr, sErr := json.Marshal(statusList)
				if sErr == nil {
					fields[i].EnumValues = string(statusStr)
				}
			}
		}
	}
	return fields, nil
}
