package mq

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5/credentials"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
)

const (
	ConsumerGroup = "test-group"
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func Consume() {
	// log to console
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()
	// In most case, you don't need to create many consumers, singleton pattern is more recommended.
	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(awaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			Topic: rmq_client.SUB_ALL,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start simpleConsumer
	err = simpleConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// graceful stop simpleConsumer
	defer simpleConsumer.GracefulStop()

	go func() {
		for {
			fmt.Println("start receive message")
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				fmt.Println(err)
			}
			// ack message
			for _, mv := range mvs {
				simpleConsumer.Ack(context.TODO(), mv)
				fmt.Println(string(mv.GetBody()))
			}
			fmt.Println("wait a moment")
			fmt.Println()
			time.Sleep(time.Second * 3)
		}
	}()
	// 阻塞主 goroutine 防止退出，观察子 goroutine 状态
	done := make(chan struct{})
	<-done
}
