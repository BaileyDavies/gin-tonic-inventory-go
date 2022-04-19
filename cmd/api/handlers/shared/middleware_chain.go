package shared

import (
	"gin-tonic-inventory-go/pkg/application"
	"gin-tonic-inventory-go/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func ValidateUuidLogAndDo(app *application.Application, fn func(*application.Application) httprouter.Handle) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.ValidateUuidRequest,
	}

	return middleware.Chain(fn(app), mdw...)

}

func LogAndDo(app *application.Application, fn func(*application.Application) httprouter.Handle) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
	}

	return middleware.Chain(fn(app), mdw...)

}
