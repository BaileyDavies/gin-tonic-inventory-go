package middleware

import (
	"context"
	"gin-tonic-inventory-go/cmd/api/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
)

const UUID_REGEXP = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"

func ValidateUuidRequest(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		uuid := params.ByName("id")
		uuidMatch, _ := regexp.MatchString(UUID_REGEXP, uuid)
		if !uuidMatch {
			err := models.Error{
				HTTPCode: http.StatusPreconditionFailed,
				Code:     123,
				Message:  "Request validation failed",
			}
			models.JSONError(writer, err)
			return
		}

		ctx := context.WithValue(request.Context(), models.CtxKey("uuid"), uuid)
		request = request.WithContext(ctx)
		next(writer, request, params)
	}
}
