package utils

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // Load .env file
)

type AppConfig struct {
	APIDomain     string
	WebsiteDomain string
	STAuthURI     string
	STAuthAPIKey  string
	AppName       string
	Port          int
	Mode          string
	LogLevel      string
	AuthDebug     bool
}

func getEnvStringOrError(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Please set the " + key + " environment variable")
	}
	return value
}

func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvIntegerOrError(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic("Please set the " + key + " environment variable")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("The " + key + " environment variable must be an integer")
	}
	return intValue
}
func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic("The " + key + " environment variable must be a boolean")
	}
	return boolValue
}

func GetAppConfig() AppConfig {
	return AppConfig{
		APIDomain:     getEnvString("API_DOMAIN", "http://localhost:8080"),
		WebsiteDomain: getEnvString("WEBSITE_DOMAIN", "http://localhost:5173"),
		STAuthURI:     getEnvStringOrError("ST_AUTH_URI"),
		STAuthAPIKey:  getEnvStringOrError("ST_AUTH_API_KEY"),
		AppName:       getEnvString("APP_NAME", "My Server"),
		Port:          getEnvIntegerOrError("PORT"),
		Mode:          getEnvString("MODE", "production"),
		LogLevel:      getEnvString("LOG_LEVEL", "info"),
		AuthDebug:     getEnvBool("AUTH_DEBUG", false),
	}
}
