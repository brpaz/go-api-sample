// +build unit

package todo_test

import (
	"encoding/json"
	"errors"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)

func TestListTodoHandler_Handle_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todo", nil)
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	ctx := testutil.CreateEchoTestContext(req, resp)

	useCaseMock := &todo.MockListUseCase{}
	useCaseMock.On("Execute").Return([]todo.Todo{
		{
			ID:          1,
			Description: faker.Lorem().Sentence(3),
			CreatedAt:   faker.Date().Backward(time.Hour * 24),
		},
		{
			ID:          2,
			Description: faker.Lorem().Sentence(3),
			CreatedAt:   faker.Date().Backward(time.Hour * 24),
		},
	}, nil)

	h := todo.NewListTodoHandler(useCaseMock)
	h.Handle(ctx)

	var respObj []todo.Todo
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		assert.Fail(t, "Failed to decode response", err)
	}

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Len(t, respObj, 2)
}

func TestListTodoHandler_Handle_Error(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todo", nil)
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	ctx := testutil.CreateEchoTestContext(req, resp)

	useCaseMock := &todo.MockListUseCase{}
	useCaseMock.On("Execute").Return([]todo.Todo{}, errors.New("some-error"))

	h := todo.NewListTodoHandler(useCaseMock)
	h.Handle(ctx)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
