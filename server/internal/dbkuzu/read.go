package dbkuzu

import (
	"log/slog"

	"github.com/EduardoOliveira/notediz/internal/types"
	"github.com/kuzudb/go-kuzu"
)

type Row struct {
	Properties map[string]any
}

func (db *Repo) GetNoteContext(noteID string) ([]any, error) {
	query := `MATCH (n:Bookmark)
	RETURN n
	`

	prep, err := db.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	result, err := db.Conn.Execute(prep, map[string]any{})
	if err != nil {
		return nil, err
	}
	defer result.Close()

	rtn := make([]any, 0)

	for result.HasNext() {
		tuple, err := result.Next()
		if err != nil {
			return nil, err
		}
		defer tuple.Close()

		rows, err := tuple.GetAsSlice()
		if err != nil {
			return nil, err
		}
		for _, row := range rows {
			n, ok := row.(kuzu.Node)
			if !ok {
				slog.Error("GetNoteContext", "error", "failed to cast node")
				continue
			}

			var note any
			switch n.Properties["kind"] {
			case string(types.NoteKindBookmark):
				note, err = types.BookmarkFromFlatTuple(n.Properties)
				if err != nil {
					return nil, err
				}
			default:
				slog.Info("GetNoteContext", "kind", "no clue")
				note = row
			}
			rtn = append(rtn, note)
		}

	}

	return rtn, nil
}
