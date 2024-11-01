package app

import (
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/server"
	"github.com/subliker/server/internal/store/item"
	"github.com/subliker/server/internal/store/photo"
)

type App struct {
	server *server.Server
}

// New creates instance of app with router and bd
func New(cfg config.Config, itemStore item.Store, photoStore photo.Store) *App {
	a := &App{
		server: server.New(cfg.Server, itemStore, photoStore),
	}
	logger.Zap.Info("App instance created")
	return a
}

// Run starts initialized app
func (a *App) Run() {
	logger.Zap.Info("App running...")
	a.server.Run()
}
