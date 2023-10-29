package app

import (
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/repository"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/db/postgresql"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/usecase"
	"github.com/K-Kizuku/Somniosus_back/proto.pb/twitter/account"
)

// App : Provider
type App struct {
	// gRPC Service
	AccountService account.AccountServiceServer

	// UseCase
	SessionUseCase usecase.AccountUsecase

	// Repository
	AccountRepository repository.Account

	// infra
	DB *postgresql.Client
}
