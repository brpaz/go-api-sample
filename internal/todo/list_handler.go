package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ListTodoHandler struct {
	useCase ListUseCase
}

func NewListTodoHandler(useCase ListUseCase) *ListTodoHandler {
	return &ListTodoHandler{
		useCase: useCase,
	}
}

func (h *ListTodoHandler) Handle(c echo.Context) error {
	todos, err := h.useCase.Execute()

	if err != nil {
		c.Error(err)
		return nil
	}

	_ = c.JSON(http.StatusOK, todos)
	return nil
}
