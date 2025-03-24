/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package auditFoxConfiguration

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		InMemory bool   `yaml:"in_memory"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

// Just for now, it will replaced by a interface contract method
func GetConfig() Configuration {

	// config file path should be passed as a variable
	f, err := os.Open("../../config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var config Configuration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		processError(err)
	}

	return config
}

func processError(err error) {
	panic("unimplemented")
}
