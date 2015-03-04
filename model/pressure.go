package model

import "time"

type Pressure struct {
    Id			int64		`json:"id"`
    Value		float32		`json:"value"`
    Timestamp	time.Time   `json:"timestamp"`
}

type Pressures []Pressure

func (this Pressure) GetId() (int64) {
    return this.Id
}

func (this Pressure) GetValue() (float32) {
    return this.Value
}

func (this Pressure) GetTimestamp() (time.Time) {
    return this.Timestamp
}