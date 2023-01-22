package main

import (
	"shortener/config"
	"shortener/internal/app"
)

func main() {
	cfg := config.New("./")

	app.Run(cfg)
}
