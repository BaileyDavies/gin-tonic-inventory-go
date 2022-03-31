package models

import (
	"context"
	"gin-tonic-inventory-go/pkg/application"
)

type Gin struct {
	GinID   string ``
	GinName string ``
	GinDesc string ``
	GinAbv  string ``
}

func (gin *Gin) AddGin(ctx context.Context, app *application.Application) error {
	query := `IMPLEMENT`

	err := app.DB.Client.QueryRowContext(
		ctx,
		query,
	).Scan(&gin.GinID)

	if err != nil {
		return err
	}

	return nil
}
