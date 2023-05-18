package models

import "time"

type Employees struct {
	Id           int8      `json:"id" gorm:"primaryKey,unique"`
	Username     string    `json:"username" gorm:"index"`
	Password     string    `json:"-"`
	FullName     string    `json:"full_name" gorm:"index"`
	Position     string    `json:"position"`
	SessionToken string    `json:"session_token"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive     bool      `json:"is_active" gorm:"index;default:true"`
}
