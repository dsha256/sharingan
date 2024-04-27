package config

import encore "encore.dev"

var (
	EnvName = encore.Meta().Environment.Name

	FeesTaskQueue = EnvName + "-fees"
)
