package queue

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

// 发送任务
func (r *RabbitMQ) ListenProducer(producer Producer) {
	// 验证链接是否正常,否则重新链接
	if r.Channel == nil {
		r.MqConnect()
	}

	// 用于检查交换机是否存在,已经存在不需要重复声明
	err := r.Channel.ExchangeDeclarePassive(r.ExchangeName, r.ExchangeType, true, false, false, true, nil)
	if err != nil{
		// 注册交换机
		// name:交换机名称,kind:交换机类型,durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;
		// noWait:是否非阻塞, true为是,不等待RMQ返回信息;args:参数,传nil即可; internal:是否为内部
		err =  r.Channel.ExchangeDeclare(r.ExchangeName, r.ExchangeType, true, false, false, true, nil)
		if err != nil {
			fmt.Printf("MQ注册交换机失败:%s \n", err)
			return
		}
	}

	// 用于检查队列是否存在,已经存在不需要重复声明
	_, err = r.Channel.QueueDeclarePassive(r.QueueName, true,false,false,true,nil)
	if err != nil{
		// 队列不存在,声明队列
		// name:队列名称;durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;noWait:是否非阻塞,
		// true为是,不等待RMQ返回信息;args:参数,传nil即可;exclusive:是否设置排他
		_, err = r.Channel.QueueDeclare(r.QueueName, true, false, false, true, nil)
		if err != nil {
			fmt.Printf("MQ注册队列失败:%s \n", err)
			return
		}
	}
	// 队列绑定
	err = r.Channel.QueueBind(r.QueueName, r.RoutingKey, r.ExchangeName, true,nil)
	if err != nil {
		fmt.Printf("MQ绑定队列失败:%s \n", err)
		return
	}

	// 发送任务消息
	err =  r.Channel.Publish(r.ExchangeName, r.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        producer.MsgContent(),
	})
	if err != nil {
		fmt.Printf("MQ任务发送失败:%s \n", err)
		return
	}
}

// 监听接收者接收任务
func (r *RabbitMQ) ListenReceiver(receiver Receiver) {
	// 处理结束关闭链接
	defer r.MqClose()
	// 验证链接是否正常
	if r.Channel == nil {
		r.MqConnect()
	}
	// 用于检查队列是否存在,已经存在不需要重复声明
	_, err := r.Channel.QueueDeclarePassive(r.QueueName, true,false,false,true,nil)
	if err != nil{
		// 队列不存在,声明队列
		// name:队列名称;durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;noWait:是否非阻塞,
		// true为是,不等待RMQ返回信息;args:参数,传nil即可;exclusive:是否设置排他
		_, err = r.Channel.QueueDeclare(r.QueueName, true, false, false, true, nil)
		if err != nil {
			fmt.Printf("MQ注册队列失败:%s \n", err)
			return
		}
	}
	// 绑定任务
	err =  r.Channel.QueueBind(r.QueueName, r.RoutingKey, r.ExchangeName, true, nil)
	if err != nil {
		fmt.Printf("绑定队列失败:%s \n", err)
		return
	}
	// 获取消费通道,确保rabbitMQ一个一个发送消息
	err =  r.Channel.Qos(1, 0, true)
	msgList, err :=  r.Channel.Consume(r.QueueName, "", false, false, false, false, nil)
	if err != nil {
		fmt.Printf("获取消费通道异常:%s \n", err)
		return
	}

	for msg := range msgList {
		// 处理数据
		//err := receiver.Consumer(msg.Body)
		//if err!=nil {
		//	err = msg.Ack(true)
		//	if err != nil {
		//		fmt.Printf("确认消息未完成异常:%s \n", err)
		//		return
		//	}
		//}else {
		//	// 确认消息,必须为false
		//	err = msg.Ack(false)
		//	if err != nil {
		//		fmt.Printf("确认消息完成异常:%s \n", err)
		//		return
		//	}
		//	return
		//}
		_ = receiver.Consumer(msg.Body)
		err := msg.Ack(false)
		if err != nil {
			fmt.Printf("确认消息完成异常:%s \n", err)
			return
		}
		s := BytesToString(&(msg.Body))
		fmt.Printf("receve msg is :%s\n", *s)

	}
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}