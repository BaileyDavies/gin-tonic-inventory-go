package router

import (
	"gin-tonic-inventory-go/cmd/api/handlers/brewer"
	"gin-tonic-inventory-go/cmd/api/handlers/gin"
	"gin-tonic-inventory-go/cmd/api/handlers/shared"
	"gin-tonic-inventory-go/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.POST("/gins", gin.AddGin(app))

	mux.POST("/brewer/country", shared.LogAndDo(app, brewer.AddCountry))
	mux.POST("/brewer", shared.LogAndDo(app, brewer.AddBrewery))
	mux.GET("/brewer/country", shared.LogAndDo(app, brewer.GetAllCountries))
	mux.GET("/brewer/country/:id", shared.ValidateUuidLogAndDo(app, brewer.GetCountryByID))
	return mux

}
