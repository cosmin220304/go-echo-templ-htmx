package main

import (
	"github.com/cosmin220304/go-echo-templ-htmx/data"
	"github.com/cosmin220304/go-echo-templ-htmx/data/repository"
	"github.com/cosmin220304/go-echo-templ-htmx/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	db := data.ConnectDB()

	todoRepo := repository.NewTodoRepo(db)

	todoHandler := handler.NewTodoHandler(todoRepo)

	app.GET("/", todoHandler.HandleGetFormTemplate)
	app.GET("/todos/:id", todoHandler.HandleGetItemTemplate)
	app.POST("/todos", todoHandler.HandlePostTodoItem)
	app.DELETE("/todos/:id", todoHandler.HandleDeleteTodoItem)

	app.Start(":8080")
}
