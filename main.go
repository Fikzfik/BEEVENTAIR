package main

import (
	"eventbe/config"
	"eventbe/database"
	"eventbe/route"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	database.ConnectMySQL()

	app := config.NewApp()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	route.RegisterRoutes(app)

	app.Static("/static", "./uploads")

	port := config.GetEnv("APP_PORT", "3000")
	app.Listen(":" + port)
}
