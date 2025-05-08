package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/aq2208/goload/internal/service"
)

type ConsumerHanlerFunc func(ctx context.Context, queueName string, payload []byte) error

// type Consumer interface {
// 	RegisterHandler(queueName string, handlerFunc ConsumerHanlerFunc)
// 	Start(ctx context.Context) error
// }

type Consumer struct {
	service service.DownloadTaskService
}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Consumed message: %s", string(message.Value))
		session.MarkMessage(message, string(message.Value))

		var id uint64
		if err := json.Unmarshal(message.Value, &id); err != nil {
			log.Default().Printf("Failed to parse message: %v", err)
			return err
		}

		if err := c.service.ProcessDownload(context.Background(), id); err != nil {
			log.Default().Printf("Failed to consumer message: %v", err)
			return err
		}
	}
	return nil
}

func newSaramaConfigConsumer() *sarama.Config {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_1_0_0
	saramaConfig.ClientID = "goload-consumer"
	saramaConfig.Metadata.Full = true
	return saramaConfig
}

func StartKafkaConsumer(service service.DownloadTaskService) error {
	brokers := []string{"localhost:9094"}
	topic := "my-topic-1"
	groupId := "goload-consumer-group-1"
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, newSaramaConfigConsumer())
	if err != nil {
		log.Printf("Error starting Kafka consumer: %v", err)
		return err
	} 
	defer consumerGroup.Close()

	log.Println("Starting Kafka consumer...")

	ctx := context.Background()
	for {
		err := consumerGroup.Consume(ctx, []string{topic}, &Consumer{service: service})
		if err != nil {
			log.Printf("Error consuming: %v", err)
		}
	}
}

// build separate handler for each type of event (in handler folder)
