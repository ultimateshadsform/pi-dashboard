package routes

import "github.com/gofiber/fiber/v3"

func SetupRouter(app *fiber.App) {

	// Catch all route
	app.Get("*", func(c fiber.Ctx) error {
		// Return index.html in ./static
		return c.SendFile("./static/index.html")
	})
}
