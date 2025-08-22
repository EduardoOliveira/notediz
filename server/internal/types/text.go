package types

import (
	"errors"
	"fmt"
	"time"

	"github.com/EduardoOliveira/notediz/internal/lib/tools"
)

type Text struct {
	ID        string
	Kind      NoteKind
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Text) Validate() error {
	var err error
	if t.Kind != NoteKindText {
		err = errors.Join(err, fmt.Errorf("invalid kind: %s", t.Kind))
	}
	if t.Content == "" {
		err = errors.Join(err, fmt.Errorf("content cannot be empty"))
	}
	return err
}

func (t *Text) CreateFromAny(v map[string]any) error {
	t.FromAny(v)

	t.ID = tools.UUID()
	t.CreatedAt = tools.Now()
	t.UpdatedAt = tools.Now()
	t.Kind = NoteKindText

	if err := t.Validate(); err != nil {
		return err
	}
	return nil
}

func (t *Text) FromAny(v map[string]any) {
	anyToSome(&t.ID, v["ID"])
	anyToSome(&t.Content, v["Content"])

	anyToTime(&t.CreatedAt, v["CreatedAt"])
	anyToTime(&t.UpdatedAt, v["UpdatedAt"])

	if kind, ok := v["Kind"].(string); ok {
		t.Kind = NoteKind(kind)
	}
}
