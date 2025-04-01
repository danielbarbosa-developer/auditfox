package services

import "auditfox/models"

type Validator interface {
	isValid(value any) models.ValidationResult
}
