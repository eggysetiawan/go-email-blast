package config

import (
	"github.com/eggysetiawan/go-email-blast/logger"
	"github.com/joho/godotenv"
)

func LoadEnv(fileName string) {
	err := godotenv.Load(fileName)

	if err != nil {
		logger.Error("Some error occured. Err: " + err.Error())
	}
}
