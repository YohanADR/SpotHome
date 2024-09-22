package messaging

import "context"

// MessageProducer définit l'interface pour émettre des messages via un système de messagerie (Kafka, RabbitMQ, etc.)
type MessageProducer interface {
	Produce(ctx context.Context, topic string, message interface{}) error
	Close() error
}
