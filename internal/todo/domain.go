package todo

import (
	"time"
)

type Todo struct {
	ID          int32     `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodo struct {
	Description string `json:"description"`
}

type GormTodo struct {
	ID          uint `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
