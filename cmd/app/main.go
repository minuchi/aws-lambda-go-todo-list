package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/fiberadapter"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/minuchi/aws-lambda-go-todo-list/pkg/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHealth)
	app.Get("/todos", handlers.GetToDos)
	app.Post("/todos", handlers.CreateToDo)
	app.Delete("/todos/:id", handlers.DeleteToDo)
}

func setLogger(app *fiber.App) {
	app.Use(logger.New())
}

func setNewRelic(app *fiber.App, newRelicLicenseKey string) {
	newrelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("todo-list"),
		newrelic.ConfigLicense(newRelicLicenseKey),
		newrelic.ConfigEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	utils.CheckError(err)

	cfg := fibernewrelic.Config{
		Application: newrelicApp,
	}

	app.Use(fibernewrelic.New(cfg))
}

var fiberLambda *fiberadapter.FiberLambda

func main() {
	app := fiber.New()

	if newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY"); newRelicLicenseKey != "" {
		setNewRelic(app, newRelicLicenseKey)
	}
	setLogger(app)
	setUpRoutes(app)

	if utils.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		err := app.Listen(":8080")
		utils.CheckError(err)
	}
}

func MarshalJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, request)
}
