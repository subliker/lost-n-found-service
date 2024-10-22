package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"gorm.io/gorm"
)

type Server struct {
	config config.ServerConfig
	router *mux.Router
	db     *gorm.DB
}

// New creates new instance of server with params from cfg
func New(cfg config.ServerConfig, db *gorm.DB) *Server {
	s := Server{
		config: cfg,
		router: mux.NewRouter(),
		db:     db,
	}
	s.initRoutes()

	logger.Zap.Info("Server instance created")
	return &s
}

// Run runs server listening
func (s *Server) Run() {
	logger.Zap.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router))
}
