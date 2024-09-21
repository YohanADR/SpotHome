package logger

import (
	"log"

	"go.uber.org/zap"
)

// Logger est une interface pour encapsuler les méthodes de logging
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Sync() error
}

// ZapLogger est l'implémentation de l'interface Logger avec zap
type ZapLogger struct {
	logger *zap.Logger
}

// NewLogger crée une nouvelle instance de ZapLogger
func NewLogger() Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("impossible de créer le logger: %v", err)
	}
	return &ZapLogger{logger: zapLogger}
}

// Info logue un message de niveau Info avec des champs structurés
func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Error logue un message de niveau Error avec des champs structurés
func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// Sync vide le buffer et écrit tous les logs restants
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// FieldString retourne un champ structuré de type chaîne de caractères
func FieldString(key, value string) zap.Field {
	return zap.String(key, value)
}

// FieldError retourne un champ structuré pour une erreur
func FieldError(err error) zap.Field {
	return zap.Error(err)
}
