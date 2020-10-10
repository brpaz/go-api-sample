// +build unit

package todo_test

import (
	"errors"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateUseCase_Execute_Created(t *testing.T) {

	createdTodo := todo.Todo{
		ID:          1,
		Description: "My todo",
		CreatedAt:   time.Now(),
	}

	repoMock := &todo.MockRepository{}
	repoMock.On("Create", mock.MatchedBy(func(input todo.Todo) bool {
		return input.Description == createdTodo.Description
	})).Return(createdTodo, nil)

	uc := todo.NewCreateUseCase(repoMock)

	ret, err := uc.Execute(todo.CreateTodo{
		Description: "My todo",
	})

	assert.Nil(t, err)
	assert.Equal(t, ret, createdTodo)

	repoMock.AssertExpectations(t)
}

func TestCreateUseCase_Execute_Error(t *testing.T) {

	repoMock := &todo.MockRepository{}
	repoMock.On("Create", mock.Anything).Return(todo.Todo{}, errors.New("error"))

	uc := todo.NewCreateUseCase(repoMock)

	_, err := uc.Execute(todo.CreateTodo{
		Description: "My todo",
	})

	assert.NotNil(t, err)

	repoMock.AssertExpectations(t)
}
