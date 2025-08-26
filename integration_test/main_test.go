package integrationtest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"Microservice-Laboratory/food-delivery-user-service/internal/api"
)

func TestHealthCheckEndpoi(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	router := api.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)

	//Act
	router.ServeHTTP(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedCacheControl := "no-store"
	if w.Header().Get("Cache-Control") != expectedCacheControl {
		t.Errorf("Expected Cache-Control %s, got %s", expectedCacheControl, w.Header().Get("Cache-Control"))
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedStatus := "UP"
	if response["status"] != expectedStatus {
		t.Errorf("Expected status 'UP', got %s", response["status"])
	}
}
