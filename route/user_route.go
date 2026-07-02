package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(api fiber.Router) {
	r := api.Group("/users")

	r.Get("/", service.GetUsers)
	r.Get("/:id", service.GetUser)
	r.Post("/", service.CreateUser)
	r.Put("/:id", service.UpdateUser)
	r.Patch("/:id", service.UpdateUser)
	r.Delete("/:id", service.DeleteUser)
}
