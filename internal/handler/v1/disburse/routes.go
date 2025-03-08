package disburse

import (
	"net/http"

	"simple-wallet/internal/app/server"
	"simple-wallet/internal/middleware"

	"github.com/gin-gonic/gin"
)

const (
	WalletPath = "/wallets"
)

func (h *HTTPHandler) RegisterRoute(auth *middleware.AuthMiddleware) []server.RouteHandler {
	routes := []server.RouteHandler{
		{
			Method: http.MethodPost,
			Path:   WalletPath + "/:user_id/deduct",
			Handler: []gin.HandlerFunc{
				auth.VerifyHeaderKey(),
				h.createDisbursement,
			},
		},
	}

	return routes
}
