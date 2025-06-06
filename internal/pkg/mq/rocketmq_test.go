package mq

import (
	"context"
	"gin-design/internal/config"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	Test()
}

func TestConsume(t *testing.T) {
	prd := InitProducer(&config.Config{})
	defer prd.GracefulStop()

	SendAsyncMsg(context.Background(), "test-topic", "test01", "test01")
	time.Sleep(time.Second * 10)

	err := SendDelayMsg(context.Background(), "test-topic", "test01", "test01", time.Now().Add(time.Second*10))
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}
