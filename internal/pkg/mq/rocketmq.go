package mq

import (
	"context"
	"fmt"
	"gin-design/internal/config"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5/credentials"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
)

var (
	prd rmq_client.Producer
	csm rmq_client.Consumer
)

const (
	Topic     = "test-topic"         // 主题名称
	Endpoint  = "192.168.0.100:8082" // grpc proxy address，换成宿主机的ip和上文安装环境的 proxy 的端口。
	AccessKey = "xxxxxx"             // 没有的话这么写就行
	SecretKey = "xxxxxx"
)

func Test() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()
	// In most case, you don't need to create many producers, singleton pattern is more recommended.
	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithTopics(Topic),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start producer
	err = producer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// graceful stop producer
	defer producer.GracefulStop()

	for i := 0; i < 10; i++ {
		// new a message
		msg := &rmq_client.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in sync
		resp, err := producer.Send(context.TODO(), msg)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}

func InitProducer(cfg *config.Config) rmq_client.Producer {
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()

	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithTopics(Topic),
	)
	if err != nil {
		log.Fatal(err)
	}

	// start producer
	err = producer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// graceful stop producer
	prd = producer
	return producer
}

func InitSimpleCosumer(cfg *config.Config) rmq_client.Consumer {
	// log to console
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()

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

	csm = simpleConsumer

	return simpleConsumer
}

func SendMsg(ctx context.Context, topic string, tag string, msg string) error {

	message := &rmq_client.Message{
		Topic: Topic,
		Body:  []byte(msg),
	}

	message.SetTag(tag)

	resp, err := prd.Send(ctx, message)
	if err != nil {
		// logger.Log().Error(err.Error())
		return err
	}

	fmt.Println(resp)

	return nil
}

func SendAsyncMsg(ctx context.Context, topic string, tag string, msg string) {
	message := &rmq_client.Message{
		Topic: Topic,
		Body:  []byte(msg),
	}
	message.SetTag(tag)
	prd.SendAsync(ctx, message, func(ctx context.Context, rec []*rmq_client.SendReceipt, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("send async success")
		}
	})
}

func SendDelayMsg(ctx context.Context, topic string, tag string, msg string, duration time.Time) error {
	message := &rmq_client.Message{
		Topic: Topic,
		Body:  []byte(msg),
	}
	message.SetTag(tag)

	message.SetDelayTimestamp(duration)

	resp, err := prd.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Println("send message success, msgId: ", resp)
	return nil
}

func ConsumeMessage(ctx context.Context, topic string, tag string) error {

	return nil
}
