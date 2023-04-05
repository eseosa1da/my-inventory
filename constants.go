package main

import "os"

func getEnvOrDefault(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}

	return defaultVal
}

var DBName = getEnvOrDefault("DB_NAME", "test")
var DbUser = getEnvOrDefault("DB_USER", "root")
var DbPassword = getEnvOrDefault("DB_PASSWORD", "Password12")
var DBHost = getEnvOrDefault("DB_HOST", "localhost")
