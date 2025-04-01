/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package models

import (
	"time"
)

type AuditLogRecord struct {
	ID            string    `json:"id" validate:"required,uuid4"`                          // Unique identifier of the event
	CorrelationID string    `json:"correlation_id" validate:"required,uuid4"`              // Unique correlationID of the event
	TraceID       string    `json:"trace_id" validate:"required,uuid4"`                    // Unique TraceID of the event
	SystemID      string    `json:"system_id" validate:"required"`                         // Unique SystemID of the architecture
	Timestamp     time.Time `json:"timestamp" validate:"required"`                         // UTC date and time of the event
	UserID        string    `json:"user_id" validate:"required,uuid4"`                     // ID of the user who performed the action
	Username      string    `json:"username" validate:"required"`                          // Name of the user
	UserRole      string    `json:"user_role" validate:"required"`                         // Role of the user
	Action        string    `json:"action" validate:"required"`                            // Action performed (e.g., "LOGIN", "DELETE", "UPDATE")
	Entity        string    `json:"entity" validate:"required"`                            // Affected entity (e.g., "file", "user", "record")
	EntityID      string    `json:"entity_id" validate:"required"`                         // Identifier of the affected entity
	SourceIP      string    `json:"source_ip" validate:"required,ip"`                      // Source IP of the action
	UserAgent     string    `json:"user_agent" validate:"required"`                        // Information about the user agent (e.g., browser, system)
	Description   string    `json:"description" validate:"required"`                       // Detailed description of the event
	Severity      string    `json:"severity" validate:"required,oneof=INFO WARN CRITICAL"` // Allowed severity levels
	Status        string    `json:"status" validate:"required,oneof=SUCCESS FAILURE"`      // Allowed status values
	FieldModified string    `json:"field_modified" validate:"omitempty"`                   // Name of the modified field (optional)
	OldValue      string    `json:"old_value" validate:"omitempty"`                        // Previous value before the change (optional)
	NewValue      string    `json:"new_value" validate:"omitempty"`                        // New value after the change (optional)
}
