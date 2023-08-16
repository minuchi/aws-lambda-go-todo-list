package middlewares

import (
	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/fiber/v2"
	"github.com/minuchi/aws-lambda-go-todo-list/pkg/utils"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicMiddleware(app *fiber.App, newRelicLicenseKey string) {
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
