package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateHandler struct {
	useCase CreateUseCase
}

func NewCreateHandler(useCase CreateUseCase) *CreateHandler {
	return &CreateHandler{
		useCase: useCase,
	}
}
func (h *CreateHandler) Handle(c echo.Context) error {
	var request CreateTodo
	if err := c.Bind(&request); err != nil {
		c.Error(err)
		return nil
	}

	if err := c.Validate(&request); err != nil {
		c.Error(err)
		return nil
	}

	newTodo, err := h.useCase.Execute(request)

	if err != nil {
		c.Error(err)
		return nil
	}

	c.JSON(http.StatusCreated, newTodo)
	return nil
}
