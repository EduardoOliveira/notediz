package types

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"
)

type Bookmark struct {
	ID          string
	URL         string
	Title       string
	Description string
	Kind        NoteKind
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b Bookmark) Validate() error {
	var err error

	if b.Kind != NoteKindBookmark {
		err = errors.Join(err, fmt.Errorf("invalid kind: %s", b.Kind))
	}

	if b.URL == "" {
		slog.Error("Empty URL", "bookmark", b)
		err = errors.Join(err, fmt.Errorf("url cannot be empty"))
	} else if _, vErr := url.Parse(b.URL); vErr != nil {
		err = errors.Join(err, fmt.Errorf("error validating url: %v", vErr))
	}
	return err
}

func (b *Bookmark) FromAny(v map[string]any) {
	if id, ok := v["id"].(string); ok {
		b.ID = id
	}
	if url, ok := v["url"].(string); ok {
		b.URL = url
	}
	if title, ok := v["title"].(string); ok {
		b.Title = title
	}
	if description, ok := v["description"].(string); ok {
		b.Description = description
	}
	anyToTime(&b.CreatedAt, v["created_at"])
	anyToTime(&b.UpdatedAt, v["updated_at"])
}

func BookmarkFromAny(v map[string]any) Bookmark {
	var b Bookmark
	anyToSome(&b.ID, v["id"])
	anyToSome(&b.URL, v["url"])
	anyToSome(&b.Title, v["title"])
	anyToSome(&b.Description, v["description"])
	anyToTime(&b.CreatedAt, v["created_at"])
	anyToTime(&b.UpdatedAt, v["updated_at"])
	b.Kind = NoteKindBookmark
	return b
}

func BookmarkFromFlatTuple(rowMap map[string]any) (Bookmark, error) {
	b := Bookmark{
		Kind: NoteKindBookmark,
	}

	if kind, ok := rowMap["kind"].(string); ok {
		if kind != string(NoteKindBookmark) {
			return b, fmt.Errorf("unexpected kind: %s", kind)
		}
	}
	if id, ok := rowMap["id"].(string); ok {
		b.ID = id
	}
	if url, ok := rowMap["url"].(string); ok {
		b.URL = url
	}
	if title, ok := rowMap["title"].(string); ok {
		b.Title = title
	}
	if description, ok := rowMap["description"].(string); ok {
		b.Description = description
	}
	if createdAt, ok := rowMap["created_at"].(time.Time); ok {
		b.CreatedAt = createdAt
	}
	if updatedAt, ok := rowMap["updated_at"].(time.Time); ok {
		b.UpdatedAt = updatedAt
	}

	return b, nil
}
