package repository

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"os"
	"sync"
	"time"
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

func (k *KafkaRepo) Close() {
	k.KafkaClient.Close()
}

func (k *KafkaRepo) InitTopic() (err error) {
	var (
		res *kmsg.CreateTopicsResponse
	)

	req := kmsg.NewPtrCreateTopicsRequest()
	topic := kmsg.NewCreateTopicsRequestTopic()
	topic.Topic = k.Cfg.TopicName
	topic.NumPartitions = int32(k.Cfg.PartitionCount)
	topic.ReplicationFactor = 1
	req.Topics = append(req.Topics, topic)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	res, err = req.RequestWith(ctx, k.KafkaClient)
	if err != nil {
		log.Errorf("[redpanda] error request new topic, %#v", err)
		return
	}

	if len(res.Topics) != 1 {
		log.Errorf("expected one topic in response, saw %d", len(res.Topics))
		return
	}
	t := res.Topics[0]

	if err = kerr.ErrorForCode(t.ErrorCode); err != nil {
		log.Errorf("topic creation failure: %#v", err)
		return
	}
	log.Infof("topic %s created successfully!", t.Topic)

	// Now we will issue a metadata request to see that topic.
	{
		var (
			res *kmsg.MetadataResponse
		)
		req := kmsg.NewPtrMetadataRequest()
		topic := kmsg.NewMetadataRequestTopic()
		topic.Topic = &k.Cfg.TopicName
		req.Topics = append(req.Topics, topic)

		res, err = req.RequestWith(ctx, k.KafkaClient)
		if err != nil {
			log.Errorf("failed request topic info: %#v", err)
			return
		}

		// Check response for Kafka error codes and print them.
		// Other requests might have top level error codes, which indicate completed but failed requests.
		for _, topic := range res.Topics {
			err := kerr.ErrorForCode(topic.ErrorCode)
			if err != nil {
				log.Errorf("topic %v response has errored: %v", topic.Topic, err.Error())
			}
		}

		log.Infof("received '%v' topics and '%v' brokers", len(res.Topics), len(res.Brokers))
	}
	return
}

func (k *KafkaRepo) PublishMsg(data model.ProductData) {
	defer k.Close()
	payload, err := json.Marshal(&data)
	if err != nil {
		log.Errorf("Error marshall file : %#v\n", err)
		return
	}
	record := &kgo.Record{
		Key:   []byte(data.UserID),
		Value: payload,
		Topic: k.Cfg.TopicName,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	k.KafkaClient.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}
	})
	wg.Wait()
}
