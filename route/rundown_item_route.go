package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerRundownItemRoutes(api fiber.Router) {
	r := api.Group("/rundown-items")

	r.Get("/", service.GetRundownItems)
	r.Get("/:id", service.GetRundownItem)
	r.Post("/", service.CreateRundownItem)
	r.Put("/:id", service.UpdateRundownItem)
	r.Patch("/:id", service.UpdateRundownItem)
	r.Delete("/:id", service.DeleteRundownItem)
}
