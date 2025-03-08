package app

import (
	"net/http"
	"simple-wallet/config"
	"simple-wallet/internal/app/server"
	disburseHandler "simple-wallet/internal/handler/v1/disburse"
	"simple-wallet/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *AppService) SetupHttpRouteHandler(cfg *config.Configuration) server.Route {
	v1Routes := a.setupV1Handler()
	otherRoutes := a.setupOtherHandler(a.SwaggerAuthMiddleware)

	return server.Route{
		V1:    v1Routes,
		Other: otherRoutes,
	}
}

func (a *AppService) setupV1Handler() []server.RouteHandler {
	routes := make([]server.RouteHandler, 0)
	routes = append(routes, disburseHandler.NewDisburseHandler(a.TransactionService, a.UserService, a.WalletService).RegisterRoute()...)

	return routes
}

func (a *AppService) setupOtherHandler(auth *middleware.SwaggerAuthMiddleware) []server.RouteHandler {
	pingHandler := server.RouteHandler{
		Method: http.MethodGet,
		Path:   "/ping",
		Handler: []gin.HandlerFunc{
			func(c *gin.Context) {
				c.JSON(http.StatusOK, "ok")
			},
		},
	}

	swaggerHandler := server.RouteHandler{
		Method: http.MethodGet,
		Path:   "/swagger/*any",
		Handler: []gin.HandlerFunc{
			auth.VerifyHeaderKey(),
			ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)),
		},
	}

	return []server.RouteHandler{
		pingHandler,
		swaggerHandler,
	}
}
