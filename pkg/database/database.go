package database

// Database définit les méthodes pour interagir avec une base de données
type Database interface {
	Connect() error
	Query(query string, args ...interface{}) (interface{}, error)
	Close() error
}
