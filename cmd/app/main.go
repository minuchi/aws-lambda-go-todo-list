package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World 👋!",
		})
	})
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON([]string{"todo1", "todo2"})
	})
}

var fiberLambda *fiberadapter.FiberLambda

func main() {
	app := fiber.New()

	setUpRoutes(app)

	if utils.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		app.Listen(":3000")
	}

}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, request)
}
