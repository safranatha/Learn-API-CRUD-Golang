package main

import (
	"go-fiber/database"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var app = fiber.New()

	// database
	database.DatabseInit()

	// migration
	// migration.RunMigration()

	// // create
	// handler.UserCreate()

	// Routes
	route.Route_init(app)

	app.Listen(":3000")
}
