package main

import (
	"flag"

	"github.com/subliker/server/internal/app"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/store/item/gorm"
	"github.com/subliker/server/internal/store/photo/minio"
)

func main() {
	flag.Parse()

	// getting config
	cfg := config.Get()
	logger.Zap.Debugf("Config: %v", cfg)

	// gorm item store init
	itemStore, err := gorm.NewMySQL(cfg.ItemStore)
	if err != nil {
		logger.Zap.Fatalf("creating gorm db error: %s", err)
	}
	defer itemStore.Close()

	// minio photo store init
	photoStore := minio.New(cfg.PhotoStore)

	// running main app
	app.New(cfg.App, itemStore, photoStore).Run()
}
