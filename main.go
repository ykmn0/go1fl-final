package main

import (
	"log"

	"go1fl-final/pkg/server"
)

func main() {
	srv := server.New()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
