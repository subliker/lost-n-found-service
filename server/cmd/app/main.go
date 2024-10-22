package main

import (
	"flag"

	"github.com/subliker/server/internal/app"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/db"
	"github.com/subliker/server/internal/logger"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()

	// getting config
	cfg := config.Get()
	logger.Zap.Debugf("Config: %v", cfg)

	// getting gorm
	gdb, err := db.NewMySQL(cfg.DB, &gorm.Config{})
	if err != nil {
		logger.Zap.Fatalf("creating gorm db error: %s", err)
	}

	// getting sql db to defer close
	sqlDB, err := gdb.DB()
	if err != nil {
		logger.Zap.Fatalf("getting sql db error: %s", err)
	}
	defer sqlDB.Close()

	// db migration
	if db.MigrateMode {
		db.Migrate(gdb)
	}

	// running main app
	app.New(cfg.App, gdb).Run()
}
