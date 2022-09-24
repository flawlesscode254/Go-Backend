package main

import (
	"github.com/gofiber/fiber/v2"
)

// Slash route
func slash(ctx *fiber.Ctx) error {
	return ctx.SendString("Slash Route")
}

// Policy route
func policy(ctx *fiber.Ctx) error {
	return ctx.SendString("Policy Route")
}

func main() {
	// App initialization
	app := fiber.New()

	// Main route
	app.Get("/", slash)

	// Defined route
	app.Get("/policy", policy)

	// Route params
	app.Get("/names/:value", func(ctx *fiber.Ctx) error {
		return ctx.SendString(ctx.Params("value"))
	})

	// Serving static files
	app.Static("/static", "./public")

	// Listening route
	app.Listen(":3450")
}
