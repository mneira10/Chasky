package utils

import (
	"fmt"
	"log"
	"os"
)

func GetEnvironmentVariable(envVarKey string) string {
	envVarVal := os.Getenv(envVarKey)

	if envVarVal == "" {
		log.Fatal(fmt.Sprintf("No environment variable for %s", envVarKey))
		// exit with input/ouput error
		os.Exit(5)
	}

	return envVarVal
}
