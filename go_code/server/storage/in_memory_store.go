package storage

import (
	"sync"
	"time"

	"server/interfaces"
)

// InMemoryStore provides in-memory key-value storage
type InMemoryStore struct {
	data map[string]interface{}
	ttl  map[string]time.Time
	mu   sync.RWMutex
}

// NewInMemoryStore creates a new in-memory store
func NewInMemoryStore() interfaces.KeyValueStore {
	return &InMemoryStore{
		data: make(map[string]interface{}),
		ttl:  make(map[string]time.Time),
	}
}

// Store stores a key-value pair with optional TTL
func (s *InMemoryStore) Store(key string, value interface{}, ttl *time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	if ttl != nil {
		s.ttl[key] = time.Now().Add(*ttl)
	} else {
		delete(s.ttl, key)
	}
	return true
}

// Retrieve retrieves a value by key
func (s *InMemoryStore) Retrieve(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.data[key]
	if !exists {
		return nil
	}

	// Check TTL
	if ttl, hasTTL := s.ttl[key]; hasTTL && time.Now().After(ttl) {
		// Value has expired, remove it
		s.mu.RUnlock()
		s.mu.Lock()
		delete(s.data, key)
		delete(s.ttl, key)
		s.mu.Unlock()
		s.mu.RLock()
		return nil
	}

	return value
}

// Delete deletes a key-value pair
func (s *InMemoryStore) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		delete(s.ttl, key)
		return true
	}
	return false
}

// Exists checks if a key exists
func (s *InMemoryStore) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.data[key]
	if !exists {
		return false
	}

	// Check TTL
	if ttl, hasTTL := s.ttl[key]; hasTTL && time.Now().After(ttl) {
		return false
	}

	return true
}
