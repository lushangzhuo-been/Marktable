package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mark3/global"
	"mark3/pkg/common"
	"mark3/repository/db/model/app/tmpl"
	model "mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/repository/db/model/ws"
	types "mark3/types/app"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type CommonSrv struct{}

func (s *CommonSrv) LogUserUserRecentlyVisitedTmpl(userid int, wsId int, tmplId int) {
	var log tmpl.UserRecentlyVisitedLogModel
	err := global.GVA_DB.Where("userid=? and ws_id=? and tmpl_id=?", userid, wsId, tmplId).First(&log).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.WsId = wsId
			log.TmplId = tmplId
			log.Userid = userid
			log.VisitedTime = common.GetCurrentTime()
			if err = global.GVA_DB.Create(&log).Error; err != nil {
				global.GVA_LOG.Error(err.Error())
				return
			}
			return
		}
		global.GVA_LOG.Error(err.Error())
		return
	}

	log.VisitedTime = common.GetCurrentTime()
	if err = global.GVA_DB.Where("userid=? and ws_id=? and tmpl_id=?", userid, wsId, tmplId).Updates(&log).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

}

func (s *CommonSrv) GetUserUserRecentlyVisitedLog(userid int) (resp interface{}, err error) {
	var logs []tmpl.UserRecentlyVisitedLogModel
	if err = global.GVA_DB.Where("userid=? order by visited_time desc", userid).Find(&logs).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	var (
		tmplIds []int
		wsIds   []int

		tmplList []model.TmplModel
		wsList   []ws.WsModel

		tmplListMap = make(map[int]model.TmplModel)
		wsListMap   = make(map[int]ws.WsModel)
	)
	for _, log := range logs {
		tmplIds = append(tmplIds, log.TmplId)
		wsIds = append(wsIds, log.WsId)
	}

	global.GVA_DB.Where("id in ?", tmplIds).Find(&tmplList)
	for _, tmpl := range tmplList {
		tmplListMap[tmpl.Id] = tmpl
	}
	global.GVA_DB.Where("id in ?", wsIds).Find(&wsList)
	for _, ws := range wsList {
		wsListMap[ws.Id] = ws
	}

	type logVo struct {
		WsId     int    `json:"ws_id"`
		TmplId   int    `json:"tmpl_id"`
		WsName   string `json:"ws_name"`
		TmplName string `json:"tmpl_name"`
		TmplMode string `json:"tmpl_mode"`
	}

	var logVoList []logVo
	for _, log := range logs {
		ws, ok := wsListMap[log.WsId]
		if !ok {
			continue
		}
		tmpl, ok := tmplListMap[log.TmplId]
		if !ok {
			continue
		}
		logVoList = append(
			logVoList,
			logVo{
				WsId:     log.WsId,
				TmplId:   log.TmplId,
				WsName:   ws.Name,
				TmplName: tmpl.Name,
				TmplMode: tmpl.Mode,
			})
	}
	return logVoList, nil
}

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Avatar   string `json:"avatar"`
}

func (s *CommonSrv) GetWsUserList(req types.CommonGetWsUserListReq, userid int) (resp interface{}, err error) {
	var commonlyUsedList []UserInfo
	if req.Key != "" {
		req.Key = strings.TrimSpace(req.Key)
	}

	if req.Key == "" {
		commonlyUsedList, _ = s.GetUserHandleCommonlyUsed(req.WsId, userid)
		if commonlyUsedList == nil {
			// 不做处理
		} else if len(commonlyUsedList) >= 10 {
			return commonlyUsedList, nil
		}
	}

	db := global.GVA_DB

	db = db.Table("ws_member as a").Select("b.id, b.username, b.full_name, b.avatar").Joins("inner join user as b on a.userid=b.id")
	db = db.Where("ws_id", req.WsId)

	if req.Key != "" {
		db = db.Where("(b.username like ? or b.full_name like ?)", "%"+req.Key+"%", "%"+req.Key+"%")
	}

	if req.Ex != "" {
		ex := strings.Split(req.Ex, ",")
		var useridList []int
		for _, s := range ex {
			userid, err := strconv.Atoi(s)
			if err == nil {
				useridList = append(useridList, userid)
			}
		}
		if len(useridList) > 0 {
			db = db.Where("b.id not in ?", useridList)
		}
	}

	if commonlyUsedList != nil && len(commonlyUsedList) > 0 {
		var commonlyUseridList []int
		for _, commonlyUsed := range commonlyUsedList {
			commonlyUseridList = append(commonlyUseridList, commonlyUsed.Id)
		}
		if len(commonlyUseridList) > 0 {
			db = db.Where("b.id not in ?", commonlyUseridList)
		}
	}

	db = db.Limit(10)

	var users []UserInfo
	db.Scan(&users)
	var commonlyUsedListLen int = 0
	if commonlyUsedList != nil {
		commonlyUsedListLen = len(commonlyUsedList)
	}
	needUsersLen := 10 - commonlyUsedListLen
	if len(users) < needUsersLen {
		needUsersLen = len(users)
	}
	commonlyUsedList = append(commonlyUsedList, users[0:needUsersLen]...)
	return commonlyUsedList, nil
}

func (s *CommonSrv) GetUserHandleCommonlyUsed(wsId int, userid int) (userVoListResp []UserInfo, err error) {
	var commonlyUsed tmpl.UserHandleCommonlyUsedModel
	if err = global.GVA_DB.Where("ws_id=? and userid=? order by updated_at desc", wsId, userid).Find(&commonlyUsed).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	if commonlyUsed.Id == 0 {
		return nil, err
	}

	var userIds []int
	commonlyUsedUseridList := strings.Split(commonlyUsed.CommonlyUsedUseridList, ",")
	var commonlyUsedUseridListInt32 []int
	for i := 0; i < len(commonlyUsedUseridList); i++ {
		var commonlyUsedUserid, _ = strconv.Atoi(commonlyUsedUseridList[i])
		commonlyUsedUseridListInt32 = append(commonlyUsedUseridListInt32, commonlyUsedUserid)
	}
	userIds = append(userIds, commonlyUsedUseridListInt32...)

	var users = GetUsersByUserIds(userIds)

	var userVoList []UserInfo
	var userVoMap = make(map[int]UserInfo)
	for _, userObj := range users {
		var userVo UserInfo
		userVo.Id = userObj.Id
		userVo.Username = userObj.Username
		userVo.FullName = userObj.FullName
		userVo.Avatar = userObj.Avatar
		userVoMap[userObj.Id] = userVo
	}

	for i := len(commonlyUsedUseridListInt32) - 1; i > 0; i-- {
		userVoList = append(userVoList, userVoMap[commonlyUsedUseridListInt32[i]])
	}

	return userVoList, nil
}

// todo
func (s *CommonSrv) LogUserOfCloseUser(wsId int, closeUseridList []int, userid int) {
	if len(closeUseridList) == 0 {
		return
	}
	var err error
	var oldCommonlyUsed tmpl.UserHandleCommonlyUsedModel
	if err = global.GVA_DB.Where("ws_id=? and userid = ? order by updated_at desc", wsId, userid).Find(&oldCommonlyUsed).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	if oldCommonlyUsed.Id != 0 {
		var oldCommonlyUsedUseridList = strings.Split(oldCommonlyUsed.CommonlyUsedUseridList, ",")
		for i := 0; i < len(closeUseridList); i++ {
			if !strings.Contains(oldCommonlyUsed.CommonlyUsedUseridList, strconv.Itoa(closeUseridList[i])) {
				if len(oldCommonlyUsedUseridList) < 10 {
					oldCommonlyUsedUseridList = append(oldCommonlyUsedUseridList, strconv.Itoa(closeUseridList[i]))
				} else {
					oldCommonlyUsedUseridList = append(oldCommonlyUsedUseridList[1:], strconv.Itoa(closeUseridList[i]))
				}
			} else {
				// 如果之前有，将下标移动到最后
				var index int
				var removeIndex int
				for index = 0; index < len(oldCommonlyUsedUseridList); index++ {
					if oldCommonlyUsedUseridList[index] == strconv.Itoa(closeUseridList[i]) {
						removeIndex = index
						break
					}
				}
				oldCommonlyUsedUseridList = append(oldCommonlyUsedUseridList[:removeIndex], oldCommonlyUsedUseridList[removeIndex+1:]...)
				oldCommonlyUsedUseridList = append(oldCommonlyUsedUseridList, strconv.Itoa(closeUseridList[i]))
			}
		}
		oldCommonlyUsed.CommonlyUsedUseridList = strings.Join(oldCommonlyUsedUseridList, ",")
		if err = global.GVA_DB.Save(&oldCommonlyUsed).Error; err != nil {
			global.GVA_DB.Rollback()
			global.GVA_LOG.Error(err.Error())
		}
		return
	}
	var commonlyUsed tmpl.UserHandleCommonlyUsedModel
	var commonlyUsedUseridList []string
	min_len := min(len(closeUseridList), 10)
	for i := 0; i < min_len; i++ {
		var commonlyUsedUserid = strconv.Itoa(closeUseridList[i])
		commonlyUsedUseridList = append(commonlyUsedUseridList, commonlyUsedUserid)
	}
	var commonlyUsedUseridListStr string
	commonlyUsedUseridListStr = strings.Join(commonlyUsedUseridList, ",")

	commonlyUsed.WsId = wsId
	commonlyUsed.CommonlyUsedUseridList = commonlyUsedUseridListStr
	commonlyUsed.Userid = userid
	commonlyUsed.CreatedAt = common.GetCurrentTime()
	commonlyUsed.UpdatedAt = common.GetCurrentTime()

	if err = global.GVA_DB.Create(&commonlyUsed).Error; err != nil {
		global.GVA_DB.Rollback()
		global.GVA_LOG.Error(err.Error())
	}
}

func GetUsersByUserIdStr(userIdStr string) []user.UserModel {
	if len(userIdStr) == 0 {
		return nil
	}
	var userIds []int
	userLists := strings.Split(userIdStr, ",")
	for _, s := range userLists {
		uid, err := strconv.Atoi(s)
		if err == nil {
			userIds = append(userIds, uid)
		}
	}
	users := GetUsersByUserIds(userIds)
	return users
}

func GetUsersByUserIds(userIds []int) []user.UserModel {
	if len(userIds) == 0 {
		return nil
	}
	var users []user.UserModel
	global.GVA_DB.Where("id in ?", userIds).Find(&users)
	return users
}

func GetDocumentsByIdStr(wsId int, tmplId int, issueIdsStr string) []bson.M {
	if len(issueIdsStr) == 0 {
		return nil
	}
	var objectIds []primitive.ObjectID
	issueIds := strings.Split(issueIdsStr, ",")
	for _, s := range issueIds {
		objectId, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			continue
		}
		objectIds = append(objectIds, objectId)
	}
	filter := bson.M{
		"ws_id":   wsId,
		"tmpl_id": tmplId,
		"_id":     bson.M{"$in": objectIds},
	}
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil
	}
	return documents
}

func GetStatusInfo(id int) model.StatusModel {
	var status model.StatusModel
	global.GVA_DB.Where("id=?", id).Find(&status)
	return status
}

func GetStatusByIds(statusIds []int) []model.StatusModel {
	if len(statusIds) == 0 {
		return nil
	}
	var status []model.StatusModel
	global.GVA_DB.Where("id in ?", statusIds).Find(&status)
	return status
}

func GetStatusListInfo(statusIdStr string) []model.StatusModel {
	if len(statusIdStr) == 0 {
		return nil
	}
	var statusIds []int
	statusLists := strings.Split(statusIdStr, ",")
	for _, s := range statusLists {
		uid, err := strconv.Atoi(s)
		if err == nil {
			statusIds = append(statusIds, uid)
		}
	}
	users := GetStatusByIds(statusIds)
	return users
}
