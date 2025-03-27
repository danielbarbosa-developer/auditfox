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

var exampleAuditLog models.AuditLogRecord

func TestPostAuditLogShouldReturnHTTP200(t *testing.T) {
	router := SetupRouters()

	exampleAuditLog = models.AuditLogRecord{}

	w := httptest.NewRecorder()

	auditLogJson, _ := json.Marshal(exampleAuditLog)
	req, _ := http.NewRequest("POST", auditAddPath, strings.NewReader(string(auditLogJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(auditLogJson), w.Body.String())
}
