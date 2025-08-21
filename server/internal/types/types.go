package types

import (
	"errors"
	"fmt"
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

func anyToTime(t *time.Time, v any) {
	switch val := v.(type) {
	case string:
		if parsed, err := time.Parse(time.RFC3339, val); err == nil {
			*t = parsed
		}
	case time.Time:
		*t = val
	default:
		*t = time.Time{}
	}
}
func anyToString(s *string, v any, fallback ...string) {
	if val, ok := v.(string); ok {
		*s = val
	} else {
		if len(fallback) > 0 {
			*s = fallback[0]
		} else {
			*s = ""
		}
	}
}

func anyToSome[T any](s *T, v any, fallback ...T) {
	if val, ok := v.(T); ok {
		*s = val
	} else {
		if len(fallback) > 0 {
			*s = fallback[0]
		} else {
			var zero T
			*s = zero
		}
	}
}

type Audio struct {
	ID        string
	NoteID    string
	URL       string
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
	URL       string
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
	NoteID    string
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
	URL       string
	Title     string
	Duration  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
