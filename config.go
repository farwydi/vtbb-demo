package main

import "github.com/farwydi/cleanwhale/config"

type AppConfig struct {
	Project config.ProjectConfig

	GraphQLTransport config.HTTPConfig
}
