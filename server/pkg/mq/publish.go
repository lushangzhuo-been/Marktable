package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"mark3/global"
)

func PublishSimple(queueName string, message string) {
	//声明队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := global.GVA_MQ.QueueDeclare(
		// queue:队列名称
		queueName,
		//durable:是否持久化，当mq重启之后，还在
		false,
		//exclusive:是否独占即只能有一个消费者监听这个队列
		false,
		//autoDelete:是否自动删除。当没有Consumer时，自动删除掉
		false,
		//noWait:是否阻塞处理。true:不阻塞，false:阻塞
		false,
		//arguments:其他属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//调用channel 发送消息到队列中
	global.GVA_MQ.Publish(
		"",
		q.Name,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
