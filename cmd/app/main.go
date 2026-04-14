package main

import (
	"log"
	"task-tracker/internal/server"
	"task-tracker/internal/storage/postgres"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=Fgrths197+ dbname=task_tracker sslmode=disable"

	db, err := postgres.ConnectToDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = postgres.InitSchema(db)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run(db)
	if err != nil {
		log.Fatal(err)
	}
}
