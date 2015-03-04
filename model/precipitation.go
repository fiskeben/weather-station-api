package model
import "time"


type Precipitation struct {
    Id			int64		`json:"id"`
    Value		float32		`json:"value"`
    Timestamp	time.Time   `json:"timestamp"`
}

type Precipitations []Precipitation

func (this Precipitation) GetId() (int64) {
    return this.Id
}

func (this Precipitation) GetValue() (float32) {
    return this.Value
}

func (this Precipitation) GetTimestamp() (time.Time) {
    return this.Timestamp
}
