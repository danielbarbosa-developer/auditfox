/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package auditFoxConfiguration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerConfigShouldReturnConfigStruct(t *testing.T) {
	config := GetConfig()

	assert.Equal(t, config.Server.Host, "localhost")
	assert.Equal(t, config.Server.Port, 8000)

	assert.Equal(t, config.Database.InMemory, true)
	assert.Equal(t, config.Database.Username, "admin")
	assert.Equal(t, config.Database.Password, "super-user-1980")
}
