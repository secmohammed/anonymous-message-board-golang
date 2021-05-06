package models

import "time"

type Thread struct {
    BaseModel
    Text           string    `json:"text" binding:"required"`
    DeletePassword string    `json:"delete_password" gorm:"not null" binding:"required"`
    Reported       bool      `json:"reported" gorm:"default:false"`
    BumpedOn       time.Time `json:"bumped_on" gorm:"default:now()"`
    Replies        []Reply   `json:"replies,omitempty"`
}
