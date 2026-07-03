package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Serve static upload folders
	app.Static("/uploads", "./uploads")

	api := app.Group("/api/v1")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api.Post("/register", service.Register)
	api.Post("/login", service.Login)
	api.Post("/upload", service.UploadFile)

	registerUserRoutes(api)
	registerEventRoutes(api)
	registerPrizeRoutes(api)
	registerRundownItemRoutes(api)
	registerTeamRoutes(api)
	registerTeamMemberRoutes(api)
	registerRoundRoutes(api)
	registerMatchRoutes(api)
	registerSubmissionRoutes(api)
	registerChatMessageRoutes(api)
}
