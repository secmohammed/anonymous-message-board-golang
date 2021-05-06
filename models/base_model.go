package models

import "time"

type BaseModel struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}
