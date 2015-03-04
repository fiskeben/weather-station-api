package model

import (
    "fmt"
	"time"
    "errors"
    "github.com/fiskeben/weather-station-api/context"
    "log"
)

func Find(appContext *context.AppContext, modelName string, id int64) (*BaseModel, error) {
    var (
        value float32
        timestamp time.Time
    )

    sql := fmt.Sprintf("SELECT value, created_at FROM %ss WHERE id = $1", modelName)

    err := appContext.Db.QueryRow(sql, id).Scan(&value, &timestamp)

    if err != nil {
        return nil, err
    }

    result, err := GetModelFromName(modelName, id, value, timestamp)

    if err != nil {
        return nil, err
    }

    return &result, nil
}

func Save(appContext *context.AppContext, modelName string, object BaseModel) (id int64, err error) {
	var generatedId int64

    sql := fmt.Sprintf("INSERT INTO %ss (value, created_at) VALUES ($1, NOW()) RETURNING id", modelName)

	err = appContext.Db.QueryRow(sql, object.GetValue()).Scan(&generatedId)

	if err != nil {
		return -1, err
	}

	return generatedId, nil
}

func FindAll(appContext *context.AppContext, modelName string, from time.Time, to time.Time) (t []BaseModel, err error) {
	var (
		id int64
		value float32
		timestamp time.Time
	)

    sql := fmt.Sprintf("SELECT id, value, created_at FROM %ss WHERE created_at BETWEEN $1 AND $2 ORDER BY created_at DESC", modelName)

	res, err := appContext.Db.Query(sql, from, to)

	if err != nil {
		return nil, err
	}

	result := BaseModels{}

	defer res.Close()

	for res.Next() {
		err = res.Scan(&id, &value, &timestamp)
		if err != nil {
			return nil, err
		}

		object, err := GetModelFromName(modelName, id, value, timestamp)

        if err != nil {
            return nil, err
        }

		result = append(result, object)
	}

	err = res.Err()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetModelFromName(modelName string, id int64, value float32, timestamp time.Time) (BaseModel, error) {
    switch modelName {
        case "temperature":
            return Temperature{id, value, timestamp}, nil
        case "humidity":
            return Humidity{id, value, timestamp}, nil
        case "precipitation":
            return Precipitation{id, value, timestamp}, nil
        case "pressure":
            return Pressure{id, value, timestamp}, nil
        default:
            log.Printf("Unknown model name '%s'", modelName)
            return nil, errors.New("Unknown model " + modelName)
    }
}

func GetZeroModelFromName(modelName string) (BaseModel, error) {
    switch modelName {
        case "temperature":
            return &Temperature{}, nil
        case "humidity":
            return &Humidity{}, nil
        case "precipitation":
            return &Precipitation{}, nil
        case "pressure":
            return &Pressure{}, nil
        default:
            log.Printf("Unknown model name '%s'", modelName)
            return nil, errors.New("Unknown model " + modelName)
    }
}