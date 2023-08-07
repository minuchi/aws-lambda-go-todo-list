package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

type ToDo struct {
	ID    string `json:"id" xml:"id" form:"id"`
	Title string `json:"title" xml:"title" form:"title"`
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World 👋!",
		})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON([]ToDo{
			{ID: "1", Title: "Buy milk"},
			{ID: "2", Title: "Buy eggs"},
		})
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		id := uuid.New().String()
		toDo := new(ToDo)
		if err := c.BodyParser(toDo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if toDo.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing title",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id, "title": toDo.Title})
	})
}

func setLogger(app *fiber.App) {
	app.Use(logger.New())
}

var fiberLambda *fiberadapter.FiberLambda

func main() {
	app := fiber.New()

	setLogger(app)
	setUpRoutes(app)

	if utils.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		app.Listen(":8080")
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, request)
}
