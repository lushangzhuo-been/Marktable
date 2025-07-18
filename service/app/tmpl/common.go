package tmpl

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mark3/global"
	"mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
)

func GetUserList(userIds []int) []user.UserModel {
	if len(userIds) == 0 {
		return nil
	}
	var users []user.UserModel
	global.GVA_DB.Where("id in ?", userIds).Find(&users)
	return users
}

func GetStatusInfo(id int) tmpl.StatusModel {
	var status tmpl.StatusModel
	global.GVA_DB.Where("id=?", id).Find(&status)
	return status
}

func GetDocument(wsId int, tmplId int, issueId string) (resp bson.M, err error) {
	collection := global.GVA_MONGO.Database("mark3").Collection("issue")
	objectID, err := primitive.ObjectIDFromHex(issueId)
	if err != nil {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"ws_id": wsId},
			{"tmpl_id": tmplId},
			{"_id": objectID},
		},
	}

	var document bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&document)
	if err != nil {
		return nil, errors.New("数据不存在")
	}
	document["issue_url"] = fmt.Sprintf(global.GVA_CONFIG.URL.Issue, wsId, tmplId, issueId)
	return document, nil
}
