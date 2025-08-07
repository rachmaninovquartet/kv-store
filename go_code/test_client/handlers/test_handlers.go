package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"test_client/models"

	"github.com/gin-gonic/gin"
)

// KeyValue represents a key-value pair for testing
type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	TTL   *int        `json:"ttl,omitempty"`
}

// getServerURL returns the server URL from environment
func getServerURL() string {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		return "http://localhost:8000"
	}
	return serverURL
}

// TestDeletionHandler handles the deletion workflow test
func TestDeletionHandler(c *gin.Context) {
	serverURL := getServerURL()

	// Set a key first
	setData := KeyValue{
		Key:   "test_delete_key",
		Value: "test_value",
	}

	setResp, err := http.Post(serverURL+"/set", "application/json",
		strings.NewReader(fmt.Sprintf(`{"key":"%s","value":"%s"}`, setData.Key, setData.Value)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to connect to main server",
			Details: err.Error(),
		})
		return
	}
	defer setResp.Body.Close()

	if setResp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to set key for deletion test",
		})
		return
	}

	// Check if key exists
	existsResp, err := http.Get(serverURL + "/exists/" + setData.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to check if key exists",
			Details: err.Error(),
		})
		return
	}
	defer existsResp.Body.Close()

	if existsResp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to check if key exists",
		})
		return
	}

	var existsResult map[string]interface{}
	if err := json.NewDecoder(existsResp.Body).Decode(&existsResult); err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to parse exists response",
			Details: err.Error(),
		})
		return
	}

	if exists, ok := existsResult["exists"].(bool); !ok || !exists {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Key does not exist after setting",
		})
		return
	}

	// Delete the key
	req, err := http.NewRequest("DELETE", serverURL+"/delete/"+setData.Key, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to create delete request",
			Details: err.Error(),
		})
		return
	}

	deleteResp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to delete key",
			Details: err.Error(),
		})
		return
	}
	defer deleteResp.Body.Close()

	if deleteResp.StatusCode == http.StatusOK {
		var deleteResult map[string]interface{}
		json.NewDecoder(deleteResp.Body).Decode(&deleteResult)
		c.JSON(http.StatusOK, models.TestResponse{
			Message: "Test deletion successful",
			Details: deleteResult,
		})
	} else {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Test deletion failed",
			Details: deleteResp.StatusCode,
		})
	}
}

// TestOverwriteHandler handles the overwrite workflow test
func TestOverwriteHandler(c *gin.Context) {
	serverURL := getServerURL()
	key := "test_overwrite_key"

	// Set initial value
	initialData := KeyValue{
		Key:   key,
		Value: "initial_value",
	}

	initialJSON, _ := json.Marshal(initialData)
	initialResp, err := http.Post(serverURL+"/set", "application/json",
		strings.NewReader(string(initialJSON)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to set initial value",
			Details: err.Error(),
		})
		return
	}
	defer initialResp.Body.Close()

	// Overwrite with new value
	overwriteData := KeyValue{
		Key:   key,
		Value: "overwritten_value",
	}

	overwriteJSON, _ := json.Marshal(overwriteData)
	overwriteResp, err := http.Post(serverURL+"/set", "application/json",
		strings.NewReader(string(overwriteJSON)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to overwrite value",
			Details: err.Error(),
		})
		return
	}
	defer overwriteResp.Body.Close()

	// Get the final value
	getResp, err := http.Get(serverURL + "/get/" + key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to get final value",
			Details: err.Error(),
		})
		return
	}
	defer getResp.Body.Close()

	if getResp.StatusCode == http.StatusOK {
		var getResult map[string]interface{}
		json.NewDecoder(getResp.Body).Decode(&getResult)
		c.JSON(http.StatusOK, models.TestResponse{
			Message: "Test overwrite successful",
			Details: getResult,
		})
	} else {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Test overwrite failed",
			Details: getResp.StatusCode,
		})
	}
}

// TestGetKeyHandler handles testing getting a specific key
func TestGetKeyHandler(c *gin.Context) {
	key := c.Param("key")
	serverURL := getServerURL()

	resp, err := http.Get(serverURL + "/get/" + key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.TestResponse{
			Message: "Failed to get key",
			Details: err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		c.JSON(http.StatusOK, models.TestResponse{
			Message: "Get key successful",
			Details: result,
		})
	} else {
		var errorResult map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errorResult)
		c.JSON(http.StatusOK, models.TestResponse{
			Message: "Get key failed",
			Details: errorResult,
		})
	}
}
