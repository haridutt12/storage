package storage

type Storage interface {
	Put(key string, value interface{})
	Get(key string) (interface{}, error)
	Delete(key string)
}
