package config

import "os"

type AwsConfig struct {
	Region    string
	AccessKey string
	SecretKey string
}

type Config struct {
	PORT           string
	AwsConfig      AwsConfig
	DynamoEndpoint string
	JWTSecret      string
}

var instance *Config

func getEnvWithDefaultValue(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Init() {
	if instance != nil {
		return
	}

	instance = &Config{
		PORT: getEnvWithDefaultValue("PORT", ""),
		AwsConfig: AwsConfig{
			Region:    getEnvWithDefaultValue("AWS_REGION", ""),
			AccessKey: getEnvWithDefaultValue("AWS_ACCESS_KEY", ""),
			SecretKey: getEnvWithDefaultValue("AWS_SECRET_KEY", ""),
		},
		DynamoEndpoint: getEnvWithDefaultValue("DYNAMO_ENDPOINT", "http://localhost:8000"),
		JWTSecret:      getEnvWithDefaultValue("JWT_SECRET_KEY", "default value"),
	}

}

func GetInstance() *Config {
	return instance
}
