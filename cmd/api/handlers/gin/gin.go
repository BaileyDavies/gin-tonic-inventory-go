package gin

import (
	"encoding/json"
	"gin-tonic-inventory-go/cmd/api/models"
	"gin-tonic-inventory-go/pkg/application"
	"gin-tonic-inventory-go/pkg/logger"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func AddGin(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		logger.Error.Printf("%v", "test")
		newGin := &models.Gin{}
		defer request.Body.Close()
		if err := json.NewDecoder(request.Body).Decode(newGin); err != nil {
			e := models.Error{
				HTTPCode: http.StatusBadRequest,
				Code:     123,
				Message:  "Could not decode request",
			}
			models.JSONError(writer, e)
			return
		}

		v := validator.New()
		err := v.Struct(newGin)
		if err != nil {
			e := models.Error{
				HTTPCode: http.StatusPreconditionFailed,
				Code:     123,
				Message:  "Request Validation failed",
			}
			models.JSONError(writer, e)
			return
		}

		if err := newGin.AddGin(request.Context(), app); err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Database insert error",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(newGin)
		_, _ = writer.Write(response)
	}
}
