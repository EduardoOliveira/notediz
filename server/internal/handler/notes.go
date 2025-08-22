package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/EduardoOliveira/notediz/internal/types"
)

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload map[string]any
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	slog.Info("Creating note", "payload", payload)
	switch payload["Kind"].(string) {
	case string(types.NoteKindText):
		txt := types.Text{}
		if err := txt.CreateFromAny(payload); err != nil {
			h.errorResponse(w, http.StatusBadRequest, err)
			return
		}
		txt, err = h.repo.CreateText(r.Context(), txt)
		if err != nil {
			h.errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		h.jsonResponse(w, http.StatusCreated, txt)

	case string(types.NoteKindBookmark):
		bm := types.Bookmark{}
		if err := bm.CreateFromAny(payload); err != nil {
			h.errorResponse(w, http.StatusBadRequest, err)
			return
		}
		bm, err = h.repo.CreateBookmark(r.Context(), bm)
		if err != nil {
			h.errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		h.jsonResponse(w, http.StatusCreated, bm)
	}

}

func (h *Handler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var payload types.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	if err := payload.Validate(); err != nil {
		h.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	bookmark, err := h.repo.CreateBookmark(r.Context(), payload)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	h.jsonResponse(w, http.StatusCreated, bookmark)
}

func (h *Handler) GetNoteContext(w http.ResponseWriter, r *http.Request) {
	noteID := r.URL.Query().Get("note_id")

	context, err := h.repo.GetNoteContext(noteID)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	h.jsonResponse(w, http.StatusOK, context)
}
