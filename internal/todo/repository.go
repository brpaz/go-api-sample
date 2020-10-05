package todo

import "gorm.io/gorm"

type Repository interface {
	FindAll() []Todo
	Create(todo CreateTodo) error
}

type PgRepository struct {
	db *gorm.DB
}

func NewPgRepository(db *gorm.DB) PgRepository {
	return PgRepository{
		db: db,
	}
}

func (repo PgRepository) FindAll() []Todo {
	var todos = make([]Todo, 0)

	return todos
}

func (repo PgRepository) CreateTodo(data CreateTodo) error {
	dbEntity := GormTodo{
		Description: data.Description,
	}

	if err := repo.db.Create(dbEntity).Error; err != nil {
		return err
	}

	return nil
}
