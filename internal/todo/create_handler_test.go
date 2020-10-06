// +build unit

package todo_test

import (
	"bytes"
	"fmt"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreteTodoHandler_InvalidPayload(t *testing.T) {

	var jsonBody = []byte(`{"abcc":""}`)

	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonBody))
	req.Header.Add("Content-type", "application/json")
	rec := httptest.NewRecorder()
	ctx := testutil.CreateEchoTestContext(req, rec)

	useCaseMock := &todo.MockCreateUseCase{}

	h := todo.NewCreateHandler(useCaseMock)
	h.Handle(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	fmt.Println(rec.Body.String())
}
