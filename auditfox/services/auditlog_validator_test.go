package services

import (
	"auditfox/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuditLogvalidatorShouldReturnValid(t *testing.T) {
	auditLog := models.AuditLogRecord{
		ID:            uuid.NewString(),
		CorrelationID: uuid.NewString(),
		TraceID:       uuid.NewString(),
		SystemID:      "System-123",
		Timestamp:     time.Now().UTC(),
		UserID:        uuid.NewString(),
		Username:      "johndoe",
		UserRole:      "admin",
		Action:        "LOGIN",
		Entity:        "user",
		EntityID:      uuid.NewString(),
		SourceIP:      "192.168.1.1",
		UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		Description:   "User logged in successfully",
		Severity:      "INFO",
		Status:        "SUCCESS",
		FieldModified: "",
		OldValue:      "",
		NewValue:      "",
	}

	validator := AuditLogValidator{}
	result := validator.isValid(auditLog)

	assert.Equal(t, true, result.IsValid)
}

func TestAuditLogValidatorShouldReturnInvalidWithAllFieldsEmpty(t *testing.T) {
	// Criando um AuditLogRecord inválido (todos os campos obrigatórios estão vazios)
	auditLog := models.AuditLogRecord{
		ID:            "",
		CorrelationID: "",
		TraceID:       "",
		SystemID:      "",
		Timestamp:     time.Time{}, // Tempo zero (inválido)
		UserID:        "",
		Username:      "",
		UserRole:      "",
		Action:        "",
		Entity:        "",
		EntityID:      "",
		SourceIP:      "",
		UserAgent:     "",
		Description:   "",
		Severity:      "",
		Status:        "",
		FieldModified: "",
		OldValue:      "",
		NewValue:      "",
	}

	validator := AuditLogValidator{}
	result := validator.isValid(auditLog)

	assert.Equal(t, false, result.IsValid)

	t.Logf("Total de erros: %d", len(result.Errors))

	assert.Equal(t, 16, len(result.Errors))
}
