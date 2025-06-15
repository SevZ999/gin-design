package mq

import (
	"context"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type RocketMQ struct {
	prod rocketmq.Producer
}

func LoadRocketMQ() (*RocketMQ, func(), error) {
	//rocketmq.InitRocketMQ()
	prod, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.0.101:9876"}))
	if err != nil {
		return nil, nil, err
	}

	if err := prod.Start(); err != nil {
		return nil, nil, err
	}

	clean := func() {
		if err := prod.Shutdown(); err != nil {
			log.Fatalf("强制关闭: %v", err)
		}
	}

	return &RocketMQ{
		prod: prod,
	}, clean, nil
}

func (prod *RocketMQ) SendSync(topic string, tag string, msg string) error {
	mq := primitive.NewMessage(topic, []byte(msg))
	mq.WithTag(tag)

	result, err := prod.prod.SendSync(context.Background(), mq)
	if err != nil {
		return err
	}

	log.Printf("发送结果: %s", result.String())

	return nil
}
