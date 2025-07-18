package ws

import (
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	enum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/ws"
	commonTypes "mark3/types/common"
	types "mark3/types/ws"
	"strconv"
	"strings"
)

type MemberSrv struct{}

func (s *MemberSrv) Config() (resp interface{}, err error) {
	var config types.MemberConfigResp
	config.RoleConfig = common.FormatConfig(enum.WsRoleMap)
	config.StatusConfig = common.FormatConfig(enum.WsStatusMap)
	return config, nil
}

func (s *MemberSrv) List(req types.MemberListReq) (resp interface{}, err error) {
	var sql = "select count(a.id) from ws_member as a inner join user as b on a.userid=b.id where a.ws_id=?"
	var cnt int64
	if req.Key != "" {
		sql += " and (b.username like ? or b.full_name like ?)"
		global.GVA_DB.Raw(sql, req.WsId, "%"+req.Key+"%", "%"+req.Key+"%").Scan(&cnt)
	} else {
		global.GVA_DB.Raw(sql, req.WsId).Scan(&cnt)
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	sql = "select a.id, a.userid, a.role, a.status, b.username, b.full_name, b.email, b.phone, b.avatar from ws_member as a inner join user as b on a.userid=b.id where a.ws_id=?"
	var members []types.MemberVo
	if req.Key != "" {
		sql += " and (b.username like ? or b.full_name like ?)"
		global.GVA_DB.Raw(sql+" limit ?, ?", req.WsId, "%"+req.Key+"%", "%"+req.Key+"%", (pageNum-1)*pageSize, pageSize).Scan(&members)
	} else {
		global.GVA_DB.Raw(sql+" limit ?, ?", req.WsId, (pageNum-1)*pageSize, pageSize).Scan(&members)
	}

	resp = commonTypes.BasePageResp{
		List: members,
		Cnt:  cnt,
	}
	return
}

func (s *MemberSrv) Create(req types.MemberCreateReq) (resp interface{}, err error) {
	all := strings.Split(req.UseridList, ",")
	var members []model.MemberModel
	global.GVA_DB.Where("ws_id=? and userid in ?", req.WsId, all).Find(&members)
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
		global.GVA_DB.Model(&model.MemberModel{}).Where("ws_id=? and userid=?", req.WsId, userid).Updates(map[string]interface{}{"role": req.Role, "status": req.Status})
	}
	//批量添加
	for _, userid := range notAdded {
		var member model.MemberModel
		useridInt, _ := strconv.Atoi(userid)
		member.WsId = req.WsId
		member.Userid = useridInt
		member.Role = req.Role
		member.Status = req.Status
		member.CreatedAt = common.GetCurrentTime()
		member.UpdatedAt = common.GetCurrentTime()
		global.GVA_DB.Model(&model.MemberModel{}).Create(&member)
	}
	return
}

func (s *MemberSrv) Update(req types.MemberUpdateReq) (resp interface{}, err error) {
	var member model.MemberModel
	global.GVA_DB.Where("ws_id=? and id=?", req.WsId, req.Id).First(&member)
	if member.Id == 0 {
		return nil, errors.New("记录不存在，请刷新页面")
	}

	if member.Role == enum.WsRoleAdmin {
		//判断管理员必须存在一个
		var cnt int64
		global.GVA_DB.Model(&model.MemberModel{}).Where("ws_id=? and role=? and id !=?", req.WsId, enum.WsRoleAdmin, req.Id).Count(&cnt)
		if cnt < 1 {
			return nil, errors.New("操作失败，至少要存在一个管理员")
		}
	}

	if err = global.GVA_DB.Where("ws_id=? and id=?", req.WsId, req.Id).Updates(&model.MemberModel{
		Role:   req.Role,
		Status: req.Status,
	}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}

func (s *MemberSrv) Delete(req types.MemberDeleteReq) (resp interface{}, err error) {
	var member model.MemberModel
	global.GVA_DB.Where("ws_id=? and id=?", req.WsId, req.Id).First(&member)
	if member.Id == 0 {
		return nil, errors.New("记录不存在，请刷新页面")
	}

	if member.Role == enum.WsRoleAdmin {
		//判断管理员必须存在一个
		var cnt int64
		global.GVA_DB.Model(&model.MemberModel{}).Where("ws_id=? and role=? and id !=?", req.WsId, enum.WsRoleAdmin, req.Id).Count(&cnt)
		if cnt < 1 {
			return nil, errors.New("操作失败，至少要存在一个管理员")
		}
	}

	if err = global.GVA_DB.Where("ws_id=? and id=?", req.WsId, req.Id).Delete(&model.MemberModel{}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("操作失败")
	}
	return
}
