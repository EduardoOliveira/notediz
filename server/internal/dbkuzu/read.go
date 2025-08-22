package dbkuzu

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/EduardoOliveira/notediz/internal/types"
	"github.com/kuzudb/go-kuzu"
)

type Row struct {
	Properties map[string]any
}

func (db *Repo) GetNoteContext(noteID string) ([]any, error) {
	query := `MATCH (b:Bookmark)
	WITH b
	MATCH (t:Text)
	RETURN b, t
	`

	prep, err := db.Conn.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer prep.Close()

	result, err := db.Conn.Execute(prep, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer result.Close()

	rtn := make([]any, 0)

	for result.HasNext() {
		tuple, err := result.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next tuple: %v", err)
		}
		defer tuple.Close()

		rows, err := tuple.GetAsSlice()
		if err != nil {
			return nil, fmt.Errorf("failed to get rows as slice: %v", err)
		}
		slog.Info("GetNoteContext", "rows", rows)
		for _, row := range rows {
			n, ok := row.(kuzu.Node)
			if !ok {
				slog.Error("GetNoteContext", "error", "failed to cast node")
				continue
			}
			slog.Info("GetNoteContext", "node", n)

			var note any
			switch strings.ToLower(n.Label) {
			case string(types.NoteKindBookmark):
				bm := types.Bookmark{}
				bm.FromAny(n.Properties)
				note = bm
			case string(types.NoteKindText):
				txt := types.Text{}
				txt.FromAny(n.Properties)
				note = txt
			default:
				slog.Info("GetNoteContext", "kind", "no clue")
				note = row
			}
			rtn = append(rtn, note)
		}

	}

	return rtn, nil
}
