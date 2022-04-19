package application

import (
	"context"
	"fmt"
	"gin-tonic-inventory-go/pkg/config"
	"gin-tonic-inventory-go/pkg/db"
	"gin-tonic-inventory-go/pkg/logger"
	"github.com/go-redis/redis/v8"
)

// Joins together the config and DB helpers to use in the application

type Application struct {
	DB    *db.DB
	Cfg   *config.Config
	Cache *redis.Client
}

func Get() (*Application, error) {
	cfg := config.Init()
	database, err := db.Get(cfg.GetDBConnStr())
	ginCache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("redis:6379"),
		Password: "",
		DB:       0,
	})
	_, err = ginCache.Ping(context.TODO()).Result()
	if err != nil {
		logger.Error.Printf("%v", err)
	}

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:    database,
		Cfg:   cfg,
		Cache: ginCache,
	}, nil
}
