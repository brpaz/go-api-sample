package todo

import (
	"github.com/brpaz/go-api-sample/internal/errors"
	"gorm.io/gorm"
	"time"
)

type GormTodo struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"column:content"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (GormTodo) TableName() string {
	return "todos"
}

// ToDomain Converts the Database entity into Domain entity
func (e GormTodo) ToDomain() Todo {
	return Todo{
		ID:          e.ID,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
	}
}

func (e GormTodo) FromDomain(t Todo) GormTodo {
	return GormTodo{
		ID:          t.ID,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
	}
}

type PgRepository struct {
	db *gorm.DB
}


func NewPgRepository(db *gorm.DB) *PgRepository {
	return &PgRepository{
		db: db,
	}
}

func (repo *PgRepository) FindAll() ([]Todo, error) {
	var dbTodos = make([]GormTodo, 0)
	var domainTodos = make([]Todo, 0)

	if err := repo.db.Find(&dbTodos).Error; err != nil {
		return domainTodos, errors.NewApplicationError(errors.ErrCodeInternalError, "failed to fetch list of todos").
			WithOriginalError(err)
	}

	for _, dbTodo := range dbTodos {
		domainTodos = append(domainTodos, dbTodo.ToDomain())
	}

	return domainTodos, nil
}

func (repo *PgRepository) Create(todo Todo) (Todo, error) {
	dbEntity := GormTodo{}.FromDomain(todo)

	if err := repo.db.Save(&dbEntity).Error; err != nil {
		return Todo{}, errors.NewApplicationError(errors.ErrCodeInternalError, "failed to create todo").
			WithOriginalError(err)
	}

	return dbEntity.ToDomain(), nil
}
