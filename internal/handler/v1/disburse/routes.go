package disburse

import (
	"net/http"

	"simple-wallet/internal/app/server"

	"github.com/gin-gonic/gin"
)

const (
	WalletPath = "/wallets"
)

func (h *HTTPHandler) RegisterRoute() []server.RouteHandler {
	routes := []server.RouteHandler{
		{
			Method: http.MethodPost,
			Path:   WalletPath + "/:user_id/deduct",
			Handler: []gin.HandlerFunc{
				h.createDisbursement,
			},
		},
	}

	return routes
}
