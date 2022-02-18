package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"os"
)

type KafkaRepo struct {
	Cfg         *model.Config
	KafkaClient *kgo.Client
}

func New(cfg *model.Config) *KafkaRepo {
	return &KafkaRepo{
		Cfg: cfg,
	}
}

func (k *KafkaRepo) InitBrokers() (err error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(k.Cfg.Host...),
		kgo.ConsumerGroup(k.Cfg.ConsumerGroup),
		kgo.ConsumeTopics(k.Cfg.TopicName),
	}

	opts = append(opts, kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)))
	k.KafkaClient, err = kgo.NewClient(opts...)
	if err != nil {
		log.Errorf("[redpanda] error init new client, %#v", err)
	}
	return
}

func (k *KafkaRepo) CreateTopic(topic string, partition int) (err error) {
	return
}

func (k *KafkaRepo) PublishMsg() {

}
