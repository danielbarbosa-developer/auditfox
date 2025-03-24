/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package auditFoxConfiguration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigShouldReturnAuditFoxConfigurationStruct(t *testing.T) {

	config := GetAuditFoxConfig()

	assert.Equal(t, config.Cluster.Nodes, 1)
	assert.Equal(t, config.InputMode.Https, true)
	assert.Equal(t, config.InputMode.Agents, false)
	assert.Equal(t, config.InputMode.Queue, false)
}
