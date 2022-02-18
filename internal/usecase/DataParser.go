package usecase

import "github.com/ujunglangit-id/redpanda-kafka-bench/internal/model"

type DataParser struct {
	cfg *model.Config
}

func New(cfg *model.Config) *DataParser {
	return &DataParser{
		cfg: cfg,
	}
}

func ParseMessage() {

}
