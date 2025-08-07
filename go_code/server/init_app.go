package main

import (
	"log"
	"os"

	"server/services"
	"server/storage"

	"github.com/gin-gonic/gin"
)

// KeyValueService interface for dependency injection
type KeyValueService interface {
	SetKeyValue(key string, value interface{}, ttl *int) bool
	GetValue(key string) interface{}
	DeleteKey(key string) bool
	KeyExists(key string) bool
}

// getStorageType returns the storage type from environment
func getStorageType() string {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		return "memory"
	}
	return storageType
}

// createStore creates a new store based on configuration
func createStore() storage.KeyValueStore {
	storageType := getStorageType()

	switch storageType {
	case "redis":
		redisAddr := os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "localhost:6379"
		}
		log.Printf("Using Redis storage at %s", redisAddr)
		return storage.NewRedisStore(redisAddr)
	default:
		log.Println("Using in-memory storage")
		return storage.NewInMemoryStore()
	}
}

// getService creates a new service instance
func getService() KeyValueService {
	store := createStore()
	return services.NewKeyValueService(store)
}

// serviceMiddleware injects the service into the Gin context
func serviceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := getService()
		c.Set("service", service)
		c.Next()
	}
}
