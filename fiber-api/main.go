package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my Awosome API")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUser)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
