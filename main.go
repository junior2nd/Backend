package main

import (
	"msa/database"
	"msa/database/migration.go"
	"msa/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Database
	database.DatabaseInit()
	//Migration
	migration.RunMigration()

	app := fiber.New()

	// TimeZoneSeting
	app.Use(logger.New(logger.Config{
		TimeFormat: "Asia/Bangkok",
	}))

	// Route
	route.RouteInit(app)

	app.Listen(":4000")

}
