package ports

// RedisPort est une interface qui définit les opérations possibles avec Redis
type RedisPort interface {
	Close() error
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}
