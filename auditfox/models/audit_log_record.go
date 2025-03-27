/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package models

import "time"

type AuditLogRecord struct {
	ID            string    `json:"id"`             // Unique identifier of the event
	CorrelationID string    `json:"correlation_id"` // Unique correlationID of the event
	TraceID       string    `json:"trace_id"`       // Unique TraceID of the event
	SystemID      string    `json:"system_id"`      // Unique SystemID of the architecture
	Timestamp     time.Time `json:"timestamp"`      // UTC date and time of the event
	UserID        string    `json:"user_id"`        // ID of the user who performed the action
	Username      string    `json:"username"`       // Name of the user
	UserRole      string    `json:"user_role"`      // Role of the user
	Action        string    `json:"action"`         // Action performed (e.g., "LOGIN", "DELETE", "UPDATE")
	Entity        string    `json:"entity"`         // Affected entity (e.g., "file", "user", "record")
	EntityID      string    `json:"entity_id"`      // Identifier of the affected entity
	SourceIP      string    `json:"source_ip"`      // Source IP of the action
	UserAgent     string    `json:"user_agent"`     // Information about the user agent (e.g., browser, system)
	Description   string    `json:"description"`    // Detailed description of the event
	Severity      string    `json:"severity"`       // Severity level (e.g., "INFO", "WARN", "CRITICAL")
	Status        string    `json:"status"`         // Outcome of the action (e.g., "SUCCESS", "FAILURE")
	FieldModified string    `json:"field_modified"` // Name of the modified field
	OldValue      string    `json:"old_value"`      // Previous value before the change
	NewValue      string    `json:"new_value"`      // New value after the change
}
