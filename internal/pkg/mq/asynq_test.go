package mq

import (
	"gin-design/internal/config"
	"testing"
)

func TestAsynqSendMessage(t *testing.T) {
	as := NewAsynqClient(&config.Config{
		Redis: config.Redis{
			Addr:     "43.142.13.187:6379",
			Password: "980508",
			DB:       0,
		},
	})
	err := as.SendMessage("test-01", "test1")
	if err != nil {
		t.Error(err)
	}
}

func TestAsynqDelayMessage(t *testing.T) {
	as := NewAsynqClient(&config.Config{
		Redis: config.Redis{
			Addr:     "43.142.13.187:6379",
			Password: "980508",
			DB:       0,
		},
	})
	err := as.SendDelayMessage("test-01", "test2", 3600)
	if err != nil {
		t.Error(err)
	}
}
