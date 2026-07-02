package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error { return list(c, repository.GetAllEvents) }
func GetEvent(c *fiber.Ctx) error  { return find(c, repository.GetEventByID, "Event not found") }
func CreateEvent(c *fiber.Ctx) error {
	return create(c, repository.CreateEvent, "Event created successfully")
}
func UpdateEvent(c *fiber.Ctx) error {
	return update(c, repository.UpdateEvent, "Event updated successfully", "Event not found")
}
func DeleteEvent(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteEvent, "Event deleted successfully", "Event not found")
}
