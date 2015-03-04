package model

import "time"

type Humidity struct {
    Id			int64		`json:"id"`
    Value		float32		`json:"value"`
    Timestamp	time.Time   `json:"timestamp"`
}

type Humidities []Humidity

func (this Humidity) GetId() (int64) {
    return this.Id
}

func (this Humidity) GetValue() (float32) {
    return this.Value
}

func (this Humidity) GetTimestamp() (time.Time) {
    return this.Timestamp
}