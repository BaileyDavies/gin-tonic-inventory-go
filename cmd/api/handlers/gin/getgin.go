package gin

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"gin-tonic-inventory-go/pkg/application"
)

func GetGins(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "hello")
	}
}
