package ws

import (
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/ws"
	model "mark3/repository/db/model/ws"
	types "mark3/types/app"
	wsTypes "mark3/types/app/ws"
)

type WsSrv struct{}

func (s *WsSrv) GetMyWsList(userid int) (resp interface{}, err error) {
	var list []model.WsModel
	global.GVA_DB.Raw("select a.* from ws as a inner join ws_member as b on b.ws_id=a.id where b.status=? and b.userid=?", enum.WsStatusOk, userid).Scan(&list)
	return list, nil
}

func (s *WsSrv) GetWsInfo(userid int, req wsTypes.WsInfoReq) (resp interface{}, err error) {
	//获取空间详情
	var ws model.WsModel
	if err = global.GVA_DB.Where("id=?", req.WsId).First(&ws).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("数据不存在")
	}

	//获取空间成员
	var memberList []types.MemberVo
	sql := "select a.id, a.userid, a.role, a.status, b.username, b.full_name, b.email, b.phone, b.avatar from ws_member as a inner join user as b on a.userid=b.id where a.ws_id=? and a.status=?"
	global.GVA_DB.Raw(sql, req.WsId, enum.WsStatusOk).Scan(&memberList)

	//获取用户权限
	var isAdmin string = "no"
	var isMember string = "no"
	//获取空间管理员
	var adminList []types.MemberVo
	for _, member := range memberList {
		if member.Userid == userid {
			isMember = "yes"
		}
		if member.Role == enum.WsRoleAdmin {
			adminList = append(adminList, member)
			if member.Userid == userid {
				isAdmin = "yes"
			}
		}
	}

	//获取空间中流程数量
	var tmplCnt int64
	global.GVA_DB.Raw("select count(id) from tmpl where ws_id=?", req.WsId).Scan(&tmplCnt)

	//获取空间中成员数量
	var memberCnt int64
	global.GVA_DB.Raw("select count(id) from ws_member where ws_id=?", req.WsId).Scan(&memberCnt)

	resp = wsTypes.WsInfoResp{
		Ws:        ws,
		AdminList: adminList,
		TmplCnt:   tmplCnt,
		MemberCnt: memberCnt,
		IsAdmin:   isAdmin,
		IsMember:  isMember,
	}

	return
}
