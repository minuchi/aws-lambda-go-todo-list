package main

import (
	"context"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"os"

	"github.com/minuchi/aws-lambda-go-todo-list/pkg/middlewares"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"
)

var fiberLambda *fiberadapter.FiberLambda

func main() {
	app := fiber.New()

	middlewares.LogMiddleware(app)
	if newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY"); newRelicLicenseKey != "" {
		middlewares.NewRelicMiddleware(app, newRelicLicenseKey)
	}

	routes.AppRouter(app)
	routes.ToDoRouter(app)
	routes.NotFoundRouter(app)

	if utils.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		err := app.Listen(":8080")
		utils.CheckError(err)
	}
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, request)
}
