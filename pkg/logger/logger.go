package logger

import "go.uber.org/zap"

// Logger est l'interface qui définit les méthodes de log génériques
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

// zapLogger est une implémentation de l'interface Logger avec Zap
type ZapLogger struct {
	zap *zap.Logger
}

// InitLogger initialise et retourne une instance de zapLogger
func InitLogger() Logger {
	// Configuration du logger Zap
	zapLogger, _ := zap.NewProduction()

	return &ZapLogger{
		zap: zapLogger,
	}
}

// Info implémente la méthode Info pour Zap
func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	l.zap.Sugar().Infow(msg, fields...)
}

// Error implémente la méthode Error pour Zap
func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	l.zap.Sugar().Errorw(msg, fields...)
}

// Fatal implémente la méthode Fatal pour Zap
func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	l.zap.Sugar().Fatalw(msg, fields...)
}

// ShutdownLogger arrête proprement le logger Zap
func ShutdownLogger(l Logger) {
	if zapL, ok := l.(*ZapLogger); ok {
		_ = zapL.zap.Sync() // Assurez-vous que tous les logs sont flushés avant l'arrêt
	}
}
