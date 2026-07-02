package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerChatMessageRoutes(api fiber.Router) {
	r := api.Group("/chat-messages")

	r.Get("/", service.GetChatMessages)
	r.Get("/:id", service.GetChatMessage)
	r.Post("/", service.CreateChatMessage)
	r.Put("/:id", service.UpdateChatMessage)
	r.Patch("/:id", service.UpdateChatMessage)
	r.Delete("/:id", service.DeleteChatMessage)
}
