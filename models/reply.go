package models

type Reply struct {
    BaseModel
    Text           string `json:"text" binding:"required"`
    DeletePassword string `json:"delete_password" gorm:"not null" binding:"required"`
    Reported       bool   `json:"reported" gorm:"default:false"`
    ThreadID       uint   `json:"thread_id"`
}
