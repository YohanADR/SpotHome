package errors

import (
	"fmt"
)

// AppError représente une erreur enrichie avec des informations supplémentaires
type AppError struct {
	Code    int    // Code de l'erreur (HTTP ou autre)
	Message string // Message utilisateur
	Details string // Détails supplémentaires (optionnels)
}

// New crée une nouvelle erreur enrichie
func New(code int, message, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Error satisfait l'interface error pour AppError
func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}

// Wrap permet d'envelopper une erreur existante avec un message personnalisé
func Wrap(err error, message string) *AppError {
	return &AppError{
		Code:    500, // Code HTTP générique pour une erreur interne
		Message: message,
		Details: err.Error(),
	}
}
