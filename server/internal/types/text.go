package types

import (
	"errors"
	"fmt"
	"time"
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

func (t *Text) FromAny(v map[string]any) {
	anyToSome(&t.ID, v["id"])
	anyToSome(&t.Content, v["content"])

	anyToTime(&t.CreatedAt, v["created_at"])
	anyToTime(&t.UpdatedAt, v["updated_at"])

	if kind, ok := v["kind"].(string); ok {
		t.Kind = NoteKind(kind)
	}
}
