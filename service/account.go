package service

import (
	"context"
	"errors"
	"fiber-ordering/entity"
	"fiber-ordering/model"
	"fiber-ordering/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type accountService struct {
	AccountRepo entity.AccountRepository
}

// NewAccountService will initiate new service
func NewAccountService(ar entity.AccountRepository) entity.AccountService {
	return &accountService{
		AccountRepo: ar,
	}
}

func (a *accountService) CreateAccount(c *fiber.Ctx, account *model.AccountRequest) (*model.AccountResponse, error) {
	util.Validate(account)

	// check username in DB
	isUsernameAlreadyExistsInDB := a.AccountRepo.CheckUsername(account.Username)
	if isUsernameAlreadyExistsInDB == true {
		return &model.AccountResponse{}, errors.New("username already exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// password hashing
	hashedPassword, _ := util.HashPassword(account.Password)
	acc := entity.Account{
		AccountID:   uuid.New().String(),
		FullName:    account.FullName,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
		Username:    account.Username,
		Password:    hashedPassword,
		CreatedAt:   time.Now(),
	}

	_, err := a.AccountRepo.Insert(ctx, &acc)
	if err != nil {
		return &model.AccountResponse{}, err
	}

	return &model.AccountResponse{
		FullName:    acc.FullName,
		Email:       acc.Email,
		PhoneNumber: acc.PhoneNumber,
		Username:    acc.Username,
		CreatedAt:   acc.CreatedAt,
	}, nil
}

func (a *accountService) GetAll() (res []model.AccountResponse, err error) {
	allAccounts, err := a.AccountRepo.GetAll()

	for _, v := range allAccounts {
		res = append(res, model.AccountResponse{
			FullName:    v.FullName,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			Username:    v.Username,
			CreatedAt:   v.CreatedAt,
		})
	}
	return
}

func (a *accountService) GetByID(AccountID string) (*model.AccountResponse, error) {
	account, err := a.AccountRepo.GetByID(AccountID)

	res := &model.AccountResponse{
		FullName:    account.FullName,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
		Username:    account.Username,
		CreatedAt:   account.CreatedAt,
	}

	return res, err
}

func (a *accountService) UpdateAccountByID(c *fiber.Ctx, accountID string, account *model.UpdateAccountRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return a.AccountRepo.UpdateByID(ctx, accountID, account)
}

func (a *accountService) DeleteByID(c *fiber.Ctx, AccountID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.AccountRepo.DeleteByID(ctx, AccountID)
	return err
}