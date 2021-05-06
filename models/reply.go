package models

type Reply struct {
    BaseModel
    Text           string `json:"text"`
    DeletePassword string `json:"delete_password" gorm:"not null"`
    Reported       bool   `json:"reported" gorm:"default:false"`
    ThreadID       uint   `json:"thread_id"`
}
