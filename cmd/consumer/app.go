package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/delivery"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/repository"
	"os"
	"sync"
)

func main() {
	//log setup
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	//load config
	cfg, err := model.InitConfig()
	if err != nil {
		log.Fatalf("err load config, %#v", err)
	}

	kfka := repository.New(cfg)
	err = kfka.InitBrokers()
	if err != nil {
		log.Fatalf("err init broker, %#v", err)
	}

	consumer := delivery.New(cfg, kfka)
	err = consumer.InitConsumer()
	if err != nil {
		log.Fatalf("err InitConsumer, %#v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	log.Infof("consumer started")
	wg.Wait()
	os.Exit(0)
}
