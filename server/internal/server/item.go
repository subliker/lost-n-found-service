package server

import (
	"encoding/json"
	"flag"
	"net/http"
	"strconv"

	"github.com/go-playground/form/v4"
	"github.com/subliker/server/internal/logger"
	"github.com/subliker/server/internal/models"
)

var decoder = form.NewDecoder()

var maxMultipartFormSize int64

func init() {
	flag.Int64Var(&maxMultipartFormSize, "mmfs", 32<<20, " setting maximum multipart form size")
}

func (s *Server) handleGetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		page := -1
		if query.Has("page") {
			var err error
			page, err = strconv.Atoi(query.Get("page"))
			if err != nil || page < 1 {
				http.Error(w, "Invalid page number", http.StatusBadRequest)
				return
			}
		}

		pageSize := -1
		if query.Has("page_size") {
			var err error
			pageSize, err = strconv.Atoi(query.Get("page_size"))
			if err != nil || pageSize < 1 {
				http.Error(w, "Invalid page size", http.StatusBadRequest)
				return
			}
		}

		// making and getting items array
		items := make([]models.Item, 0)

		if page > 0 && pageSize > 0 {
			if err := s.itemStore.FindWithPagination(&items, page, pageSize); err != nil {
				http.Error(w, "Error getting items with pagination", http.StatusInternalServerError)
				return
			}
		} else {
			if err := s.itemStore.Find(&items); err != nil {
				http.Error(w, "Error getting items", http.StatusInternalServerError)
				return
			}
		}

		if err := json.NewEncoder(w).Encode(items); err != nil {
			logger.Zap.Errorf("error encoding items: %s", err)
		}
	}
}

func (s *Server) handleCreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parsing multipart form from request
		err := r.ParseMultipartForm(maxMultipartFormSize)
		if err != nil {
			http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// making item struct
		var item models.Item
		if err := decoder.Decode(&item, r.MultipartForm.Value); err != nil {
			http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		logger.Zap.Debug(item)

		// getting photo
		photo, photoHeader, err := r.FormFile("photo_content")
		// if no getting photo from form errors
		if err == nil {
			// getting photo file name in storage
			photoFileName, err := s.photoStore.Put(photo, photoHeader.Filename, photoHeader.Size)
			if err != nil {
				http.Error(w, "Error writing photo: "+err.Error(), http.StatusInternalServerError)
				return
			}
			// adding photo file name into item struct
			item.PhotoFileName = photoFileName
		} else if err != http.ErrMissingFile {
			http.Error(w, "Error getting file from form: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.itemStore.Create(&item); err != nil {
			http.Error(w, "Error creating item: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
