package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHealth)
	app.Get("/todos", handlers.GetToDos)
	app.Post("/todos", handlers.CreateToDo)
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
