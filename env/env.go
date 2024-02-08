package env

import (
	"github.com/joho/godotenv"
	"os"
)

func GetVarFromDotEnv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}
