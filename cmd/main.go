package main

import (
	"fmt"
	"shortener/config"
	"shortener/internal/app"
)

func main() {
	cfg := config.New("./")
	fmt.Println(cfg)

	app.Run(cfg)
}
