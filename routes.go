package main

import (
    "github.com/fiskeben/weather-station-api/handler"
    "github.com/fiskeben/weather-station-api/context"

	"github.com/gorilla/mux"
    "net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
    Model string
	HandlerFunc func(*context.AppContext, string, http.ResponseWriter, *http.Request) (interface{}, int, error)
}

type Routes []Route

func NewRouter(context *context.AppContext) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
        handler := handler.AppHandler{context, route.Model, route.HandlerFunc}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"TemperatureIndex",
		"GET",
		"/temperatures",
        "temperature",
        handler.Index,
	},
	Route{
		"Temperature",
		"GET",
		"/temperatures/{id}",
        "temperature",
        handler.Show,
	},
	Route{
		"TemperatureSave",
		"POST",
		"/temperatures",
        "temperature",
        handler.Save,
	},
    Route{
        "PrecipitationIndex",
        "GET",
        "/precipitations",
        "precipitation",
        handler.Index,
    },
    Route{
        "PrecipitationShow",
        "GET",
        "/precipitations/{id}",
        "precipitation",
        handler.Show,
    },
    Route{
        "PrecipitationSave",
        "POST",
        "/precipitations",
        "precipitation",
        handler.Save,
    },
}
