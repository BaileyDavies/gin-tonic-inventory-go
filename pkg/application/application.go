package application

import (
	"gin-tonic-inventory-go/pkg/config"
	"gin-tonic-inventory-go/pkg/db"
)

// Joins together the config and DB helpers to use in the application

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Init()
	database, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  database,
		Cfg: cfg,
	}, nil
}
