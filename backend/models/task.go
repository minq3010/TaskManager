package models

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed" gorm:"defalt:false"`
	CreatedAt time.Time `json:"created_at"`
}
