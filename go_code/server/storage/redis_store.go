package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisStore provides Redis-based key-value storage
type RedisStore struct {
	redisClient *redis.Client
}

// NewRedisStore creates a new Redis store
func NewRedisStore(addr string) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStore{
		redisClient: client,
	}
}

// Store stores a key-value pair with optional TTL
func (s *RedisStore) Store(key string, value interface{}, ttl *time.Duration) bool {
	ctx := context.Background()

	// Serialize value to JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return false
	}

	if ttl != nil {
		err = s.redisClient.Set(ctx, key, jsonValue, *ttl).Err()
	} else {
		err = s.redisClient.Set(ctx, key, jsonValue, 0).Err()
	}

	return err == nil
}

// Retrieve retrieves a value by key
func (s *RedisStore) Retrieve(key string) interface{} {
	ctx := context.Background()

	result, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	// Try to unmarshal as JSON first
	var value interface{}
	if err := json.Unmarshal([]byte(result), &value); err != nil {
		// If JSON unmarshaling fails, return as string
		return result
	}

	return value
}

// Delete deletes a key-value pair
func (s *RedisStore) Delete(key string) bool {
	ctx := context.Background()

	result, err := s.redisClient.Del(ctx, key).Result()
	return err == nil && result > 0
}

// Exists checks if a key exists
func (s *RedisStore) Exists(key string) bool {
	ctx := context.Background()

	result, err := s.redisClient.Exists(ctx, key).Result()
	return err == nil && result > 0
}
