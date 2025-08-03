package message

import (
	"context"
	"errors"
	"mark3/global"
	enum "mark3/repository/db/enum/message"
	"mark3/service/app"
	types "mark3/types/message"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageSrv struct{}

func (s *MessageSrv) GetUnReadCount(userid int) (resp interface{}, err error) {
	filter := bson.M{"userid": userid, "read_sign": enum.ReadSignNo}
	collection := global.GVA_MONGO.Database("mark3").Collection("user_message")
	cnt, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return cnt, nil
}

func (s *MessageSrv) GetMessageList(userid int, req types.GetMessageListReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("user_message")
	var filter bson.M
	switch req.Mode {
	case enum.ModeReassign:
		filter = bson.M{"userid": userid, "mode": enum.ModeReassign}
	case enum.ModeAtPerson:
		filter = bson.M{"userid": userid, "mode": enum.ModeAtPerson}
	}

	var unReadFilter bson.M
	switch req.Mode {
	case enum.ModeReassign:
		unReadFilter = bson.M{"userid": userid, "mode": enum.ModeReassign, "read_sign": enum.ReadSignNo}
	case enum.ModeAtPerson:
		unReadFilter = bson.M{"userid": userid, "mode": enum.ModeAtPerson, "read_sign": enum.ReadSignNo}
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	pageNum := req.PageNum
	if pageNum <= 0 {
		pageNum = 1
	}
	skip := (pageNum - 1) * pageSize

	findOptions := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).
		SetLimit(int64(pageSize)).
		SetSkip(int64(skip))

	cnt, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	unReadCnt, err := collection.CountDocuments(context.TODO(), unReadFilter)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	var documents []bson.M
	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	for _, document := range documents {
		originUserid := int(document["origin_userid"].(int32))
		var userIds = []int{originUserid}
		users := app.GetUsersByUserIds(userIds)
		if len(users) >= 1 {
			document["origin_user"] = users[0]
		} else {
			document["origin_user"] = nil
		}
	}

	resp = types.GetMessageListResp{
		List:      documents,
		Cnt:       cnt,
		UnReadCnt: unReadCnt,
	}
	return
}

func (s *MessageSrv) ReadMessage(userid int, req types.ReadMessageReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("user_message")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"userid": userid},
			{"_id": objectID},
		},
	}
	var message bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&message)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	message["read_sign"] = "yes"
	update := bson.M{
		"$set": message,
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return
}

func (s *MessageSrv) ReadAllMessage(userid int, req types.ReadAllMessageReq) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("user_message")

	var unReadFilter bson.M
	switch req.Mode {
	case enum.ModeReassign:
		unReadFilter = bson.M{"userid": userid, "mode": enum.ModeReassign, "read_sign": enum.ReadSignNo}
	case enum.ModeAtPerson:
		unReadFilter = bson.M{"userid": userid, "mode": enum.ModeAtPerson, "read_sign": enum.ReadSignNo}
	}

	update := bson.M{
		"$set": bson.M{"read_sign": "yes"},
	}
	_, err = collection.UpdateMany(context.TODO(), unReadFilter, update)
	if err != nil {
		return nil, err
	}
	return

}
