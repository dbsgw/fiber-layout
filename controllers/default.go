package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type DefaultController struct{}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (t *DefaultController) Api(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ ---- %s", c.Params("*"))
	return c.SendString(msg) // => ✋ register
}

func (t *DefaultController) Flights(c *fiber.Ctx) error {
	msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
	return c.SendString(msg) // => 💸 From: LAX, To: SFO
}

func (t *DefaultController) File(c *fiber.Ctx) error {
	msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
	return c.SendString(msg) // => 📃 dictionary.txt
}

func (t *DefaultController) Gender(c *fiber.Ctx) error {
	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
	return c.SendString(msg) // => 👴 john is 75 years old
}

func (t *DefaultController) Name(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", string(c.Params("name")))
	return c.SendString(msg) // => Hello john 👋!
}
