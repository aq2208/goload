package producer

import (
	"log"

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
	saramaConfig.ClientID = "goload-producer"
	saramaConfig.Metadata.Full = true
	return saramaConfig
}

func NewKafkaProducer() (Producer, error) {
	brokers := []string{"localhost:9094"}
	topic := "my-topic-1"
	producer, err := sarama.NewSyncProducer(brokers, newSaramaConfig())
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
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