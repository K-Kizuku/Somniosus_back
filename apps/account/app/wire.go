//go:build wireinject
// +build wireinject

package app

import (
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/infra"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/ui"
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/usecase"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	ui.Set,
	usecase.Set,
	infra.Set,
)

func NewApp() (*App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		superSet,
	)
	return nil, nil
}
