package db

import (
	"context"
	"log/slog"

	"github.com/EduardoOliveira/notediz/internal/db/gen"
	"github.com/EduardoOliveira/notediz/internal/types"
)

func (r *Repo) CreateText(ctx context.Context, note types.Text) (types.Text, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		return types.Text{}, err
	}
	defer tx.Rollback()
	qtx := r.queries.WithTx(tx)

	noteID := r.uuid()
	txtID := r.uuid()
	err = qtx.CreateNote(ctx, gen.CreateNoteParams{
		ID:        noteID,
		Kind:      string(types.NoteKindText),
		CreatedAt: r.now(),
		UpdatedAt: r.now(),
	})
	if err != nil {
		slog.Error("Error creating note", "error", err, "noteID", noteID)
		return types.Text{}, err
	}

	err = qtx.CreateText(ctx, gen.CreateTextParams{
		ID:        txtID,
		NoteID:    noteID,
		Content:   note.Content,
		CreatedAt: r.now(),
		UpdatedAt: r.now(),
	})
	if err != nil {
		return types.Text{}, err
	}

	rtn, err := qtx.GetTextByID(ctx, txtID)
	if err != nil {
		return types.Text{}, err
	}

	return types.Text(rtn), tx.Commit()
}

func (r *Repo) CreateBookmark(ctx context.Context, bookmark types.Bookmark) (types.Bookmark, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		return types.Bookmark{}, err
	}
	defer tx.Rollback()
	qtx := r.queries.WithTx(tx)

	noteID := r.uuid()
	bookmarkID := r.uuid()
	err = qtx.CreateNote(ctx, gen.CreateNoteParams{
		ID:        noteID,
		Kind:      string(types.NoteKindBookmark),
		CreatedAt: r.now(),
		UpdatedAt: r.now(),
	})
	if err != nil {
		slog.Error("Error creating note for bookmark", "error", err, "noteID", noteID)
		return types.Bookmark{}, err
	}

	err = qtx.CreateBookmark(ctx, gen.CreateBookmarkParams{
		ID:        bookmarkID,
		NoteID:    noteID,
		Url:       bookmark.Url,
		Title:     bookmark.Title,
		CreatedAt: r.now(),
		UpdatedAt: r.now(),
	})
	if err != nil {
		return types.Bookmark{}, err
	}

	rtn, err := qtx.GetBookmarkByID(ctx, bookmarkID)
	if err != nil {
		return types.Bookmark{}, err
	}

	return types.Bookmark(rtn), tx.Commit()
}
