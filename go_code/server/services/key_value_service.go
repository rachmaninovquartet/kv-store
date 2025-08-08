package services

import (
	"server/interfaces"
	"time"
)

// KeyValueService provides business logic for key-value operations
type KeyValueService struct {
	store interfaces.KeyValueStore
}

// NewKeyValueService creates a new key-value service
func NewKeyValueService(store interfaces.KeyValueStore) *KeyValueService {
	return &KeyValueService{
		store: store,
	}
}

// SetKeyValue sets a key-value pair with optional TTL
func (s *KeyValueService) SetKeyValue(key string, value interface{}, ttl *int) bool {
	var ttlDuration *time.Duration
	if ttl != nil {
		duration := time.Duration(*ttl) * time.Second
		ttlDuration = &duration
	}

	return s.store.Store(key, value, ttlDuration)
}

// GetValue retrieves a value by key
func (s *KeyValueService) GetValue(key string) interface{} {
	return s.store.Retrieve(key)
}

// DeleteKey deletes a key-value pair
func (s *KeyValueService) DeleteKey(key string) bool {
	return s.store.Delete(key)
}

// KeyExists checks if a key exists
func (s *KeyValueService) KeyExists(key string) bool {
	return s.store.Exists(key)
}
