package exception

import (
	"fiber-ordering/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return c.JSON(model.WebResponse{
			Header: c.Status(fiber.StatusBadRequest),
			Data:   err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Header: c.Status(fiber.StatusInternalServerError),
		Data:   err.Error(),
	})
}