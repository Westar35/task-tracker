package main

import (
	"log"
	"task-tracker/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}