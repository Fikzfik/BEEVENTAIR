package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerRoundRoutes(api fiber.Router) {
	r := api.Group("/rounds")

	r.Get("/", service.GetRounds)
	r.Get("/:id", service.GetRound)
	r.Post("/", service.CreateRound)
	r.Put("/:id", service.UpdateRound)
	r.Patch("/:id", service.UpdateRound)
	r.Delete("/:id", service.DeleteRound)
}
