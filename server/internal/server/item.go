package server

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

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

type CreateItemRequest struct {
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	FoundTime    time.Time `json:"found_time"`
	PhotoContent string    `json:"photo_content"`
}

func (s *Server) createItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// body, _ := io.ReadAll(r.Body)

	var reqBody CreateItemRequest
	// if err := json.Unmarshal(body, &reqBody); err != nil {
	// 	http.Error(w, "Error decoding item data: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	ft, _ := time.Parse(time.RFC3339, r.FormValue("found_time"))

	reqBody = CreateItemRequest{
		Name:      r.FormValue("name"),
		Location:  r.FormValue("location"),
		FoundTime: ft,
	}

	defer r.Body.Close()
	logger.Zap.Info(reqBody)

	item := models.Item{
		Name:      reqBody.Name,
		Location:  reqBody.Location,
		FoundTime: reqBody.FoundTime,
	}

	photo, _, _ := r.FormFile("photo_content")

	if photo != nil {
		photoContent, _ := io.ReadAll(photo)
		fName, err := s.storage.PutPublicPhoto(string(photoContent))
		if err != nil {
			http.Error(w, "Error writing phto: "+err.Error(), http.StatusInternalServerError)
			return
		}
		item.PhotoFileName = fName
	}

	if err := s.db.Create(&item); err.Error != nil {
		http.Error(w, "Error creating item: "+err.Error.Error(), http.StatusInternalServerError)
		return
	}
}
