package azeventhubs

import (
	"github.com/go-kratos/kratos/v2/log"

	"faceto-ai/internal/conf"
)

type SnapshotProducer struct {
	*AZProducer
}

func NewSnapshotProducer(conf *conf.Event, logger log.Logger) *SnapshotProducer {
	producer := NewAZProducer(conf, logger, WithHubName(conf.AzureHubs.IndexHubName))
	return &SnapshotProducer{
		AZProducer: producer,
	}
}

type SnapshotConsumer struct {
	*AZConsumer
}

func NewSnapshotConsumer(conf *conf.Event, logger log.Logger) *SnapshotConsumer {
	consumer := NewAZConsumer(conf, logger, WithHubName(conf.AzureHubs.IndexHubName))
	return &SnapshotConsumer{
		AZConsumer: consumer,
	}
}
