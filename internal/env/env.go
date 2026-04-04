package env

import "os"

func Getstring(key, fallback string) string {
	if val := os.Getenv(key); val != "" {

		return val
	}

	return fallback
}
