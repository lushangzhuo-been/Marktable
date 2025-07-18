package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"mark3/global"
)

func InitRabbitMQSimple() *amqp.Channel {
	//连接mq
	host := global.GVA_CONFIG.Rabbit.Host
	port := global.GVA_CONFIG.Rabbit.Port
	username := global.GVA_CONFIG.Rabbit.Username
	password := global.GVA_CONFIG.Rabbit.Password
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("连接mq:%s", err)
	}
	//获取channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("获取channel:%s", err)
	}

	return channel
}
