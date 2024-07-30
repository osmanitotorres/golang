package main

import (
	"github.com/gofiber/fiber/v2"
)

// Define o handler para o upload
func uploadHandler(c *fiber.Ctx) error {
	// Lógica do upload
	return nil
}

// Configura as rotas
func setupRoutes(app *fiber.App) {
	app.Post("/upload", uploadHandler)
}
