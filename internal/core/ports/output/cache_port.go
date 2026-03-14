package output

type CachePort interface {
	Get(key string) (string, error)
	Set(key string, value string, ttlSeconds int) error
	Delete(key string) error
}
