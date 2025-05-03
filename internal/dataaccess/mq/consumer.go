package mq

import (
	"context"
	"github.com/IBM/sarama"
	"log"
)

type ConsumerHanlerFunc func(ctx context.Context, queueName string, payload []byte) error

// type Consumer interface {
// 	RegisterHandler(queueName string, handlerFunc ConsumerHanlerFunc)
// 	Start(ctx context.Context) error
// }

type Consumer struct {}

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Consumed message: %s", string(message.Value))

		// var task model.DownloadTask
		// if err := json.Unmarshal(message.Value, &task); err != nil {
		// 	log.Printf("Unmarshal error: %v", err)
		// 	continue
		// }

		// // Simulate processing the download task
		// log.Printf("Processing task %d: URL = %s", task.ID, task.URL)

		// // (Optional) update DB status to in_progress or completed
		// err := consumer.dbRepo.UpdateStatus(task.ID, "in_progress") // define this method
		// if err != nil {
		// 	log.Printf("Failed to update status: %v", err)
		// 	continue
		// }

		// // Simulate download...
		// log.Printf("Finished downloading task %d", task.ID)
		// _ = consumer.dbRepo.UpdateStatus(task.ID, "completed")

		session.MarkMessage(message, "")

		// Call download processor here
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

func StartKafkaConsumer() error {
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
		err := consumerGroup.Consume(ctx, []string{topic}, &Consumer{})
		if err != nil {
			log.Printf("Error consuming: %v", err)
		}
	}	
}

// build separate handler for each type of event (in handler folder)
