package handler

import (
    "net/http"
    "strings"

    "github.com/fiskeben/weather-station-api/context"

    "encoding/json"
    "encoding/xml"
)

type AppHandler struct {
    *context.AppContext
    ModelName string
    Handle func(*context.AppContext, string, http.ResponseWriter, *http.Request) (interface{}, int, error)
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    if payload, status, err := ah.Handle(ah.AppContext, ah.ModelName, w, r); err != nil {
        switch status {
            case http.StatusNotFound:
                notFound(w, r)
            case http.StatusInternalServerError:
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            default:
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
    } else {
        acceptsContentType := r.Header.Get("Accept")

        if strings.Contains(acceptsContentType, "xml") {
            w.Header().Set("Content-Type", "application/xml")
            xml.NewEncoder(w).Encode(payload)
        } else {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(payload)
        }
    }
}

func notFound(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w,r)
}
