// +build unit

package todo_test

import (
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListUseCase_Execute(t *testing.T) {
	repo := &todo.MockRepository{}
	repo.On("FindAll").Return([]todo.Todo{
		{
			ID:          1,
			Description: "test",
			CreatedAt:   time.Now(),
		},
	}, nil)
	uc := todo.NewListUseCase(repo)

	result, err := uc.Execute()

	assert.Nil(t, err)
	assert.Len(t, result, 1)

}
