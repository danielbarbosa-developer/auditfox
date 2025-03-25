/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package auditfoxconfiguration

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AuditFoxConfiguration struct {
	Cluster struct {
		Nodes int `yaml:"nodes"`
	} `yaml:"cluster"`
	InputMode struct {
		Https  bool `yaml:"https"`
		Agents bool `yaml:"agents"`
		Queue  bool `yaml:"queue"`
	} `yaml:"input_mode"`
}

// Just for now, it will replaced by a interface contract method
func GetAuditFoxConfig() AuditFoxConfiguration {

	// config file path should be passed as a variable
	f, err := os.Open("../../auditfox_config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var config AuditFoxConfiguration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		processError(err)
	}

	return config
}
