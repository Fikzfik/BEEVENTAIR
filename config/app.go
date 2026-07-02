package config

import (
	f "github.com/gofiber/fiber/v2"
)

func NewApp() *f.App {
	return f.New(f.Config{
		AppName: "EventNear API",
	})
}
