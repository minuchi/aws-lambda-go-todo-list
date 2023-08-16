package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/handlers"
)

func AppRouter(app *fiber.App) {
	app.Get("/health", handlers.GetHealth)
}
