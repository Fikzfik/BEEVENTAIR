package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetPrizes(c *fiber.Ctx) error { return list(c, repository.GetAllPrizes) }
func GetPrize(c *fiber.Ctx) error  { return find(c, repository.GetPrizeByID, "Prize not found") }
func CreatePrize(c *fiber.Ctx) error {
	return create(c, repository.CreatePrize, "Prize created successfully")
}
func UpdatePrize(c *fiber.Ctx) error {
	return update(c, repository.UpdatePrize, "Prize updated successfully", "Prize not found")
}
func DeletePrize(c *fiber.Ctx) error {
	return destroy(c, repository.DeletePrize, "Prize deleted successfully", "Prize not found")
}
