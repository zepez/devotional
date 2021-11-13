package utils

import "os"

func GetEnvUtil(env string, fallback string) string {
	value := os.Getenv(env)

	if len(value) < 1 {
		value = fallback
	}

	return value
}
