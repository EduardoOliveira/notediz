package dbkuzu

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/EduardoOliveira/notediz/internal/types"
)

func (r *Repo) CreateText(ctx context.Context, note types.Text) (types.Text, error) {
	query := `MERGE (t:Text {ID: $ID})
	ON CREATE SET t.CreatedAt = $CreatedAt
	SET t.Content = $Content, t.Kind = $Kind, t.UpdatedAt = $UpdatedAt
	return t
	`

	prep, err := r.Conn.Prepare(query)
	if err != nil {
		return types.Text{}, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer prep.Close()

	var createdText types.Text
	result, err := r.Conn.Execute(prep, map[string]any{
		"ID":        r.uuid(),
		"Content":   note.Content,
		"Kind":      string(note.Kind),
		"CreatedAt": r.now().UTC(),
		"UpdatedAt": r.now().UTC(),
	})
	if err != nil {
		return types.Text{}, fmt.Errorf("failed to execute query: %v", err)
	}
	defer result.Close()

	if result.HasNext() {
		tuple, err := result.Next()
		if err != nil {
			return types.Text{}, fmt.Errorf("failed to get next tuple: %v", err)
		}
		slog.Info("Created text", "tuple", tuple.GetAsString())

		tm, err := tuple.GetAsMap()
		if err != nil {
			return types.Text{}, fmt.Errorf("failed to get text as map: %v", err)
		}

		var createdText types.Text
		createdText.FromAny(tm)
		tuple.Close()
		return createdText, nil
	}

	return createdText, errors.New("no text created, result is empty")
}

func (r *Repo) CreateBookmark(ctx context.Context, bookmark types.Bookmark) (types.Bookmark, error) {
	query := `MERGE (b:Bookmark {ID: $ID})
	ON CREATE SET b.CreatedAt = $CreatedAt
	SET b.Title = $Title, b.URL = $URL, b.Kind = $Kind, b.UpdatedAt = $UpdatedAt
	return b
	`

	prep, err := r.Conn.Prepare(query)
	if err != nil {
		return types.Bookmark{}, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer prep.Close()

	slog.Info("Creating bookmark", "bookmark", bookmark)
	result, err := r.Conn.Execute(prep, map[string]any{
		"ID":        r.uuid(),
		"Title":     bookmark.Title,
		"URL":       bookmark.URL,
		"Kind":      string(bookmark.Kind),
		"CreatedAt": r.now().UTC(),
		"UpdatedAt": r.now().UTC(),
	})
	if err != nil {
		return types.Bookmark{}, fmt.Errorf("failed to execute query: %v", err)
	}
	defer result.Close()

	var createdBookmark types.Bookmark
	if result.HasNext() {
		tuple, err := result.Next()
		if err != nil {
			return types.Bookmark{}, fmt.Errorf("failed to get next tuple: %v", err)
		}
		slog.Info("Created bookmark", "tuple", tuple.GetAsString())

		tm, err := tuple.GetAsMap()
		if err != nil {
			return types.Bookmark{}, fmt.Errorf("failed to get bookmark as map: %v", err)
		}
		var createdBookmark types.Bookmark
		createdBookmark.FromAny(tm)
		tuple.Close()
		return createdBookmark, nil
	}

	return createdBookmark, errors.New("no bookmark created, result is empty")
}
