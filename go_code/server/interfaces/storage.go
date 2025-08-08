package interfaces

import (
	"time"
)

// KeyValueStore defines the interface for key-value storage
type KeyValueStore interface {
	Store(key string, value interface{}, ttl *time.Duration) bool
	Retrieve(key string) interface{}
	Delete(key string) bool
	Exists(key string) bool
}
