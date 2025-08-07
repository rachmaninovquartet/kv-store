package api

import (
	"test_client/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the test routes
func SetupRoutes(r *gin.Engine) {
	// Test deletion workflow
	r.GET("/test_deletion", handlers.TestDeletionHandler)

	// Test overwrite workflow
	r.GET("/test_overwrite", handlers.TestOverwriteHandler)

	// Test get specific key
	r.GET("/test_get/:key", handlers.TestGetKeyHandler)
}
