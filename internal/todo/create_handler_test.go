// +build unit

package todo_test

import (
	"bytes"
	"encoding/json"
	"github.com/brpaz/go-api-sample/internal/errors"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreteTodoHandler_InvalidPayload(t *testing.T) {

	var jsonBody = []byte(`{"abcc":""}`)

	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	ctx := testutil.CreateEchoTestContext(req, resp)

	useCaseMock := &todo.MockCreateUseCase{}

	h := todo.NewCreateHandler(useCaseMock)
	h.Handle(ctx)

	var respObj appHttp.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		assert.Fail(t, "Failed to decode response", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
	assert.Equal(t, errors.ErrCodeValidationFailed, respObj.Code)
	assert.Len(t, respObj.Fields, 1)
}

func TestCreteTodoHandler_Success(t *testing.T) {

	var jsonBody = []byte(`{"description":"test"}`)

	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	ctx := testutil.CreateEchoTestContext(req, resp)

	useCaseMock := &todo.MockCreateUseCase{}
	useCaseMock.On("Execute", todo.CreateTodo{
		Description: "test",
	}).Return(todo.Todo{ID: 1, Description: "test", CreatedAt: time.Now()}, nil)

	h := todo.NewCreateHandler(useCaseMock)
	h.Handle(ctx)

	/*var respObj appHttp.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		assert.Fail(t, "Failed to decode response", err)
	}*/

	assert.Equal(t, http.StatusCreated, resp.Code)
}
