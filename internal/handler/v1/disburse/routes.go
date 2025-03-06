package disburse

import (
	"net/http"

	"simple-wallet/internal/app/server"

	"github.com/gin-gonic/gin"
)

const (
	DisbursePath = "/disburse"
)

func (h *HTTPHandler) RegisterRoute() []server.RouteHandler {
	routes := []server.RouteHandler{
		{
			Method: http.MethodPost,
			Path:   DisbursePath,
			Handler: []gin.HandlerFunc{
				h.getIdentityImageData,
			},
		},
		{
			Method: http.MethodGet,
			Path:   DisbursePath,
			Handler: []gin.HandlerFunc{
				h.testHTTP,
			},
		},
	}

	return routes
}
