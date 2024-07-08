package pkg

import "os"

// Config - represents a configuration object
type Config struct {
	URL string
}

// LoadConfig - initializes and returns a Config struct with configuration values
func LoadConfig() Config {
	return Config{
		URL: getEnv("URL", "https://pokeapi.co/api/v2/pokemon/"),
	}
}

// getEnv - retrieves the value of an environment variable
func getEnv(key, defaultvalue string) string {
	value, err := os.LookupEnv(key)
	if !err {
		return defaultvalue
	}
	return value
}
