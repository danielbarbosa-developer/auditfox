package routes

import (
	"auditfox/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostAuditLogShouldReturnHTTP200(t *testing.T) {
	router := setupRouter()
	router = postAuditLog(router)

	w := httptest.NewRecorder()

	exampleAuditLog := models.AuditLogRecord{}

	userJson, _ := json.Marshal(exampleAuditLog)
	req, _ := http.NewRequest("POST", "/auditlog/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
