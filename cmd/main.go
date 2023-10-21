package main

import (
	"log"

	"github.com/kozyrev-m/effective-mobile-task/internal/server"
)

func main() {
	if err := server.StartApp(); err != nil {
		log.Fatal(err)
	}
}
