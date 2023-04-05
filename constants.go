package main

import "os"

// func getEnvOrDefault(env, defaultVal string) string {
// 	if val := os.Getenv(env); val != "" {
// 		return val
// 	}

// 	return defaultVal
// }

// var DBName, DBUser, DBPassword, DBHost string

var DBName = os.Getenv("DB_NAME")
var DBUser = os.Getenv("DB_USER")
var DBPassword = os.Getenv("DB_PASSWORD")
var DBHost = os.Getenv("DB_HOST")

// DBName = getEnvOrDefault("DB_NAME", "inventory")
// DBUser = getEnvOrDefault("DB_USER", "admin")
// DBPassword = getEnvOrDefault("DB_PASSWORD", "Password12")
// DBHost = getEnvOrDefault("DB_HOST", "goproject.cwqtgpeadrce.us-east-1.rds.amazonaws.com:3306")
