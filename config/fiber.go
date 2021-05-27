package config

import (
	"fiber-ordering/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config  {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}