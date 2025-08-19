package dbsql

import (
	"database/sql"
	"path/filepath"
	"time"

	"github.com/EduardoOliveira/notediz/internal/dbsql/gen"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Repo struct {
	Db      *sql.DB
	queries *gen.Queries
	now     func() time.Time
	uuid    func() string
}

func MustNew(basePath string) *Repo {
	db, err := sql.Open("sqlite", filepath.Join(basePath, "notediz.db"))
	if err != nil {
		panic(err)
	}
	return &Repo{
		Db:      db,
		queries: gen.New(db),
		now:     time.Now,
		uuid:    uuid.NewString,
	}
}
