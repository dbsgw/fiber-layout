package routers

import (
	"fiber-layout/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	// GET /register 	get
	app.Get("/register", main.Register)

	// GET /login 	json
	app.Post("/login", main.Login)

}
