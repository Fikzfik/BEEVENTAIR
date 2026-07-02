package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerMatchRoutes(api fiber.Router) {
	r := api.Group("/matches")

	r.Get("/", service.GetMatches)
	r.Get("/:id", service.GetMatch)
	r.Post("/", service.CreateMatch)
	r.Put("/:id", service.UpdateMatch)
	r.Patch("/:id", service.UpdateMatch)
	r.Delete("/:id", service.DeleteMatch)
}
