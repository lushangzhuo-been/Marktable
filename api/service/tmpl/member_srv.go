package tmpl

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/tmpl"
	model "mark3/repository/db/model/tmpl"
	commonTypes "mark3/types/common"
	types "mark3/types/tmpl"
	"strconv"
	"strings"
)

type MemberSrv struct{}

func (s *MemberSrv) Config() (resp interface{}, err error) {
	var config types.MemberConfigResp
	config.StatusConfig = common.FormatConfig(enum.TmplStatusMap)
	return config, nil
}

func (s *MemberSrv) List(req types.MemberListReq) (resp interface{}, err error) {
	var sql = "select count(a.id) from tmpl_member as a left join user as b on a.userid=b.id where a.tmpl_id=?"
	var cnt int64
	if req.Key != "" {
		sql += " and (b.username like ? or b.full_name like ?)"
		global.GVA_DB.Raw(sql, req.TmplId, "%"+req.Key+"%", "%"+req.Key+"%").Scan(&cnt)
	} else {
		global.GVA_DB.Raw(sql, req.TmplId).Scan(&cnt)
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}

	sql = "select a.id, a.role_id, a.status, a.userid, b.username, b.full_name, b.email, b.phone, b.avatar from tmpl_member as a inner join user as b on a.userid=b.id where a.tmpl_id=?"
	var members []*types.MemberVo
	if req.Key != "" {
		sql += " and (b.username like ? or b.full_name like ?)"
		global.GVA_DB.Raw(sql+" limit ?,?", req.TmplId, "%"+req.Key+"%", "%"+req.Key+"%", (pageNum-1)*pageSize, pageSize).Scan(&members)
	} else {
		global.GVA_DB.Raw(sql+" limit ?,?", req.TmplId, (pageNum-1)*pageSize, pageSize).Scan(&members)
	}

	for _, member := range members {
		var role model.RoleModel
		global.GVA_DB.Where("id=?", member.RoleId).First(&role)
		member.RoleName = role.Name
	}
	resp = commonTypes.BasePageResp{
		List: members,
		Cnt:  cnt,
	}
	return
}

func (s *MemberSrv) Create(req types.MemberCreateReq) (resp interface{}, err error) {
	all := strings.Split(req.UserIdList, ",")
	var members []model.MemberModel
	global.GVA_DB.Where("tmpl_id=? and userid in ?", req.TmplId, all).Find(&members)
	var added []string
	for _, member := range members {
		added = append(added, strconv.Itoa(member.Userid))
	}

	var addedMap = make(map[string]string, len(added))
	for _, v := range added {
		addedMap[v] = v
	}
	var notAdded []string
	for _, v := range all {
		if _, ok := addedMap[v]; !ok {
			notAdded = append(notAdded, v)
		}
	}

	//批量更新
	for _, userid := range added {
		global.GVA_DB.Model(&model.MemberModel{}).Where("tmpl_id=? and userid=?", req.TmplId, userid).Updates(map[string]interface{}{"role": req.RoleId, "status": req.Status})
	}
	//批量添加
	for _, userid := range notAdded {
		var member model.MemberModel
		useridInt, _ := strconv.Atoi(userid)
		member.WsId = req.WsId
		member.TmplId = req.TmplId
		member.Userid = useridInt
		member.RoleId = req.RoleId
		member.Status = req.Status
		member.CreatedAt = common.GetCurrentTime()
		member.UpdatedAt = common.GetCurrentTime()
		global.GVA_DB.Model(&model.MemberModel{}).Create(&member)
	}
	return
}

func (s *MemberSrv) Update(req types.MemberUpdateReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Updates(&model.MemberModel{
		RoleId:    req.RoleId,
		Status:    req.Status,
		UpdatedAt: common.GetCurrentTime(),
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *MemberSrv) Delete(req types.MemberDeleteReq) (resp interface{}, err error) {
	if err = global.GVA_DB.Where("id=? and tmpl_id=?", req.Id, req.TmplId).Delete(&model.MemberModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
