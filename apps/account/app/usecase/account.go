package usecase

import (
	"context"
	"time"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/repository"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/token"
	"github.com/K-Kizuku/Somniosus_back/lib/hash"
	"github.com/K-Kizuku/Somniosus_back/lib/ulid"
)

type (
	accountUsecase struct {
		accountRepo repository.Account
		authRepo    repository.Authentication
		tokenRepo   token.TokenManager
	}
	AccountUsecase interface {
		CreateTempAccount(ctx context.Context, input *CreateTempAccountInput) (string, error)
		CreateAccount(ctx context.Context, jwt, displayID, password string) (string, error)
		GetAccount(ctx context.Context, id string) (*account.Account, error)
		UpdateAccount(ctx context.Context, ac *account.Account) error
		VerifyToken(ctx context.Context, jwt string) (string, error)
		GenerateToken(ctx context.Context, id, password string, status account.AccountStatus) (string, error)
	}
)

func NewAccount(accountRepo repository.Account, authentication repository.Authentication, tokenRepo token.TokenManager) AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
		authRepo:    authentication,
		tokenRepo:   tokenRepo,
	}
}

type CreateTempAccountInput struct {
	Name     string
	TelNum   string
	BirthDay time.Time
}

func (a *accountUsecase) CreateTempAccount(ctx context.Context, input *CreateTempAccountInput) (string, error) {
	token, err := a.tokenRepo.GenerateTempToken(ctx, ulid.GenerateID(), input.Name, input.TelNum, input.BirthDay, time.Now().Add(time.Duration(token.ExpiresTime)*time.Second))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *accountUsecase) CreateAccount(ctx context.Context, jwt, displayID, password string) (string, error) {
	claims, err := a.tokenRepo.VerifyTempToken(ctx, jwt)
	if err != nil {
		return "", err
	}
	ac := account.Assign(account.General, claims.ID, displayID, claims.Name, "", "", claims.TelNum, claims.BirthDay)
	if err := a.accountRepo.CreateAccount(ctx, ac); err != nil {
		return "", err
	}
	hashedPass, err := hash.Generate(password)
	if err != nil {
		return "", err
	}
	au := account.NewAuthentication(ac.ID(), hashedPass)
	if err := a.authRepo.CreateAuthentication(ctx, au); err != nil {
		return "", nil
	}
	newJwt, err := a.tokenRepo.GenerateToken(ctx, ac.ID(), time.Now().Add(time.Duration(token.ExpiresTime)*time.Second), ac.Status())
	if err != nil {
		return "", err
	}
	return newJwt, nil
}

func (a *accountUsecase) GetAccount(ctx context.Context, id string) (*account.Account, error) {
	ac, err := a.accountRepo.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (a *accountUsecase) UpdateAccount(ctx context.Context, ac *account.Account) error {
	if err := a.accountRepo.UpdateAccount(ctx, ac); err != nil {
		return err
	}
	return nil
}

func (a *accountUsecase) VerifyToken(ctx context.Context, jwt string) (string, error) {
	claims, err := a.tokenRepo.VerifyToken(ctx, jwt)
	if err != nil {
		return "", err
	}
	return claims.ID, nil
}

func (a *accountUsecase) GenerateToken(ctx context.Context, id, password string, status account.AccountStatus) (string, error) {
	hashedPass, err := a.authRepo.GetAuthentication(ctx, id)
	if err != nil {
		return "", err
	}
	if err := hash.Compare(hashedPass, password); err != nil {
		return "", err
	}
	jwt, err := a.tokenRepo.GenerateToken(ctx, id, time.Now().Add(time.Duration(token.ExpiresTime)*time.Second), status)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
