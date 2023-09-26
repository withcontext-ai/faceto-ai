package azeventhubs

import (
	"context"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/checkpoints"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/event"
	"faceto-ai/internal/pkg/middleware"
)

var (
	_ event.Sender   = (*AZProducer)(nil)
	_ event.Receiver = (*AZConsumer)(nil)
)

type Option func(*options)

type options struct {
	hubName string
}

func WithHubName(name string) Option {
	return func(o *options) {
		o.hubName = name
	}
}

type AZProducer struct {
	conf   *conf.Event
	log    *log.Helper
	client *azeventhubs.ProducerClient
}

func NewAZProducer(conf *conf.Event, logger log.Logger, opts ...Option) *AZProducer {
	o := options{
		hubName: conf.AzureHubs.DefaultHubName,
	}

	for _, opt := range opts {
		opt(&o)
	}

	l := log.NewHelper(logger)
	client, err := azeventhubs.NewProducerClientFromConnectionString(conf.AzureHubs.ConnectionString, o.hubName, nil)
	if err != nil {
		l.Fatal(fmt.Sprintf("NewAZProducer: %s %v", o.hubName, err))
	}
	return &AZProducer{
		conf:   conf,
		log:    l,
		client: client,
	}
}

func (p *AZProducer) Send(ctx context.Context, msg event.Event) error {
	p.log.WithContext(ctx).Debugf("Send: %s", string(msg.Value()))

	batch, err := p.client.NewEventDataBatch(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "NewEventDataBatch")
	}

	key := msg.Key()
	err = batch.AddEventData(&azeventhubs.EventData{
		MessageID: &key,
		Body:      msg.Value(),
	}, nil)
	if err != nil {
		return errors.Wrap(err, "AddEventData")
	}

	err = p.client.SendEventDataBatch(ctx, batch, nil)
	if err != nil {
		return errors.Wrap(err, "SendEventDataBatch")
	}

	return nil
}

func (p *AZProducer) SendBatch(ctx context.Context, msgs []event.Event) error {
	p.log.WithContext(ctx).Debugf("SendBatch: %v", msgs)

	batch, err := p.client.NewEventDataBatch(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "NewEventDataBatch")
	}

	for _, msg := range msgs {
		key := msg.Key()
		err = batch.AddEventData(&azeventhubs.EventData{
			MessageID: &key,
			Body:      msg.Value(),
		}, nil)
		if err != nil {
			return errors.Wrap(err, "AddEventData")
		}
	}

	err = p.client.SendEventDataBatch(ctx, batch, nil)
	if err != nil {
		return errors.Wrap(err, "SendEventDataBatch")
	}

	return nil
}

func (p *AZProducer) Close(ctx context.Context) error {
	return p.client.Close(ctx)
}

type AZConsumer struct {
	conf   *conf.Event
	log    *log.Helper
	client *azeventhubs.ConsumerClient
}

func NewAZConsumer(conf *conf.Event, logger log.Logger, opts ...Option) *AZConsumer {
	o := options{
		hubName: conf.AzureHubs.IndexHubName,
	}

	for _, opt := range opts {
		opt(&o)
	}

	l := log.NewHelper(logger)
	client, err := azeventhubs.NewConsumerClientFromConnectionString(
		conf.AzureHubs.ConnectionString,
		o.hubName,
		azeventhubs.DefaultConsumerGroup,
		nil,
	)

	if err != nil {
		l.Fatal(fmt.Sprintf("NewAZConsumer: %s %v", o.hubName, err))
	}
	return &AZConsumer{
		conf:   conf,
		log:    l,
		client: client,
	}
}

func (s *AZConsumer) Receive(ctx context.Context, handler event.Handler) error {
	containerClient, err := container.NewClientFromConnectionString(s.conf.AzureHubs.BlobConnectionString, s.conf.AzureHubs.BlobContainerName, nil)
	if err != nil {
		return errors.Wrap(err, "NewClientFromConnectionString")
	}

	checkpointStore, err := checkpoints.NewBlobStore(containerClient, nil)
	if err != nil {
		return errors.Wrap(err, "NewBlobStore")
	}

	processor, err := azeventhubs.NewProcessor(s.client, checkpointStore, &azeventhubs.ProcessorOptions{})

	dispatchPartitionClients := func() {
		for {
			partitionClient := processor.NextPartitionClient(ctx)

			if partitionClient == nil {
				break
			}

			go func() {
				s.log.WithContext(ctx).Debugf("Starting partition client %s", partitionClient.PartitionID())
				if err := s.processEvents(ctx, partitionClient, handler); err != nil {
					s.log.WithContext(ctx).Errorf("processEvents: %v", err)
					return
				}
				s.log.WithContext(ctx).Debugf("Stopped partition client %s", partitionClient.PartitionID())
			}()
		}
	}

	go dispatchPartitionClients()

	processorCtx, processorCancel := context.WithCancel(ctx)
	defer processorCancel()

	if err := processor.Run(processorCtx); err != nil {
		return errors.Wrap(err, "processor.Run")
	}

	return nil
}

func (s *AZConsumer) Close(ctx context.Context) error {
	return s.client.Close(ctx)
}

func (s *AZConsumer) processEvents(ctx context.Context, partitionClient *azeventhubs.ProcessorPartitionClient, handler event.Handler) error {
	defer partitionClient.Close(ctx)

	for {
		childCtx := context.WithValue(ctx, middleware.SpanID, helper.Generator())
		s.log.WithContext(childCtx).Debugf("Waiting for events on partition %s", partitionClient.PartitionID())

		receiveCtx, receiveCtxCancel := context.WithTimeout(childCtx, time.Minute)
		events, err := partitionClient.ReceiveEvents(receiveCtx, 1, nil)
		receiveCtxCancel()
		if err != nil && !errors.Is(err, context.DeadlineExceeded) {
			return errors.Wrapf(err, "ReceiveEvents for partition %s", partitionClient.PartitionID())
		}

		if len(events) == 0 {
			continue
		}

		s.log.WithContext(childCtx).Debugf("Processing %s event %d", partitionClient.PartitionID(), len(events))

		for _, value := range events {
			s.log.WithContext(childCtx).Debugf("Processing event %s %s", *value.MessageID, string(value.Body))

			message := event.NewMessage(*value.MessageID, value.Body)

			err = handler(childCtx, message)
			if err != nil {
				return errors.Wrapf(err, "handler event %s %s", *value.MessageID, string(value.Body))
			}
		}

		if len(events) != 0 {
			if err := partitionClient.UpdateCheckpoint(childCtx, events[len(events)-1]); err != nil {
				return errors.Wrap(err, "UpdateCheckpoint")
			}
		}
	}
}
