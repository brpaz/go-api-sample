// +build integrationdb

package todo_test

import (
	"fmt"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)

func TestPgRepository_CreateTodo(t *testing.T) {
	db := testutil.GetTestDBConnection()

	tx := db.Begin()
	defer tx.Rollback()

	repo := todo.NewPgRepository(tx)

	td, err := repo.CreateTodo(todo.Todo{
		Description: faker.Lorem().Sentence(5),
		CreatedAt:   faker.Date().Backward(24 * 7 * time.Hour),
	})

	assert.Nil(t, err)
	assert.NotNil(t, td.ID)
}

func TestPgRepository_CreateTodoWithInvalidData(t *testing.T) {
	db := testutil.GetTestDBConnection()

	tx := db.Begin()
	defer tx.Rollback()

	repo := todo.NewPgRepository(tx)

	_, err := repo.CreateTodo(todo.Todo{})

	fmt.Println(err)
	assert.NotNil(t, err)
}

func TestPgRepository_FindAll(t *testing.T) {
	db := testutil.GetTestDBConnection()

	tx := db.Begin()
	defer tx.Rollback()

	repo := todo.NewPgRepository(tx)

	todosToCreate := []todo.GormTodo{
		{
			Description: faker.Lorem().Sentence(5),
			CreatedAt:   faker.Date().Backward(24 * 7 * time.Hour),
		},
		{
			Description: faker.Lorem().Sentence(5),
			CreatedAt:   faker.Date().Backward(24 * 7 * time.Hour),
		},
	}

	if err := tx.Create(&todosToCreate).Error; err != nil {
		t.Fatal(err)
	}

	result, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}
