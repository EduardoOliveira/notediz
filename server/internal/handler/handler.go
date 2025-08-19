package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/EduardoOliveira/notediz/internal/types"
)

type Repo interface {
	CreateBookmark(ctx context.Context, bookmark types.Bookmark) (types.Bookmark, error)
	CreateText(ctx context.Context, note types.Text) (types.Text, error)
}

type Handler struct {
	repo        Repo
	HTTPHandler http.Handler
}

func New(repo Repo) *Handler {
	h := &Handler{
		repo: repo,
	}
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("POST /api/notes", h.CreateNote)
	httpHandler.HandleFunc("POST /api/bookmarks", h.CreateBookmark)
	h.HTTPHandler = httpHandler
	return h
}

func (_ Handler) jsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func (_ Handler) errorResponse(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]string{"error": err.Error()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}

func (_ *Handler) noContentResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
