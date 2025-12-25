package main

import (
	"log"

	"pbi/internal/config/container"
	"pbi/internal/config/db"
)

func main() {
	cont := container.InitContainer()

	if err := db.SeedAdmin(cont.DB.Raw); err != nil {
		log.Fatal("seed admin failed:", err)
	}

	log.Println("seed admin SUCCESS")
}
