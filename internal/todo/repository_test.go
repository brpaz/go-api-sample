// +build integrationdb

package todo_test

import (
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgRepository_CreateTodo(t *testing.T) {
	db, err := testutil.GetTestDBConnection()

	if err != nil {
		t.Fatal(err)
	}

	repo := todo.NewPgRepository(db)

	err = repo.CreateTodo(todo.CreateTodo{
		Description: "some-todo",
	})

	assert.Nil(t, err)
}
