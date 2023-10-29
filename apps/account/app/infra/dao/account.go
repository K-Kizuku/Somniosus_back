package dao

import (
	"context"
	"database/sql"
	"errors"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/repository"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/db/postgresql"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/dto"
)

type accountRepository struct {
	db *postgresql.Client
}

func NewAccountRepository(db *postgresql.Client) repository.Account {
	return &accountRepository{
		db: db,
	}
}

func (a accountRepository) GetAccount(ctx context.Context, id string) (*account.Account, error) {
	var ac dto.Account
	q := `
		SELECT
			id,
			status,
			display_id,
			name,
			description,
			image_url,
			birth_day,
			tel_num,
			created_at,
			updated_at
		FROM accounts
		WHERE id = ?
		LIMIT 1
	`
	if err := a.db.Do(ctx).GetContext(ctx, &ac, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotExistAccount
		}
		return nil, err
	}

	return account.Assign(ac.Status, ac.ID, ac.DisplayID, ac.Name, ac.Description, ac.ImageURL, ac.TelNum, ac.BirthDay), nil
}

func (a accountRepository) CreateAccount(ctx context.Context, account *account.Account) error {
	q := `
		INSERT INTO 
			accounts
		(
			id,
			status,
			display_id,
			name,
			description,
			image_url,
			birth_day,
			tel_num
		)
		VALUES
		(
			:id,
			:status,
			:display_id,
			:name,
			:description,
			:image_url,
			:birth_day,
			:tel_num
		)

	`
	ac := &dto.Account{
		ID:          account.ID(),
		Status:      account.Status(),
		DisplayID:   account.DisplayID(),
		Name:        account.Name(),
		Description: account.Description(),
		ImageURL:    account.ImageURL(),
		BirthDay:    account.BirthDay(),
		TelNum:      account.TelNum(),
	}
	if _, err := a.db.Do(ctx).NamedExecContext(ctx, q, ac); err != nil {
		return err
	}
	return nil
}

func (a accountRepository) UpdateAccount(ctx context.Context, account *account.Account) error {
	q := `
	UPDATE 
		accounts
	SET
		status=:status,
		display_id=:display_id,
		name=:name,
		description=:description,
		image_url=:image_url,
		birth_day=:birth_day,
		tel_num=:tel_num
	WHERE
		id = :id
	`
	ac := &dto.Account{
		ID:          account.ID(),
		Status:      account.Status(),
		DisplayID:   account.DisplayID(),
		Name:        account.Name(),
		Description: account.Description(),
		ImageURL:    account.ImageURL(),
		BirthDay:    account.BirthDay(),
		TelNum:      account.TelNum(),
	}
	if _, err := a.db.Do(ctx).NamedExecContext(ctx, q, ac); err != nil {
		return err
	}
	return nil
}
