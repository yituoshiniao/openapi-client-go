package task

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewEnter,
	NewCronSingleAppleProductPriceTask,
)
