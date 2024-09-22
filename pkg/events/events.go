package events

// Event représente un événement métier dans l'application
type Event struct {
	Name    string
	Payload interface{}
}
