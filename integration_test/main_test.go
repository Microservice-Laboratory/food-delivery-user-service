package integrationtest_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"Microservice-Laboratory/food-delivery-user-service/internal/api"
)

func TestHealthCheckEndpoi(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
		},
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	// Ensure the container is cleaned up after the test finishes
	defer postgresC.Terminate(ctx)

	host, err := postgresC.Host(ctx)
	require.NoError(t, err)
	port, err := postgresC.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err)

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=test password=test dbname=test sslmode=disable", host, port.Port())

	// Open the database connection
	db, err := sql.Open("pgx", connStr)
	require.NoError(t, err)
	defer db.Close()

	router := api.SetupRouter(db)

	w := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)

	//Act
	router.ServeHTTP(w, request)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedCacheControl := "no-store"
	if w.Header().Get("Cache-Control") != expectedCacheControl {
		t.Errorf("Expected Cache-Control %s, got %s", expectedCacheControl, w.Header().Get("Cache-Control"))
	}

	var response map[string]string
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &response)
	if errUnmarshal != nil {
		t.Fatalf("Failed to unmarshal response body: %v", errUnmarshal)
	}

	expectedStatus := "UP"
	if response["status"] != expectedStatus {
		t.Errorf("Expected status 'UP', got %s", response["status"])
	}

	defer postgresC.Terminate(ctx)
}
