package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitConsumer inicializa e retorna um consumidor Kafka.
func InitConsumer(brokers []string, groupID string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0 // ou a versão do Kafka que você está usando
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatal("Erro ao criar o grupo de consumidores:", err)
	}
	return consumerGroup
}

// ConsumeMessages consome mensagens de um tópico Kafka e insere no MongoDB.
func ConsumeMessages(consumer sarama.ConsumerGroup, client *mongo.Client, topic string) {
	ctx := context.Background()
	for {
		handler := exampleConsumerGroupHandler{client}
		if err := consumer.Consume(ctx, []string{topic}, &handler); err != nil {
			log.Panicf("Erro ao consumir mensagens: %v", err)
		}
	}
}

type exampleConsumerGroupHandler struct {
	mongoClient *mongo.Client
}

func (h *exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Mensagem recebida: chave=%s, valor=%s, tópico=%s, partition=%d, offset=%d\n",
			string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
		// Aqui você pode inserir lógica para inserir a mensagem no MongoDB
		sess.MarkMessage(msg, "")
	}
	return nil
}
