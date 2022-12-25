package main

import (
	"fiber-layout/routers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	routers.SetRoute(app)

	log.Fatal(app.Listen(":3000"))
}