package keeper

import (
	"newsletter/x/newsletter/types"
)

var _ types.QueryServer = Keeper{}
