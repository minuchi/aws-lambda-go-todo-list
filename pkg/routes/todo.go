package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/handlers"
)

func ToDoRouter(app *fiber.App) {
	app.Get("/todos", handlers.GetToDos)
	app.Post("/todos", handlers.CreateToDo)
	app.Delete("/todos/:id", handlers.DeleteToDo)
}
