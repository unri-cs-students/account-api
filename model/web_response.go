package model

import "github.com/gofiber/fiber/v2"

type WebResponse struct {
	Header *fiber.Ctx `json:"-"`
	Data   interface{} `json:"data"`
}