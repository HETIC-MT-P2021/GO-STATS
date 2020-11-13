package service

import (
	"os"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable Allow to get environment variable in .env file
func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

// GetVarsBot Allow to get environment variable for setting up bot in .env file
func GetVarsBot() (BotConfig, error) {
	var config BotConfig

	token, err := GoDotEnvVariable("TOKENDISCORD")
	if err != nil {
		return config, err
	}

	config = BotConfig{
		Token: token,
	}

	return config, nil
}
