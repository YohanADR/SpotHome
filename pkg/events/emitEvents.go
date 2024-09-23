package events

import (
	"context"
	"fmt"

	"github.com/YohanADR/SpotHome/infrastructure/logger" // Import pour le logger
	"github.com/YohanADR/SpotHome/pkg/messaging"         // Import pour le producteur Kafka
)

var kafkaProducer messaging.MessageProducer
var log logger.Logger

// InitEventSystem initialise le système d'événements avec KafkaProducer et Logger
func InitEventSystem(producer messaging.MessageProducer, logger logger.Logger) {
	kafkaProducer = producer
	log = logger
}

// emitEventWithParams envoie un événement à Kafka avec un nom d'événement et un message.
func EmitEvent(event Event) {
	// Format du message de l'événement
	eventMessage := fmt.Sprintf("Event: %s, Payload: %v", event.Name, event.Payload)

	// Envoyer l'événement au producteur Kafka
	err := kafkaProducer.Produce(context.Background(), "events-topic", eventMessage)
	if err != nil {
		// Logger l'erreur en cas d'échec de l'envoi
		log.Error("Erreur lors de l'envoi de l'événement à Kafka", "event", event.Name, "error", err)
	} else {
		// Logger le succès de l'envoi
		log.Info("Événement envoyé à Kafka avec succès", "event", event.Name)
	}
}
