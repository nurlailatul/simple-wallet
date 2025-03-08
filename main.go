package main

import (
	"context"
	console "simple-wallet/cmd"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ukautz/clif.v1"
)

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

	cli.Run()
}
