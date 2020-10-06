package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateHandler struct {
	useCase CreateUseCase
}

func NewCreateHandler(useCase CreateUseCase)  CreateHandler {
	return CreateHandler{
		useCase: useCase,
	}
}
func (h  *CreateHandler)Handle(c echo.Context) {
	var request CreateTodo
	if err := c.Bind(&request); err != nil {
		c.Error(err)
		return
	}

	if err := c.Validate(&request); err != nil {
		c.Error(err)
		return
	}

	newTodo, err := h.useCase.Execute(request)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}
