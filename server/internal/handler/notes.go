package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/EduardoOliveira/notediz/internal/types"
)

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	r.Body = io.NopCloser(io.TeeReader(r.Body, &b))

	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	if payload["Content"] == nil {
		h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("Content is required"))
		return
	}

	if v, ok := payload["kind"].(string); !ok || v == "" {
		if isValidURL(payload["Content"].(string)) {
			bm, err := h.repo.CreateBookmark(r.Context(), types.BookmarkFromAny(payload))
			if err != nil {
				h.errorResponse(w, http.StatusInternalServerError, err)
				return
			}
			h.jsonResponse(w, http.StatusCreated, bm)
			return
		} else {
			txt, err := h.repo.CreateText(r.Context(), types.Text{
				Content: payload["Content"].(string),
			})
			if err != nil {
				h.errorResponse(w, http.StatusInternalServerError, err)
				return
			}
			h.jsonResponse(w, http.StatusCreated, txt)
			return
		}
	}

	var err error
	switch payload["kind"].(types.NoteKind) {
	case types.NoteKindText:
		var txt types.Text
		if err := json.Unmarshal(b.Bytes(), &txt); err != nil {
			h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
			return
		}

		txt, err = h.repo.CreateText(r.Context(), txt)
		if err != nil {
			h.errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		h.jsonResponse(w, http.StatusCreated, txt)
	case types.NoteKindBookmark:
		var bm types.Bookmark
		if err := json.Unmarshal(b.Bytes(), &bm); err != nil {
			h.errorResponse(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
			return
		}

		bm, err = h.repo.CreateBookmark(r.Context(), bm)
		if err != nil {
			h.errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		h.jsonResponse(w, http.StatusCreated, bm)
		// other types
	}

}

func isValidURL(url string) bool {
	// A simple URL validation logic, can be improved
	if len(url) < 5 || (len(url) > 2048 && len(url) < 10) {
		return false
	}
	if url[:4] != "http" && url[:3] != "www" {
		return false
	}
	return true
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
