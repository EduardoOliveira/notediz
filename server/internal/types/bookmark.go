package types

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/EduardoOliveira/notediz/internal/lib/tools"
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
	} else if u, vErr := url.Parse(b.URL); vErr != nil {
		if u.Scheme != "http" && u.Scheme != "https" {
			err = errors.Join(err, fmt.Errorf("url must have a scheme (http or https)"))
		}
		if u.Host == "" {
			err = errors.Join(err, fmt.Errorf("url must have a host"))
		}
		err = errors.Join(err, fmt.Errorf("error validating url: %v", vErr))
	}
	return err
}

func (b *Bookmark) CreateFromAny(v map[string]any) error {
	b.FromAny(v)

	b.ID = tools.UUID()
	b.CreatedAt = tools.Now()
	b.UpdatedAt = tools.Now()
	b.Kind = NoteKindBookmark

	if err := b.Validate(); err != nil {
		return err
	}
	return nil
}

func (b *Bookmark) FromAny(v map[string]any) {
	if id, ok := v["ID"].(string); ok {
		b.ID = id
	}
	if url, ok := v["URL"].(string); ok {
		b.URL = url
	}
	if title, ok := v["Title"].(string); ok {
		b.Title = title
	}
	if description, ok := v["Description"].(string); ok {
		b.Description = description
	}
	anyToTime(&b.CreatedAt, v["CreatedAt"])
	anyToTime(&b.UpdatedAt, v["UpdatedAt"])
	if kind, ok := v["Kind"].(string); ok {
		b.Kind = NoteKind(kind)
	}
}
