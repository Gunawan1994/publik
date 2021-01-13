package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver

	"alfa/internal/config"

	thnHdl "alfa/internal/delivery/http/alfa"
	thnRes "alfa/internal/resource/alfa"
	thnSvc "alfa/internal/service/alfa"
)

// Do initialization and injection on this function
func startApp() error {
	config.Init()
	cfg := config.Get()

	db, err := sqlx.Open(cfg.DB.Driver, cfg.DB.Master)
	if err != nil {
		return err
	}

	thnRes := thnRes.New(db)
	thnSvc := thnSvc.New(thnRes)
	thnHdl := thnHdl.New(thnSvc)
	r := newRouter(
		thnHdl,
	)
	return startServer(r, cfg.Server.HTTPPort)
}
