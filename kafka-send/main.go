package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("this is 2 test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"172.20.90.161:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	for {
		time.Sleep(time.Second * 1)
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
