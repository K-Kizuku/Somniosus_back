package usecase

import (
	"github.com/google/wire"
)

// Set : UseCase層のwire set
var Set = wire.NewSet(
	NewAccount,
)
