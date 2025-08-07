package api

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the API
func SetupRoutes(r *gin.Engine) {
	// Root endpoint
	r.GET("/", handlers.RootHandler)

	// Set key-value pair
	r.POST("/set", handlers.SetKeyValueHandler)

	// Get value by key
	r.GET("/get/:key", handlers.GetValueHandler)

	// Delete key
	r.DELETE("/delete/:key", handlers.DeleteKeyHandler)

	// Check if key exists
	r.GET("/exists/:key", handlers.KeyExistsHandler)
}
