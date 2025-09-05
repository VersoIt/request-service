package kafka

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

func (k *Kafka) SendMessage(
	ctx context.Context,
	topic string,
	key []byte,
	value []byte,
) error {
	record := &kgo.Record{
		Value: value,
		Key:   key,
		Topic: topic,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	var err error
	k.client.Produce(ctx, record, func(record *kgo.Record, err error) {
		defer wg.Done()

		if err != nil {
			log.Errorf("Failed to send message to Kafka for topic '%s': %v", topic, err)
			return
		}

		log.Info("Message successfully sent to Kafka for topic '%s'", topic)
	})

	wg.Wait()

	return err
}
