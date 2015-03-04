package handler

import (
	"log"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
    "github.com/fiskeben/weather-station-api/model"
    "github.com/fiskeben/weather-station-api/context"

	"github.com/gorilla/mux"
)

func Index(context *context.AppContext, modelName string, w http.ResponseWriter, r *http.Request) (interface {}, int, error) {
	from := startAt(r)
	duration := duration(r)
	to := from.Add(duration)

	log.Printf("Getting %ss from %s to %s", modelName, from, to)

	objects, err := model.FindAll(context, modelName, from, to)

	if err != nil {
		log.Printf("Error retrieving list of %s: %s", modelName, err.Error())
		return nil, http.StatusInternalServerError, err
	}
    w.Header().Add("x-next-timestamp", to.Format(time.RFC822Z))

    return objects, http.StatusOK, nil
}

func Show(context *context.AppContext, modelName string, w http.ResponseWriter, r *http.Request) (interface {}, int, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	log.Printf("Getting %s with id '%d'", modelName, id)

	if err != nil {
		log.Printf("Unable to parse id '%s'", vars["id"])
		return nil, http.StatusBadRequest, err
	}

	object, err := model.Find(context, modelName, id)

	if (err != nil) {
        log.Printf("Unable to find %s with id '%d'", modelName, id)
        return nil, http.StatusNotFound, err
    }

    return *object, http.StatusOK, nil
}

func Save(context *context.AppContext, modelName string, w http.ResponseWriter, r *http.Request) (interface {}, int, error) {
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
        log.Printf("Unable to read body: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}
	
	object, err := model.GetZeroModelFromName(modelName)

    if err != nil {
        return nil, http.StatusInternalServerError, err
    }

	err = json.Unmarshal(buffer, object)

	if err != nil {
		log.Printf("Unable to unmarshal JSON as %s: %s", modelName, err.Error())
		return nil, http.StatusBadRequest, err
	}

	var id int64
	id, err = model.Save(context, modelName, object)

	if err != nil {
		log.Printf("Error saving %s: %s", modelName, err.Error())
		return nil, http.StatusInternalServerError, err
	}

	result, err := model.Find(context, modelName, id)

	if err != nil {
		log.Printf("Unable to load %s after save: %s", modelName, err.Error())
		return nil, http.StatusBadRequest, err
	}

    return *result, http.StatusOK, nil
}

func defaultStartAt() (time.Time) {
	duration, _ := time.ParseDuration("-24h")
	start := time.Now().Add(duration)
	return start
}

func defaultDuration() (time.Duration) {
	duration, _ := time.ParseDuration("24h")
	return duration
}

func startAt(r *http.Request) (time.Time) {
	startAtHeader := r.Header.Get("X-Start-At")

	if startAtHeader == "" {
		return defaultStartAt()
	}

	start, err := time.Parse(time.RFC822Z, startAtHeader)

	if err != nil {
		log.Printf("Unable to parse start at header: %s", err.Error())
		return defaultStartAt()
	}

	return start
}

func duration(r *http.Request) (time.Duration) {
	durationHeader := r.Header.Get("X-Duration")
	if durationHeader == "" {
		return defaultDuration()
	}

	duration, err := time.ParseDuration(durationHeader)
	
	if err != nil {
		log.Printf("Unable to parse duration: %s", err.Error())
		return defaultDuration()
	}

	return duration
}