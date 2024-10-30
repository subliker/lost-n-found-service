package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/server/internal/config"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/store/item"
	"github.com/subliker/server/internal/store/photo"
)

type Server struct {
	config     config.Server
	router     *mux.Router
	itemStore  item.Store
	photoStore photo.Store
}

// New creates new instance of server with params from cfg
func New(cfg config.Server, itemStore item.Store, photoStore photo.Store) *Server {
	s := Server{
		config:     cfg,
		router:     mux.NewRouter(),
		itemStore:  itemStore,
		photoStore: photoStore,
	}
	s.initRoutes()

	logger.Zap.Info("Server instance created")
	return &s
}

// Run runs server listening
func (s *Server) Run() {
	logger.Zap.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router))
}
