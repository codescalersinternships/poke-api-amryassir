package pkg

import (
	"os"
	"strconv"
	"time"
)

// Config - represents a configuration object
type Config struct {
	URL     string
	Timeout time.Duration
}

// LoadConfig - initializes and returns a Config struct with configuration values
func LoadConfig() Config {
	timeoutStr := getEnv("TIMEOUT", "10")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		timeout = 10
	}

	return Config{
		URL:     getEnv("URL", "https://pokeapi.co/api/v2/pokemon/"),
		Timeout: time.Duration(timeout) * time.Second,
	}
}

// getEnv - retrieves the value of an environment variable
func getEnv(key, defaultvalue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultvalue
	}
	return value
}
