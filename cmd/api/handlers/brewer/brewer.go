package brewer

import (
	"encoding/json"
	"gin-tonic-inventory-go/cmd/api/models"
	"gin-tonic-inventory-go/pkg/application"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func AddBrewery(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		newBrewery := &models.Brewer{}
		defer request.Body.Close()
		if err := json.NewDecoder(request.Body).Decode(newBrewery); err != nil {
			e := models.Error{
				HTTPCode: http.StatusBadRequest,
				Code:     123,
				Message:  "Could not decode request",
			}
			models.JSONError(writer, e)
			return
		}

		v := validator.New()
		err := v.Struct(newBrewery)
		if err != nil {
			e := models.Error{
				HTTPCode: http.StatusPreconditionFailed,
				Code:     123,
				Message:  "Request Validation failed",
			}
			models.JSONError(writer, e)
			return
		}

		if err := newBrewery.AddBrewer(request.Context(), app); err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Database insert error",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(newBrewery)
		_, _ = writer.Write(response)
	}
}

func GetAllBreweries(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		defer request.Body.Close()
		breweries, err := models.GetAllBreweries(request.Context(), app)

		if err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Error fetching from database",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(breweries)
		_, _ = writer.Write(response)
	}
}

func AddCountry(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		newCountry := &models.Country{}
		defer request.Body.Close()
		if err := json.NewDecoder(request.Body).Decode(newCountry); err != nil {
			e := models.Error{
				HTTPCode: http.StatusBadRequest,
				Code:     123,
				Message:  "Could not decode request",
			}
			models.JSONError(writer, e)
			return
		}

		v := validator.New()
		err := v.Struct(newCountry)
		if err != nil {
			e := models.Error{
				HTTPCode: http.StatusPreconditionFailed,
				Code:     123,
				Message:  "Request Validation failed",
			}
			models.JSONError(writer, e)
			return
		}

		if err := newCountry.AddCountry(request.Context(), app); err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Database insert error",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(newCountry)
		_, _ = writer.Write(response)
	}
}

func GetCountryByID(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := request.Context().Value(models.CtxKey("uuid"))
		getCountry := &models.Country{CountryID: id.(string)}
		defer request.Body.Close()
		if err := getCountry.GetCountryByID(request.Context(), app); err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Database insert error",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(getCountry)
		_, _ = writer.Write(response)
	}
}

func GetAllCountries(app *application.Application) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		defer request.Body.Close()
		countries, err := models.GetAllCountries(request.Context(), app)

		if err != nil {
			e := models.Error{
				HTTPCode: http.StatusInternalServerError,
				Code:     123,
				Message:  "Error fetching from database",
			}
			models.JSONError(writer, e)
			return
		}

		writer.Header().Set("Content/Type", "application/json")
		response, _ := json.Marshal(countries)
		_, _ = writer.Write(response)
	}
}
