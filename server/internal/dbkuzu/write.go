package dbkuzu

import (
	"context"
	"log/slog"

	"github.com/EduardoOliveira/notediz/internal/types"
)

func (r *Repo) CreateText(ctx context.Context, note types.Text) (types.Text, error) {
	return types.Text{}, nil
}

func (r *Repo) CreateBookmark(ctx context.Context, bookmark types.Bookmark) (types.Bookmark, error) {
	query := `MERGE (b:Bookmark {id: $id})
	SET b.title = $title, b.url = $url
	return b
	`
	// url: $url, title: $title, created_at: $created_at, updated_at: $updated_at
	prep, err := r.Conn.Prepare(query)
	if err != nil {
		return types.Bookmark{}, err
	}
	defer prep.Close()

	result, err := r.Conn.Execute(prep, map[string]any{
		"id":    r.uuid(),
		"title": bookmark.Title,
		"url":   bookmark.URL,
	})
	if err != nil {
		return types.Bookmark{}, err
	}
	defer result.Close()

	var createdBookmark types.Bookmark
	for result.HasNext() {
		row, err := result.Next()
		if err != nil {
			return types.Bookmark{}, err
		}
		slog.Info("Created bookmark", "bookmark", row)
	}

	return createdBookmark, nil
}
