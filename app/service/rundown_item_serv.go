package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetRundownItems(c *fiber.Ctx) error { return list(c, repository.GetAllRundownItems) }
func GetRundownItem(c *fiber.Ctx) error {
	return find(c, repository.GetRundownItemByID, "Rundown item not found")
}
func CreateRundownItem(c *fiber.Ctx) error {
	return create(c, repository.CreateRundownItem, "Rundown item created successfully")
}
func UpdateRundownItem(c *fiber.Ctx) error {
	return update(c, repository.UpdateRundownItem, "Rundown item updated successfully", "Rundown item not found")
}
func DeleteRundownItem(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteRundownItem, "Rundown item deleted successfully", "Rundown item not found")
}
