package types

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/EduardoOliveira/notediz/internal/lib/opt"
)

type NoteKind string

var (
	NoteKindText     NoteKind = "text"
	NoteKindAudio    NoteKind = "audio"
	NoteKindBookmark NoteKind = "bookmark"
	NoteKindFile     NoteKind = "file"
	NoteKindImage    NoteKind = "image"
	ToteKindVideo    NoteKind = "video"
	NoteKindTodo     NoteKind = "todo"
)

type Bookmark struct {
	ID          string
	NoteID      string
	Url         string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b Bookmark) Validate() error {
	var err error
	if b.Url == "" {
		slog.Error("Empty URL", "bookmark", b)
		err = errors.Join(err, fmt.Errorf("url cannot be empty"))
	} else if _, vErr := url.Parse(b.Url); vErr != nil {
		err = errors.Join(err, fmt.Errorf("error validating url: %v", vErr))
	}
	return err
}

func BookmarkFromAny(v map[string]any) Bookmark {
	var b Bookmark
	if id, ok := v["id"].(string); ok {
		b.ID = id
	}
	if noteID, ok := v["note_id"].(string); ok {
		b.NoteID = noteID
	}
	if url, ok := v["url"].(string); ok {
		b.Url = url
	}
	if title, ok := v["title"].(string); ok {
		b.Title = title
	}
	if description, ok := v["description"].(string); ok {
		b.Description = description
	}

	return b
}

type Audio struct {
	ID        string
	NoteID    string
	Url       string
	Title     string
	Duration  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Connection struct {
	ID        string
	From      string
	To        string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type File struct {
	ID        string
	NoteID    string
	Name      string
	Path      string
	Size      int64
	MimeType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Image struct {
	ID        string
	NoteID    string
	Url       string
	Alt       opt.Optional[string]
	Width     int64
	Height    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateNote struct {
	Content string
	Kind    NoteKind
}

func (c CreateNote) Validate() error {
	var err error
	if c.Content == "" {
		err = errors.Join(err, fmt.Errorf("content cannot be empty"))
	}
	return err
}

type Note struct {
	ID        string
	Kind      NoteKind
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NoteTag struct {
	NoteID string
	TagID  string
}

type Tag struct {
	ID        string
	Name      string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Text struct {
	ID        string
	NoteID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID        string
	NoteID    string
	Content   string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Video struct {
	ID        string
	NoteID    string
	Url       string
	Title     string
	Duration  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
