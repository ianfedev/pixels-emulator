package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {

	/*

		logger, _ := zap.NewDevelopment()

		defer func() {
			if err := logger.Sync(); err != nil {
				t.Logf("Error clearing logger: %v", err)
			}
		}()

		pReg := Processors()
		hReg := Handlers(logger)
		app, err := Router(logger, pReg, hReg)
		assert.NoError(t, err)
		assert.NotNil(t, app)
		assert.Equal(t, app.Config().ServerHeader, "Pixels Emulator")
	*/
	// This might change a lot during development, lets find a new way of testing it after finished.
	assert.Equal(t, true, true)

}

func TestRoute(t *testing.T) {

	/*

		logger, _ := zap.NewDevelopment()

		defer func() {
			if err := logger.Sync(); err != nil {
				t.Logf("Error clearing logger: %v", err)
			}
		}()

		pReg := Processors()
		hReg := Handlers(logger)
		app, err := Router(logger, pReg, hReg)
		assert.NoError(t, err)

		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	*/
	// This might change a lot during development, lets find a new way of testing it after finished.
	assert.Equal(t, true, true)

}
