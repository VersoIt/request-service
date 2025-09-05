package kafka

import (
	"RequestService/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Kafka struct {
	client *kgo.Client
}

func New(cfg config.Config) (*Kafka, error) {
	client, err := kgo.NewClient(kgo.SeedBrokers(cfg.KafkaProducer.Brokers...))
	if err != nil {
		return nil, err
	}

	return &Kafka{
		client: client,
	}, nil
}
