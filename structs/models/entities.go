package models

import "time"

type PhoneNumber struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CallId    string    `gorm:"type:varchar(64);unique_index:call_id" json:"call_id"`
	Number    string    `gorm:"type:varchar(64)" json:"number"`
	Email     string    `gorm:"type:varchar(64)" json:"email"`
	Type      int       `gorm:"type:int" json:"type"`
	Status    int       `gorm:"type:int;default:0;index" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
