package repository

import (
	"net/http"

	"github.com/cosmin220304/go-echo-templ-htmx/data/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoRepo interface {
	GetAll() ([]model.Todo, error)
	GetById(id string) (model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(todo model.Todo) (model.Todo, error)
	DeleteById(id string) error
}

// implement User interface

type TodoRepoConfig struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepoConfig {
	return &TodoRepoConfig{
		db: db,
	}
}

func (r *TodoRepoConfig) GetAll() ([]model.Todo, error) {
	var todos []model.Todo = make([]model.Todo, 0)
	result := r.db.First(&todos)
	return todos, result.Error
}

func (r *TodoRepoConfig) GetById(id string) (model.Todo, error) {
	var todo model.Todo
	result := r.db.Find(&todo, id)
	if result.RowsAffected == 0 {
		return todo, echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	}
	return todo, result.Error
}

func (r *TodoRepoConfig) Create(todo model.Todo) (model.Todo, error) {
	result := r.db.Create(&todo)
	return todo, result.Error
}

func (r *TodoRepoConfig) Update(todo model.Todo) (model.Todo, error) {
	result := r.db.Save(&todo)
	return todo, result.Error
}

func (r *TodoRepoConfig) DeleteById(id string) error {
	var todo model.Todo
	result := r.db.Delete(&todo, id)
	return result.Error
}
