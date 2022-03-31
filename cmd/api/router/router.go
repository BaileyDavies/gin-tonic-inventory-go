package router

import (
	"gin-tonic-inventory-go/cmd/api/handlers/gin"
	"gin-tonic-inventory-go/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/gins", gin.GetGins(app))
	return mux

}
