package main

import (
	_ "fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<b>It's Work!!!</b>")
	})

	app.Listen(":3000")
}
