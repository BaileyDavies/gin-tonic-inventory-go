package main

import (
	"gin-tonic-inventory-go/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	cfg := config.Init()

	direction := cfg.GetMigration()
	if direction != "down" && direction != "up" {
		log.Println("can only migrate up or down")
		return
	}

	m, err := migrate.New("file://db/migrations", cfg.GetDBConnStr())
	if err != nil {
		log.Printf("%s", err)
		return
	}

	if direction == "up" {
		err := m.Up()
		if err != nil {
			log.Printf("failed to migrate with error: %s", err)
		}
	}

	if direction == "down" {
		err = m.Down()
		if err != nil {
			log.Printf("failed to migrate with error %s", err)
			return
		}
	}
}
