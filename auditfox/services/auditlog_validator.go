package services

import (
	"auditfox/models"

	"github.com/go-playground/validator/v10"
)

type AuditLogValidator struct {
}

func (auditValidator AuditLogValidator) isValid(value any) models.ValidationResult {
	validate := validator.New()
	validationResult := models.ValidationResult{
		IsValid: false,
		Errors:  make([]string, 0),
	}

	auditLog, ok := value.(models.AuditLogRecord)

	if !ok {
		validationResult.Errors = append(validationResult.Errors, "Not a valid auditlog type")
		return validationResult
	}

	err := validate.Struct(auditLog)

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationResult.Errors = append(validationResult.Errors, fieldErr.Error())
		}
		return validationResult
	}

	validationResult.IsValid = true

	return validationResult
}
