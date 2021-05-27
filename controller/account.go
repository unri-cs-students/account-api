package controller

import (
	"fiber-ordering/entity"
	"fiber-ordering/exception"
	"fiber-ordering/model"
	"github.com/gofiber/fiber/v2"
)

type accountAPI struct {
	accountService entity.AccountService
}

// NewAccountAPI will initiate new account api
func NewAccountAPI(app *fiber.App, accountService entity.AccountService) {
	accountAPI := &accountAPI{
		accountService: accountService,
	}

	api := app.Group("/api/v1")
	api.Post("/account", accountAPI.CreateAccount)
	api.Get("/account", accountAPI.GetAllAccount)
	api.Get("/account/:account_id", accountAPI.GetAccountByID)
	api.Put("/account/:account_id", accountAPI.UpdateAccountByID)
	api.Delete("/account/:account_id", accountAPI.DeleteAccountByID)

}

func (api *accountAPI) CreateAccount(c *fiber.Ctx) error {
	var request model.AccountRequest
	exception.PanicIfNeeded(c.BodyParser(&request))

	response, _ := api.accountService.CreateAccount(c, &request)
	return c.JSON(model.WebResponse{
		Header: c.Status(fiber.StatusCreated),
		Data:   response,
	})
}

func (api *accountAPI) GetAccountByID(c *fiber.Ctx) error {
	accountID := c.Params("account_id")
	result, err := api.accountService.GetByID(accountID)

	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "failed to get account by account by id",
		})
	}

	return c.JSON(model.WebResponse{
		Header: c.Status(fiber.StatusOK),
		Data:   result,
	})
}

func (api *accountAPI) GetAllAccount(c *fiber.Ctx) error {
	result, err := api.accountService.GetAll()
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "failed to get all account",
		})
	}

	return c.JSON(model.WebResponse{
		Header: c.Status(fiber.StatusOK),
		Data:   result,
	})
}

func (api *accountAPI) DeleteAccountByID(c *fiber.Ctx) error {
	accountID := c.Params("account_id")
	err := api.accountService.DeleteByID(c, accountID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return nil
}

func (api *accountAPI) UpdateAccountByID(c *fiber.Ctx) error {
	accountID := c.Params("account_id")

	var body model.UpdateAccountRequest

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := api.accountService.UpdateAccountByID(c, accountID, &body)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.JSON(fiber.StatusOK)
	return nil
}