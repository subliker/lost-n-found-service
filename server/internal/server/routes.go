package server

import "github.com/subliker/server/internal/logger"

func (s *Server) initRoutes() {
	if s.router == nil {
		logger.Zap.Fatalf("routes init error: router wasn't initialized")
	}

	s.router.HandleFunc("/items", s.handleGetItems()).Methods("GET")
	s.router.HandleFunc("/item", s.handleCreateItem()).Methods("POST")

	logger.Zap.Info("Server routes was initialized")
}
