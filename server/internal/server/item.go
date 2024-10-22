package server

import (
	"encoding/json"
	"net/http"

	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/models"
)

func (s *Server) getItems(w http.ResponseWriter, r *http.Request) {
	items := make([]models.Item, 0)
	if err := s.db.Find(&items).Error; err != nil {
		http.Error(w, "Error getting items", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		logger.Zap.Errorf("error encoding items: %s", err)
	}
}

func (s *Server) createItem(w http.ResponseWriter, r *http.Request) {
	var i models.Item
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "Error decoding item data: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// var err error
	// if i.FoundTime, err = time.Parse(time.RFC3339, r.FormValue("found_time")); err != nil {
	// 	http.Error(w, "Unable parse time "+, http.StatusInternalServerError)
	// 	return
	// }
	logger.Zap.Info(i)

	if err := s.db.Create(&i); err != nil {
		http.Error(w, "Error creating item ", http.StatusInternalServerError)
		return
	}
}
