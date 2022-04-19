package models

import (
	"context"
	"gin-tonic-inventory-go/pkg/application"
	"gin-tonic-inventory-go/pkg/logger"
	"github.com/google/uuid"
)

type Brewer struct {
	BrewerID          string `json:"brewer_id" validate:"required"`
	BrewerName        string `json:"brewer_name" validate:"required"`
	BrewerDescription string `json:"brewer_description" validate:"required"`
	BrewerWebsite     string `json:"brewer_website" validate:"required"`
	CountryID         string `json:"country_id" validate:"required"`
}

func (brewer *Brewer) AddBrewer(ctx context.Context, app *application.Application) error {
	query := `INSERT INTO Brewers VALUES ($1, $2, $3, $4, $5)`
	brewer.BrewerID = uuid.New().String()

	_, err := app.DB.Client.ExecContext(
		ctx,
		query,
		brewer.BrewerID,
		brewer.BrewerName,
		brewer.BrewerDescription,
		brewer.BrewerWebsite,
		brewer.CountryID,
	)

	if err != nil {
		logger.Error.Printf("%s, %v", "Could not add brewer with error: ", err)
		return err
	}

	return nil
}

func GetAllBreweries(ctx context.Context, app *application.Application) ([]Brewer, error) {
	query := `SELECT * FROM brewers`

	res, err := app.DB.Client.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		logger.Error.Printf("%v", err)
		return nil, err
	}

	var breweries []Brewer

	for res.Next() {
		logger.Error.Printf("%v", res)
		var brewer Brewer
		err := res.Scan(&brewer.BrewerID, &brewer.BrewerName, &brewer.BrewerDescription, &brewer.BrewerWebsite, brewer.CountryID)
		if err != nil {
			logger.Error.Printf("%v", err)
			continue
		}
		breweries = append(breweries, brewer)
	}
	return breweries, nil
}
