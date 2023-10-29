package service

import (
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/usecase"
	"github.com/K-Kizuku/Somniosus_back/proto.pb/twitter/account"
)

type Service struct {
	au usecase.AccountUsecase
}

func NewService(au usecase.AccountUsecase) account.AccountServiceServer {
	return &Service{
		au: au,
	}
}
