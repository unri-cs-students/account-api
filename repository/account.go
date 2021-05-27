package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	"fiber-ordering/model"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

const TableAccount = "account"

// accountRepo for accountRepo struct
type accountRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewAccountRepo for creates AccountRepository impl
func NewAccountRepo(reader, writer *sql.DB) entity.AccountRepository {
	return &accountRepo{Reader: reader, Writer: writer}
}

func (r *accountRepo) CheckUsername(username string) bool {
	query := sq.Select("*").
		From(TableAccount).
		Where(sq.Eq{
			"username": username,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return false
	}
	defer rows.Close()

	var acc entity.Account
	for rows.Next() {
		err = rows.Scan(
			&acc.AccountID,
			&acc.FullName,
			&acc.Email,
			&acc.PhoneNumber,
			&acc.Username,
			&acc.Password,
			&acc.CreatedAt,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
	}

	if acc.AccountID == "" {
		return false
	}
	return true
}

func (r *accountRepo) Insert(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	query := sq.Insert(TableAccount).
		Columns(
			"account_id",
			"full_name",
			"email",
			"phone_number",
			"username",
			"password",
			"created_at",
		).
		Values(
			account.AccountID,
			account.FullName,
			account.Email,
			account.PhoneNumber,
			account.Username,
			account.Password,
			account.CreatedAt,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	_, err = r.Writer.ExecContext(ctx, sqlInsert, argsInsert...)
	if err != nil {
		return &entity.Account{}, err
	}

	return account, nil
}

func (r *accountRepo) GetAll() (res []entity.Account, err error) {
	query := sq.Select("*").
		From(TableAccount).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Account
		err = rows.Scan(
			&r.AccountID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.CreatedAt,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (r *accountRepo) GetByID(AccountID string) (res *entity.Account, err error) {
	query := sq.Select("*").
		From(TableAccount).
		Where(sq.Eq{
			"account_id": AccountID,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	var rows *sql.Rows
	var errQuery error
	rows, errQuery = query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Account
		err = rows.Scan(
			&r.AccountID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.CreatedAt,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	err = errQuery

	return
}

func (r *accountRepo) UpdateByID(ctx context.Context, accountID string, account *model.UpdateAccountRequest) error {
	query := sq.Update(TableAccount).
		Where(sq.Eq{
			"account_id": accountID,
		}).
		SetMap(map[string]interface{}{
			"full_name": account.FullName,
			"email": account.Email,
			"phone_number": account.PhoneNumber,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepo) DeleteByID(ctx context.Context, AccountID string) error {
	query := sq.Delete("").
		From(TableAccount).
		Where(sq.Eq{
			"account_id": AccountID,
		}).
		RunWith(r.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}