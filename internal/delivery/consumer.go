package delivery

import (
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/repository"
)

type KfaConsumer struct {
	Cfg  *model.Config
	Repo *repository.KafkaRepo
}

func New(cfg *model.Config, repo *repository.KafkaRepo) (cns *KfaConsumer) {
	return &KfaConsumer{
		Cfg:  cfg,
		Repo: repo,
	}
}

func (k *KfaConsumer) InitConsumer() (err error) {
	return
}
