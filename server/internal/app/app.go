package app

import (
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/server"
	"gorm.io/gorm"
)

type App struct {
	config config.AppConfig
	server *server.Server
	db     *gorm.DB
}

// New creates instance of app with router and bd
func New(cfg config.AppConfig, db *gorm.DB) *App {
	a := &App{
		config: cfg,
		server: server.New(cfg.Server, db),
		db:     db,
	}
	logger.Zap.Info("App instance created")
	return a
}

// Run starts initialized app
func (a *App) Run() {
	a.server.Run()
}
