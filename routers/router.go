package routers

import (
	"fiber-layout/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := controllers.NewDefaultController()
	// GET /api/register
	app.Get("/api/*", main.Api)

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", main.Flights)

	// GET /dictionary.txt
	app.Get("/:file.:ext", main.File)

	// GET /john/75
	app.Get("/:name/:age/:gender?", main.Gender)

	// GET /john
	app.Get("/:name", main.Name)
}
