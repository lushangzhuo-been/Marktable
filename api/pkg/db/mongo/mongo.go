package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mark3/global"
)

func InitMongo() *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(global.GVA_CONFIG.Mongo.MongoUrl)
	// 连接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	return client
}
