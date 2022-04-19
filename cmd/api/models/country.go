package models

import (
	"context"
	"encoding/json"
	"gin-tonic-inventory-go/pkg/application"
	"gin-tonic-inventory-go/pkg/logger"
	"github.com/google/uuid"
)

type Country struct {
	CountryID   string `json:"country_id"`
	CountryName string `json:"country_name"`
}

func (country *Country) AddCountry(ctx context.Context, app *application.Application) error {
	countryUuid := uuid.New().String()
	query := `INSERT INTO Country VALUES ($1, $2)`

	logger.Error.Printf("%v", country.CountryName)
	logger.Error.Printf("%s", countryUuid)

	_, err := app.DB.Client.ExecContext(
		ctx,
		query,
		countryUuid,
		country.CountryName,
	)
	if err != nil {
		logger.Error.Printf("%v", err)
		return err
	}
	return nil
}

func (country *Country) GetCountryByID(ctx context.Context, app *application.Application) error {
	err := getCountryFromCache(ctx, app, country)
	if country.CountryName != "" {
		return nil
	}
	logger.Error.Printf("%s", "Not in cache!")
	query := `SELECT country_id, country_name FROM Country WHERE country_id = $1`
	err = app.DB.Client.QueryRowContext(
		ctx,
		query,
		country.CountryID,
	).Scan(&country.CountryID, &country.CountryName)
	if err != nil {
		logger.Error.Printf("%v", err)
		return err
	}
	err = storeCountryInCache(ctx, app, country)
	if err != nil {
		logger.Error.Printf("%v", err)
	}
	return nil
}

func GetAllCountries(ctx context.Context, app *application.Application) ([]Country, error) {
	query := `SELECT country_id, country_name FROM Country`

	res, err := app.DB.Client.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		logger.Error.Printf("%v", err)
		return nil, err
	}

	var countries []Country

	for res.Next() {
		logger.Error.Printf("%v", res)
		var country Country
		err := res.Scan(&country.CountryID, &country.CountryName)
		if err != nil {
			logger.Error.Printf("%v", err)
			continue
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func storeCountryInCache(ctx context.Context, app *application.Application, country *Country) error {
	countryJsonRaw, err := json.Marshal(country)
	if err != nil {
		return err
	}
	err = app.Cache.Set(ctx, country.CountryID, string(countryJsonRaw), 0).Err()
	if err != nil {
		logger.Error.Printf("%v", err)
		return err
	}
	return nil
}

func getCountryFromCache(ctx context.Context, app *application.Application, country *Country) error {
	cacheCountry, err := app.Cache.Get(ctx, country.CountryID).Result()
	if err != nil {
		logger.Error.Printf("%s", err)
		return err
	}
	err = json.Unmarshal([]byte(cacheCountry), country)
	if err != nil {
		logger.Error.Printf("%v", err)
		return nil
	}
	return nil
}
