package infra

import (
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/dao"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/db/postgresql"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra/token"
	"github.com/google/wire"
)

var Set = wire.NewSet(

	dao.NewAccountRepository,
	dao.NewAuthenticationRepository,

	token.NewTokenManager,

	postgresql.NewClient,
)
