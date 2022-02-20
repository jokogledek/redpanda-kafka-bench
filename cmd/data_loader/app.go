package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/repository"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/usecase"
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

	parser := usecase.New(cfg, kfka)
	err = parser.LoadCSV()
	if err != nil {
		log.Fatalf("err loac csv, %#v", err)
	}
}
