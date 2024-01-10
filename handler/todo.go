package handler

import (
	"net/http"

	"github.com/cosmin220304/go-echo-templ-htmx/data/model"
	"github.com/cosmin220304/go-echo-templ-htmx/data/repository"
	view "github.com/cosmin220304/go-echo-templ-htmx/view/feature/user"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoRepo repository.TodoRepo
}

func NewTodoHandler(todoRepo repository.TodoRepo) *TodoHandler {
	return &TodoHandler{
		todoRepo: todoRepo,
	}
}

func (h *TodoHandler) HandleGetFormTemplate(c echo.Context) error {
	todos, err := h.todoRepo.GetAll()
	if err != nil {
		return err
	}
	return render(c, view.TodoForm(todos))
}

func (h *TodoHandler) HandleGetItemTemplate(c echo.Context) error {
	id := c.Param("id")
	todo, err := h.todoRepo.GetById(id)
	if err != nil {
		return err
	}
	return render(c, view.TodoItem(todo))
}

func (h *TodoHandler) HandlePostTodoItem(c echo.Context) error {
	newTodo := model.Todo{Name: c.FormValue("name")}
	if newTodo.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	createdTodo, err := h.todoRepo.Create(newTodo)
	if err != nil {
		return err
	}
	return render(c, view.TodoItem(createdTodo))
}

func (h *TodoHandler) HandleDeleteTodoItem(c echo.Context) error {
	id := c.Param("id")
	err := h.todoRepo.DeleteById(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
