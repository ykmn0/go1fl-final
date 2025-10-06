package main

import (
	"log"
	"os"

	"go1fl-final/pkg/api"
	"go1fl-final/pkg/db"
	"go1fl-final/pkg/server"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api.Init()

	srv := server.New()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
