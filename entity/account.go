package entity

import (
	"context"
	"fiber-ordering/model"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Account for account table
type Account struct {
	AccountID     string
	FullName      string
	Email         string
	PhoneNumber   string
	Username      string
	Password      string
	CreatedAt 	  time.Time
}

// AccountRepository for account repo
type AccountRepository interface {
	Insert(ctx context.Context, account *Account) (*Account, error)
	GetAll() ([]Account, error)
	GetByID(AccountID string) (*Account, error)
	UpdateByID(ctx context.Context, accountID string, account *model.UpdateAccountRequest) error
	DeleteByID(ctx context.Context, AccountID string) error
	CheckUsername(username string) bool
}

// AccountService for service
type AccountService interface {
	CreateAccount(c *fiber.Ctx, account *model.AccountRequest) (*model.AccountResponse, error)
	GetAll() ([]model.AccountResponse, error)
	GetByID(AccountID string) (*model.AccountResponse, error)
	UpdateAccountByID(c *fiber.Ctx, accountID string, account *model.UpdateAccountRequest) error
	DeleteByID(c *fiber.Ctx, AccountID string) error
}