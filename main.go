package main

import (
	"context"
	console "simple-wallet/cmd"

	"simple-wallet/config"
	"simple-wallet/docs"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ukautz/clif.v1"
)

//	@title			Wallet Service
//	@version		1.0.0
//	@description	This is a Wallet Service server.

// @host		api-stg.paper.id
// @schemes	http https
func main() {
	ctx := context.Background()
	cli := clif.New("b2b-company-management-service", "1.0.0", "Company Management Service for FFB")
	cmd, err := console.Init()
	if err != nil {
		log.Fatal("failed init console", err)
	}

	cli.Add(cmd.StartServer())

	cli.Add(cmd.MigrateRun(ctx))
	cli.Add(cmd.MigrateRollback())
	cli.Add(cmd.MigrateReset())
	cli.Add(cmd.MigrateRefresh())

	cfg := config.All()
	docs.SwaggerInfo.Title = cfg.Swagger.Title
	docs.SwaggerInfo.Description = cfg.Swagger.Description
	docs.SwaggerInfo.Version = cfg.Swagger.Version
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	cli.Run()
}
