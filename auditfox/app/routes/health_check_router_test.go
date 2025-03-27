package routes

import (
	"auditfox/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedHealthCheck models.HealthCheck

func TestHealthCheckGetShouldReturnHTTP200(t *testing.T) {
	router := SetupRouters()

	w := httptest.NewRecorder()

	expectedHealthCheck = models.HealthCheck{
		Status: "RUNNING",
	}

	healthCheckJson, _ := json.Marshal(expectedHealthCheck)
	req, _ := http.NewRequest("GET", healthCheckPath, nil)
	router.ServeHTTP(w, req)

	result := w.Body.String()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(healthCheckJson), result)
}
