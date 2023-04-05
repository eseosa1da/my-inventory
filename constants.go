package main

import "os"

func getEnvOrDefault(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}

	return defaultVal
}

var DBName = getEnvOrDefault("DB_NAME", "inventory")
var DbUser = getEnvOrDefault("DB_USER", "admin")
var DbPassword = getEnvOrDefault("DB_PASSWORD", "Password12")
var DBHost = getEnvOrDefault("DB_HOST", "goproject.cwqtgpeadrce.us-east-1.rds.amazonaws.com:3306")
