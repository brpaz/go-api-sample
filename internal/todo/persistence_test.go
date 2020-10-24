// +build unit

package todo_test

import (
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)

func TestGormTodo_ToDomain(t *testing.T) {
	creationDate := faker.Date().Backward(5 * time.Hour)

	dbTodo := todo.GormTodo{
		ID:          1,
		Description: "test",
		CreatedAt:   creationDate,
	}

	domainTodo := dbTodo.ToDomain()
	assert.Equal(t, "test", domainTodo.Description)
	assert.Equal(t, creationDate, domainTodo.CreatedAt)
	assert.Equal(t, uint(1), domainTodo.ID)
}

func TestGormTodo_FromDomain(t *testing.T) {

	domainTodo := todo.NewTodo("test")

	dbTodo := todo.GormTodo{}.FromDomain(domainTodo)

	assert.Equal(t, "test", dbTodo.Description)
	assert.NotNil(t, dbTodo.CreatedAt)
}
