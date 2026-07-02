package route

import (
	"eventbe/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerSubmissionRoutes(api fiber.Router) {
	r := api.Group("/submissions")

	r.Get("/", service.GetSubmissions)
	r.Get("/:id", service.GetSubmission)
	r.Post("/", service.CreateSubmission)
	r.Put("/:id", service.UpdateSubmission)
	r.Patch("/:id", service.UpdateSubmission)
	r.Delete("/:id", service.DeleteSubmission)
}
