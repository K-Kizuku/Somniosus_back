package repository

import (
	"context"
	"errors"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
)

var (
	ErrNotExistAuthentication = errors.New("認証情報が存在しません")
)

type Authentication interface {
	GetAuthentication(ctx context.Context, id string) (string, error)
	CreateAuthentication(ctx context.Context, auth *account.Authentication) error
}
