package mq

import (
	"github.com/IBM/sarama"
	"github.com/aq2208/goload/internal/model"
)

type Producer interface {
	SendMessage(payload []byte) error
}

type DownloadTaskEvent struct {
	DownloadTask model.DownloadTask
}

type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func newSaramaConfig() *sarama.Config {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Retry.Max = 1
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.ClientID = ""
	saramaConfig.Metadata.Full = true
	return saramaConfig
}

func NewKafkaProducer(brokers []string, topic string) (Producer, error) {
	producer, err := sarama.NewSyncProducer(brokers, newSaramaConfig())
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: producer, topic: topic}, nil
}

func (p *KafkaProducer) SendMessage(payload []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(payload),
	}
	_, _, err := p.producer.SendMessage(msg)
	return err
}