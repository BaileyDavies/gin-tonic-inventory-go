package models

import (
	"context"
	"gin-tonic-inventory-go/pkg/application"
	"github.com/google/uuid"
)

type Gin struct {
	GinID    string `json:"gin_id"`
	GinName  string `json:"gin_name"`
	GinDesc  string `json:"gin_desc"`
	GinAbv   string `json:"gin_abv"`
	BrewerID string `json:"brewer_id"`
}

func (gin *Gin) AddGin(ctx context.Context, app *application.Application) error {
	ginUuid := uuid.New().String()
	query := `INSERT INTO Gins VALUES($1, $2, $3, $4, $5)`
	_, err := app.DB.Client.QueryContext(
		ctx,
		query,
		ginUuid,
		gin.GinName,
		gin.GinDesc,
		gin.GinAbv,
		gin.BrewerID,
	)
	if err != nil {
		return err
	}
	gin.GinID = ginUuid
	return nil
}
