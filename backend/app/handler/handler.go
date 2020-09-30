package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/test-layered2/backend/domain/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TodoHandler interface{
	Todos(c echo.Context) error
}

type todoHandler struct {
	todo service.Todo
}

type NewTodoHandler(todo service.Todo) TodoHandler {
	return &todoHandler{todo: todo}
}

func (t *todoHandler) Todos (c echo.Context) error {
	todoList, err := t.todo.ITodo()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, todoList)
}