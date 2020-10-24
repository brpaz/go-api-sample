package todo

import (
	"time"
)

type Repository interface {
	FindAll() ([]Todo, error)
	Create(todo Todo) (Todo, error)
}

type Todo struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(description string) Todo {
	return Todo{
		CreatedAt:   time.Now(),
		Description: description,
	}
}

type CreateTodo struct {
	Description string `json:"description" validate:"required"`
}
