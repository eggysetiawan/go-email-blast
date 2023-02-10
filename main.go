package main

import (
	"github.com/eggysetiawan/go-email-blast/app"
	"github.com/eggysetiawan/go-email-blast/config"
	"github.com/eggysetiawan/go-email-blast/logger"
)

func main() {
	config.LoadEnv(".env")

	logger.Info("Starting application...")

	app.Router()

	//fmt.Println(os.Getenv("SMTP"))
}
