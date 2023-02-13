package main

import (
	"github.com/eggysetiawan/go-email-blast/app"
	"github.com/eggysetiawan/go-email-blast/config"
)

func main() {
	config.LoadEnv(".env")

	app.Console()
}
