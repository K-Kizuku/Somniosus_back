package repository

import (
	"context"
	"errors"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
)

var (
	ErrNotExistAccount = errors.New("アカウントが存在しません")
)

type Account interface {
	CreateAccount(ctx context.Context, account *account.Account) error
	GetAccount(ctx context.Context, id string) (*account.Account, error)
	UpdateAccount(ctx context.Context, account *account.Account) error
}
