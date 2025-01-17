package config

import (
	"log"
	"os"
)

func Get(key EnvKey) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		log.Fatalf("Env lookup failed: %s", key)
	}
	return value
}
