package service

import (
	"errors"
	"strings"
	"time"

	"eventbe/app/repository"
	"eventbe/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func list(c *fiber.Ctx, fetch func() ([]map[string]any, error)) error {
	data, err := fetch()
	if err != nil {
		return helper.InternalError(c, err.Error())
	}

	return helper.APIResponse(c, fiber.StatusOK, "Success", data)
}

func find(c *fiber.Ctx, fetch func(string) (map[string]any, error), notFoundMessage string) error {
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	data, err := fetch(id)
	if err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, "Success", data)
}

func create(c *fiber.Ctx, save func(map[string]any) (map[string]any, error), message string) error {
	input, err := parseInput(c)
	if err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}

	now := time.Now()
	input["id"] = uuid.New().String()
	setDefaultTimestamp(input, "created_at", now)
	setDefaultTimestamp(input, "updated_at", now)
	setDefaultTimestamp(input, "timestamp", now)
	setDefaultTimestamp(input, "joined_at", now)

	data, err := save(input)
	if err != nil {
		return helper.InternalError(c, err.Error())
	}

	return helper.APIResponse(c, fiber.StatusCreated, message, data)
}

func update(c *fiber.Ctx, save func(string, map[string]any) (map[string]any, error), message, notFoundMessage string) error {
	input, err := parseInput(c)
	if err != nil {
		return helper.BadRequest(c, "Invalid input payload")
	}
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	delete(input, "id")

	data, err := save(id, input)
	if err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, message, data)
}

func destroy(c *fiber.Ctx, remove func(string) error, message, notFoundMessage string) error {
	id := c.Params("id")
	if id == "" {
		return helper.BadRequest(c, "ID is required")
	}

	if err := remove(id); err != nil {
		return respondRepositoryError(c, err, notFoundMessage)
	}

	return helper.APIResponse(c, fiber.StatusOK, message, nil)
}

func parseInput(c *fiber.Ctx) (map[string]any, error) {
	var input map[string]any
	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}
	if input == nil {
		input = map[string]any{}
	}

	// Format ISO 8601 date strings to MySQL format
	autoFormatISODates(input)

	return input, nil
}

func autoFormatISODates(input map[string]any) {
	layouts := []string{
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05.000Z07:00",
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05",
	}

	for k, v := range input {
		str, ok := v.(string)
		if !ok || str == "" {
			continue
		}

		if !strings.Contains(str, "T") {
			continue
		}

		for _, layout := range layouts {
			if t, err := time.Parse(layout, str); err == nil {
				input[k] = t.Format("2006-01-02 15:04:05")
				break
			}
		}
	}
}

func setDefaultTimestamp(input map[string]any, field string, value time.Time) {
	if _, ok := input[field]; !ok {
		input[field] = value
	}
}

func respondRepositoryError(c *fiber.Ctx, err error, notFoundMessage string) error {
	if errors.Is(err, repository.ErrNotFound) {
		return helper.NotFound(c, notFoundMessage)
	}
	return helper.InternalError(c, err.Error())
}
