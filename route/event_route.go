package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerEventRoutes(api fiber.Router) {
	r := api.Group("/events")

	r.Get("/", service.GetEvents)
	r.Get("/:id", service.GetEvent)
	r.Post("/", service.CreateEvent)
	r.Put("/:id", service.UpdateEvent)
	r.Patch("/:id", service.UpdateEvent)
	r.Delete("/:id", service.DeleteEvent)
}
