package azeventhubs

import (
	"github.com/go-kratos/kratos/v2/log"

	"faceto-ai/internal/conf"
)

type IndexProducer struct {
	*AZProducer
}

func NewIndexProducer(conf *conf.Event, logger log.Logger) *IndexProducer {
	producer := NewAZProducer(conf, logger, WithHubName(conf.AzureHubs.IndexHubName))
	return &IndexProducer{
		AZProducer: producer,
	}
}

type IndexConsumer struct {
	*AZConsumer
}

func NewIndexConsumer(conf *conf.Event, logger log.Logger) *IndexConsumer {
	consumer := NewAZConsumer(conf, logger, WithHubName(conf.AzureHubs.IndexHubName))
	return &IndexConsumer{
		AZConsumer: consumer,
	}
}
