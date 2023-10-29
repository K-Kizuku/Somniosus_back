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

type authenticationRepository struct {
	db *postgresql.Client
}

func NewAuthenticationRepository(db *postgresql.Client) repository.Authentication {
	return &authenticationRepository{
		db: db,
	}
}

func (a authenticationRepository) GetAuthentication(ctx context.Context, id string) (string, error) {
	var hashedPass string
	q := `
		SELECT 
			hashed_password
		FROM
			authentications
		WHERE
			id = ?
	`
	if err := a.db.Do(ctx).GetContext(ctx, &hashedPass, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.ErrNotExistAuthentication
		}
		return "", err
	}
	return hashedPass, nil
}

func (a authenticationRepository) CreateAuthentication(ctx context.Context, auth *account.Authentication) error {
	q := `
		INSERT INTO 
			authentications
		(
			id,
			hashed_password
		)
		VALUES
		(
			:id,
			:hashed_password
		)
	`
	au := &dto.Authentication{
		ID:             auth.ID(),
		HashedPassword: auth.HashedPassword(),
	}
	if _, err := a.db.Do(ctx).NamedExecContext(ctx, q, au); err != nil {
		return err
	}
	return nil
}
