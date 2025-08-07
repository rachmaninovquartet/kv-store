package handlers

import (
	"net/http"

	"server/models"

	"github.com/gin-gonic/gin"
)

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
	TTL   *int        `json:"ttl,omitempty"`
}

// KeyValueService interface for dependency injection
type KeyValueService interface {
	SetKeyValue(key string, value interface{}, ttl *int) bool
	GetValue(key string) interface{}
	DeleteKey(key string) bool
	KeyExists(key string) bool
}

// RootHandler handles the root endpoint
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "KV Service is running"})
}

// SetKeyValueHandler handles setting a key-value pair
func SetKeyValueHandler(c *gin.Context) {
	var item KeyValue
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := c.MustGet("service").(KeyValueService)
	success := service.SetKeyValue(item.Key, item.Value, item.TTL)

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store key-value pair"})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Key '" + item.Key + "' stored successfully"})
}

// GetValueHandler handles getting a value by key
func GetValueHandler(c *gin.Context) {
	key := c.Param("key")
	service := c.MustGet("service").(KeyValueService)
	value := service.GetValue(key)

	if value == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, models.ValueResponse{Value: value})
}

// DeleteKeyHandler handles deleting a key
func DeleteKeyHandler(c *gin.Context) {
	key := c.Param("key")
	service := c.MustGet("service").(KeyValueService)
	success := service.DeleteKey(key)

	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Key '" + key + "' deleted successfully"})
}

// KeyExistsHandler handles checking if a key exists
func KeyExistsHandler(c *gin.Context) {
	key := c.Param("key")
	service := c.MustGet("service").(KeyValueService)
	exists := service.KeyExists(key)

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
