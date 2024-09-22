package kafka

import (
	"context"
	"fmt"

	"github.com/YohanADR/SpotHome/infrastructure/logger"
	"github.com/YohanADR/SpotHome/pkg/messaging"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
	Logger logger.Logger
}

// Vérification que KafkaProducer implémente bien l'interface MessageProducer
var _ messaging.MessageProducer = (*KafkaProducer)(nil)

// NewKafkaProducer initialise un nouveau producteur Kafka avec logger
func NewKafkaProducer(brokers []string, topic string, log logger.Logger) (*KafkaProducer, error) {
	if len(brokers) == 0 {
		log.Error("Erreur: aucun broker Kafka spécifié")
		return nil, fmt.Errorf("aucun broker Kafka spécifié")
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	log.Info("Producteur Kafka initialisé", "brokers", brokers, "topic", topic)

	return &KafkaProducer{
		Writer: writer,
		Logger: log,
	}, nil
}

// Produce envoie un message à Kafka (implémente MessageProducer)
func (kp *KafkaProducer) Produce(ctx context.Context, topic string, message interface{}) error {
	err := kp.Writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte(message.(string)),
		},
	)
	if err != nil {
		kp.Logger.Error("Erreur lors de l'envoi du message Kafka", "error", err)
		return err
	}
	kp.Logger.Info("Message envoyé à Kafka avec succès", "topic", topic, "message", message)
	return nil
}

// Close ferme le producteur Kafka (implémente MessageProducer)
func (kp *KafkaProducer) Close() error {
	if err := kp.Writer.Close(); err != nil {
		kp.Logger.Error("Erreur lors de la fermeture du producteur Kafka", "error", err)
		return err
	}
	kp.Logger.Info("Producteur Kafka fermé avec succès")
	return nil
}
