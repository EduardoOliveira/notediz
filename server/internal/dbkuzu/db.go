package dbkuzu

import (
	"context"
	"fmt"
	"log/slog"
	"os"
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
	repo.initSchema(basePath)

	return repo
}

func (r *Repo) initSchema(basePath string) {
	revision, err := r.getRevision(basePath)
	if err != nil {
		panic(err)
	}

	queries := []string{
		//TODO: explore kuzu uuid type
		`CREATE NODE TABLE Text(id STRING, kind STRING, content STRING, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id))`,
		`CREATE NODE TABLE Bookmark(id STRING, kind STRING, title STRING, url STRING, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id))`,
	}

	for i := revision; i < len(queries); i++ {
		if _, err := r.Conn.Query(queries[i]); err != nil {
			slog.Error("Failed to execute query", "query", queries[i], "error", err)
			panic(err)
		}
	}

	if err := r.setRevision(basePath, len(queries)); err != nil {
		panic(err)
	}
}

func (r *Repo) getRevision(basePath string) (int, error) {
	file, err := os.ReadFile(filepath.Join(basePath, "revision.txt"))
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	var revision int
	if _, err := fmt.Sscanf(string(file), "%d", &revision); err != nil {
		return 0, err
	}
	return revision, nil
}

func (r *Repo) setRevision(basePath string, revision int) error {
	file, err := os.Create(filepath.Join(basePath, "revision.txt"))
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := fmt.Fprintf(file, "%d", revision); err != nil {
		return err
	}
	return nil
}
