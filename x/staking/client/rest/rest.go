package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx client.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}
