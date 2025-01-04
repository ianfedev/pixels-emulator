package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSetupRouter(t *testing.T) {

	logger, _ := zap.NewDevelopment()

	defer func() {
		if err := logger.Sync(); err != nil {
			t.Logf("Error clearing logger: %v", err)
		}
	}()

	app, err := SetupRouter(logger)
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, app.Config().ServerHeader, "Pixels Emulator")

}

func TestRoute(t *testing.T) {

	logger, _ := zap.NewDevelopment()

	defer func() {
		if err := logger.Sync(); err != nil {
			t.Logf("Error clearing logger: %v", err)
		}
	}()

	app, err := SetupRouter(logger)
	assert.NoError(t, err)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

}
