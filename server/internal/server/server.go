package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/storage"
	"gorm.io/gorm"
)

type Server struct {
	config  config.ServerConfig
	router  *mux.Router
	db      *gorm.DB
	storage storage.Storage
}

// New creates new instance of server with params from cfg
func New(cfg config.ServerConfig, db *gorm.DB, storage storage.Storage) *Server {
	s := Server{
		config:  cfg,
		router:  mux.NewRouter(),
		db:      db,
		storage: storage,
	}
	s.initRoutes()

	logger.Zap.Info("Server instance created")
	return &s
}

// Run runs server listening
func (s *Server) Run() {
	logger.Zap.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router))
}
