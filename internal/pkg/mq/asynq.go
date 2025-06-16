package mq

import (
	"encoding/json"
	"gin-design/internal/config"
	"time"

	"github.com/hibiken/asynq"
)

const (
	AdsConclusion = "conclusion"
	AdsAttribute  = "attribute"
	AdsConversion = "conversion"
)

var HashQueueMap = map[int]string{
	0: "q1",
	1: "q2",
	2: "q3",
	3: "q4",
	4: "q5",
	5: "q6",
	6: "q7",
	7: "q8",
}

type Asynq struct {
	client *asynq.Client
}

func NewAsynqClient(cfg *config.Config) *Asynq {
	opt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	client := asynq.NewClient(opt)

	return &Asynq{
		client: client,
	}
}

func (a *Asynq) SendMessage(name string, payload any, opts ...asynq.Option) error {
	py, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(name, py)

	_, err = a.client.Enqueue(task, opts...)

	return err
}

func (a *Asynq) SendDelayMessage(name string, payload any, delay int64) error {
	py, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(name, py, asynq.ProcessIn(time.Duration(delay)*time.Second))

	_, err = a.client.Enqueue(task)

	return err
}
