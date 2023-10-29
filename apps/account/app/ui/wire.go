package ui

import (
	"github.com/K-Kizuku/Somniosus_back/apps/account/app/ui/grpc/service"
	"github.com/google/wire"
)

// Set : ui層のwire set
var Set = wire.NewSet(
	service.NewService,
)
