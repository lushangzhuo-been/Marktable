package ws

import (
	"mark3/global"
	model "mark3/repository/db/model/user"
	"mark3/types/common"
	types "mark3/types/ws"
	"strings"
)

type CommonSrv struct{}

func (s *CommonSrv) GetUserList(req types.GetUserListReq) (resp interface{}, err error) {
	db := global.GVA_DB.Model(&model.UserModel{})
	if req.Key != "" {
		req.Key = strings.TrimSpace(req.Key)
		db = db.Where("email = ? or full_name = ? or username = ?", req.Key, req.Key, req.Key)
	}
	var cnt int64
	db.Count(&cnt)

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	var users []model.UserModel
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users)

	resp = common.BasePageResp{
		List: users,
		Cnt:  cnt,
	}
	return
}
