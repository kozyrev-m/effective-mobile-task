package main

import (
	"log"

	"github.com/kozyrev-m/effective-mobile-task/internal/config"
	"github.com/kozyrev-m/effective-mobile-task/internal/server"
)

func main() {
	cfg := config.InitConfig() 
	if err := server.StartApp(cfg); err != nil {
		log.Fatal(err)
	}
}
