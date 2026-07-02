package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetSubmissions(c *fiber.Ctx) error { return list(c, repository.GetAllSubmissions) }
func GetSubmission(c *fiber.Ctx) error {
	return find(c, repository.GetSubmissionByID, "Submission not found")
}
func CreateSubmission(c *fiber.Ctx) error {
	return create(c, repository.CreateSubmission, "Submission created successfully")
}
func UpdateSubmission(c *fiber.Ctx) error {
	return update(c, repository.UpdateSubmission, "Submission updated successfully", "Submission not found")
}
func DeleteSubmission(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteSubmission, "Submission deleted successfully", "Submission not found")
}
