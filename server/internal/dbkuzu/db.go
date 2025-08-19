package dbkuzu

import (
	"context"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/kuzudb/go-kuzu"
)

type Repo struct {
	Db   *kuzu.Database
	Conn *kuzu.Connection
	now  func() time.Time
	uuid func() string
}

func MustNew(ctx context.Context, basePath string) *Repo {
	systemConfig := kuzu.DefaultSystemConfig()
	systemConfig.BufferPoolSize = 1024 * 1024 * 1024
	db, err := kuzu.OpenDatabase(filepath.Join(basePath, "notediz.kuzu"), systemConfig)
	if err != nil {
		panic(err)
	}

	// Open a connection to the database.
	conn, err := kuzu.OpenConnection(db)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		db.Close()
		conn.Close()
	}()

	repo := &Repo{
		Db:   db,
		Conn: conn,
		now:  time.Now,
		uuid: uuid.NewString,
	}
	repo.initSchema()

	return repo
}

func (r *Repo) initSchema() {
	queries := []string{
		//TODO: explore kuzu uuid type
		`CREATE NODE TABLE Bookmark(id STRING, title STRING, url STRING, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id))`,
	}

	for _, query := range queries {
		if _, err := r.Conn.Query(query); err != nil {
			slog.Error("Failed to execute query", "query", query, "error", err)
			panic(err)
		}
	}
}
