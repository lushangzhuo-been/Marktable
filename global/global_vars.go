package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"mark3/config"
)

var (
	GVA_DB     *gorm.DB
	GVA_MONGO  *mongo.Client
	GVA_LOG    *logrus.Logger
	GVA_CONFIG *config.Conf
	GVA_RDB    *redis.Client
	GVA_MQ     *amqp.Channel
)
