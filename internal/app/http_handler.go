package app

import (
	"simple-wallet/config"
	"simple-wallet/internal/app/server"
	disburseHandler "simple-wallet/internal/handler/v1/disburse"
)

func (a *AppService) SetupHttpRouteHandler(cfg *config.Configuration) server.Route {
	v1Routes := a.setupV1Handler()

	return server.Route{
		V1: v1Routes,
	}
}

func (a *AppService) setupV1Handler() []server.RouteHandler {
	routes := make([]server.RouteHandler, 0)
	routes = append(routes, disburseHandler.NewDisburseHandler(a.UserService).RegisterRoute()...)

	return routes
}
