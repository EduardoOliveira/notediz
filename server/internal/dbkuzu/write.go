package dbkuzu

import (
	"context"
	"errors"
	"log/slog"

	"github.com/EduardoOliveira/notediz/internal/types"
)

func (r *Repo) CreateText(ctx context.Context, note types.Text) (types.Text, error) {
	return types.Text{}, nil
}

func (r *Repo) CreateBookmark(ctx context.Context, bookmark types.Bookmark) (types.Bookmark, error) {
	query := `MERGE (b:Bookmark {id: $id})
	ON CREATE SET b.created_at = $created_at
	ON MATCH SET b.updated_at = $updated_at
	SET b.title = $title, b.url = $url, b.kind = $kind
	return b
	`

	prep, err := r.Conn.Prepare(query)
	if err != nil {
		return types.Bookmark{}, err
	}
	defer prep.Close()

	slog.Info("Creating bookmark", "bookmark", bookmark)
	result, err := r.Conn.Execute(prep, map[string]any{
		"id":         r.uuid(),
		"title":      bookmark.Title,
		"url":        bookmark.URL,
		"kind":       string(bookmark.Kind),
		"created_at": r.now().UTC(),
		"updated_at": r.now().UTC(),
	})
	if err != nil {
		return types.Bookmark{}, err
	}
	defer result.Close()

	var createdBookmark types.Bookmark
	if result.HasNext() {
		tuple, err := result.Next()
		if err != nil {
			return types.Bookmark{}, err
		}
		slog.Info("Created bookmark", "tuple", tuple.GetAsString())

		tm, err := tuple.GetAsMap()
		if err != nil {
			return types.Bookmark{}, err
		}

		createdBookmark, err = types.BookmarkFromFlatTuple(tm)
		if err != nil {
			return types.Bookmark{}, err
		}
		tuple.Close()
		return createdBookmark, nil
	}

	return createdBookmark, errors.New("no bookmark created, result is empty")
}
