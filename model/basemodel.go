package model

import "time"

type BaseModel interface {
    GetId() int64
    GetValue() float32
    GetTimestamp() time.Time
}

type BaseModels []BaseModel