package model
import "time"


type Temperature struct {
    Id			int64		`json:"id"`
    Value		float32		`json:"value"`
    Timestamp	time.Time   `json:"timestamp"`
}

type Temperatures []Temperature

func (this Temperature) GetId() (int64) {
    return this.Id
}

func (this Temperature) GetValue() (float32) {
    return this.Value
}

func (this Temperature) GetTimestamp() (time.Time) {
    return this.Timestamp
}
