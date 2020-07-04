package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

var mqConn *amqp.Connection
var mqChan *amqp.Channel

type RabbitMQ struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	QueueName    string            // 队列名称
	RoutingKey   string            // key名称
	ExchangeName string           // 交换机名称
	ExchangeType string           // 交换机类型
}

// 定义队列交换机对象
type QueueExchange struct {
	QuName  string           // 队列名称
	RtKey   string           // key值
	ExName  string           // 交换机名称
	ExType  string           // 交换机类型
}

// 创建一个新的操作对象
func New(q *QueueExchange) *RabbitMQ {
	return &RabbitMQ{
		QueueName:q.QuName,
		RoutingKey:q.RtKey,
		ExchangeName: q.ExName,
		ExchangeType: q.ExType,
	}
}

// 链接rabbitMQ
func (r *RabbitMQ)MqConnect() {
	var err error
	//RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", "guest", "guest", "******", 5673)
	mqConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	r.Connection = mqConn   // 赋值给RabbitMQ对象
	if err != nil {
		fmt.Printf("MQ打开链接失败:%s \n", err)
	}
	mqChan, err = mqConn.Channel()
	r.Channel = mqChan // 赋值给RabbitMQ对象
	if err != nil {
		fmt.Printf("MQ打开管道失败:%s \n", err)
	}
}

// 关闭RabbitMQ连接
func (r *RabbitMQ)MqClose() {
	// 先关闭管道,再关闭链接
	err := r.Channel.Close()
	if err != nil {
		fmt.Printf("MQ管道关闭失败:%s \n", err)
	}
	err = r.Connection.Close()
	if err != nil {
		fmt.Printf("MQ链接关闭失败:%s \n", err)
	}
}