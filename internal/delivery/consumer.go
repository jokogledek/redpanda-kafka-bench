package delivery

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/repository"
	"github.com/ujunglangit-id/redpanda-kafka-bench/internal/usecase"
)

type KfaConsumer struct {
	Cfg            *model.Config
	Repo           *repository.KafkaRepo
	DataParserCase *usecase.DataParser
}

func New(cfg *model.Config, repo *repository.KafkaRepo) (cns *KfaConsumer) {
	return &KfaConsumer{
		Cfg:  cfg,
		Repo: repo,
	}
}

func (k *KfaConsumer) InitConsumer(parser *usecase.DataParser) (err error) {
	k.DataParserCase = parser
	for v := 0; v < k.Cfg.ConsumerCount; v++ {
		go func(index int) {
			k.startConsumer(index)
		}(v)
	}

	return
}

func (k *KfaConsumer) startConsumer(idx int) {
	ctx := context.Background()
	for {
		log.Infof("starting consumer ... %d", idx)
		fetches := k.Repo.KafkaClient.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}

		// or a callback function.
		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				k.DataParserCase.ParseMessage(record, idx)
			}
		})
	}
}
