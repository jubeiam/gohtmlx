package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // hot reload

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, Iwona!",
		}, "layouts/main")
	})

	app.Post("/random-name", func(c *fiber.Ctx) error {
		return c.Render("partials/name", fiber.Map{
			"Title": "Hello, Leszek!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
