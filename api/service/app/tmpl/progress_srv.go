package tmpl

import (
	"context"
	"errors"
	"mark3/global"
	"mark3/pkg/common"
	types "mark3/types/app/tmpl"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProgressSrv struct{}

func (s *ProgressSrv) getProgressCount(wsId int, tmplId int, issueId string) int64 {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue_progress")
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": wsId},
			{"tmpl_id": tmplId},
			{"issue_id": issueId},
		},
	}
	cnt, _ := collection.CountDocuments(context.TODO(), filter)
	return cnt
}

func (s *ProgressSrv) AddProgress(userid int, req types.AddProgress) (resp interface{}, err error) {
	document := bson.M{}
	document["issue_id"] = req.IssueId
	document["ws_id"] = req.WsId
	document["tmpl_id"] = req.TmplId
	document["content"] = req.Content
	document["creator"] = userid
	document["created_at"] = common.GetCurrentTime()
	document["updated_at"] = common.GetCurrentTime()
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue_progress")
	_, err = collection.InsertOne(context.TODO(), &document)
	if err != nil {
		return
	}
	go AtPersonNotice(userid, req.WsId, req.TmplId, document)
	return
}

func (s *ProgressSrv) UpdateProgress(userid int, req types.UpdateProgress) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue_progress")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
			{"_id": objectID},
		},
	}
	var progress bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		return nil, errors.New("数据不存在")
	}

	creator, ok := progress["creator"].(int32)
	if !ok {
		return nil, err
	}
	if int(creator) != userid {
		return nil, errors.New("无权限修改")
	}

	progress["content"] = req.Content
	update := bson.M{
		"$set": progress,
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	go AtPersonNotice(userid, req.WsId, req.TmplId, progress)
	return
}

func (s *ProgressSrv) DeleteProgress(userid int, req types.DeleteProgress) (resp interface{}, err error) {
	collection := global.GVA_MONGO.Database(global.GVA_CONFIG.Mongo.MongoDataBase).Collection("issue_progress")
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": req.WsId},
			{"tmpl_id": req.TmplId},
			{"issue_id": req.IssueId},
			{"_id": objectID},
		},
	}
	var progress bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		return nil, errors.New("数据不存在")
	}
	creator, ok := progress["creator"].(int32)
	if !ok {
		return nil, err
	}
	if int(creator) != userid {
		return nil, errors.New("无权限修改")
	}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return
}
