package service

import (
	"eventbe/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error { return list(c, repository.GetAllUsers) }
func GetUser(c *fiber.Ctx) error  { return find(c, repository.GetUserByID, "User not found") }
func CreateUser(c *fiber.Ctx) error {
	return create(c, repository.CreateUser, "User created successfully")
}
func UpdateUser(c *fiber.Ctx) error {
	return update(c, repository.UpdateUser, "User updated successfully", "User not found")
}
func DeleteUser(c *fiber.Ctx) error {
	return destroy(c, repository.DeleteUser, "User deleted successfully", "User not found")
}
