package middleware

import (
	"gin-tonic-inventory-go/pkg/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LogRequest(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		logger.Info.Printf("%s - %s %s %s", request.RemoteAddr, request.Proto, request.Method, request.URL.RequestURI())
		next(writer, request, params)

	}
}
