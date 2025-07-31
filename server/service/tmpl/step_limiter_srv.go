package tmpl

import (
	"encoding/json"
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	types "mark3/types/tmpl"
)

type StepLimiterSrv struct{}

func (s *StepLimiterSrv) List(req types.StepLimiterListReq) (resp interface{}, err error) {
	type StepLimiterVo struct {
		model.StepLimiterModel
		Mode_Desc string `json:"mode_desc"`
	}
	var limits []StepLimiterVo
	if err = global.GVA_DB.Where("tmpl_id=? and step_id=?", req.TmplId, req.StepId).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}

	for k, v := range limits {
		desc, ok := enum.ButtonLimitModeMap[v.Mode]
		if ok {
			limits[k].Mode_Desc = desc
		} else {
			limits[k].Mode_Desc = ""
		}
	}

	return limits, nil
}

func (s *StepLimiterSrv) Info(req types.StepLimiterInfoReq) (resp interface{}, err error) {
	var limit model.StepLimiterModel
	if err = global.GVA_DB.Where("tmpl_id=? and id=?", req.TmplId, req.Id).First(&limit).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}

	type Rule struct {
		Name           string           `json:"name"`
		UserList       []int            `json:"user_list"`
		IssueRoleList  []string         `json:"issue_role_list"`
		UserListFormat []user.UserModel `json:"user_list_format"`
	}
	var rule Rule
	if err = json.Unmarshal([]byte(limit.Rule), &rule); err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("json解析错误")
	}

	if len(rule.UserList) > 0 {
		var userList []user.UserModel
		if err = global.GVA_DB.Where("id in (?)", rule.UserList).Find(&userList).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("参数错误")
		}
		rule.UserListFormat = userList
	}
	rule.Name = limit.Name
	return rule, nil
}

func (s *StepLimiterSrv) Create(req types.StepLimiterCreateReq) (resp interface{}, err error) {
	if req.UserList == "" && req.IssueRoleList == "" {
		return nil, errors.New("参数错误")
	}

	var userList []int
	var issueRoleList []string
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

	var rule = struct {
		UserList      []int    `json:"user_list"`
		IssueRoleList []string `json:"issue_role_list"`
	}{
		UserList:      userList,
		IssueRoleList: issueRoleList,
	}
	r, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	var limits []model.StepLimiterModel
	if err = global.GVA_DB.Where("tmpl_id=? and step_id=? and mode=?", req.TmplId, req.StepId, req.Mode).Find(&limits).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("参数错误")
	}
	if len(limits) > 0 {
		return nil, errors.New("此类型的限制策略，不能重复添加")
	}

	if err = global.GVA_DB.Create(&model.StepLimiterModel{
		WsId:      req.WsId,
		TmplId:    req.TmplId,
		StepId:    req.StepId,
		Mode:      req.Mode,
		Name:      req.Name,
		Rule:      string(r),
		CreatedAt: common.GetCurrentTime(),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StepLimiterSrv) Update(req types.StepLimiterUpdateReq) (resp interface{}, err error) {
	var userList []int
	var issueRoleList []string
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
	var rule = struct {
		UserList      []int    `json:"user_list"`
		IssueRoleList []string `json:"issue_role_list"`
	}{
		UserList:      userList,
		IssueRoleList: issueRoleList,
	}
	r, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Updates(&model.StepLimiterModel{
		Name:      req.Name,
		Rule:      string(r),
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *StepLimiterSrv) Delete(req types.StepLimiterDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Delete(&model.StepLimiterModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
