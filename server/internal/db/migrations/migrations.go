package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var migrations embed.FS

func MustMigrate(ctx context.Context, db *sql.DB) {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(fmt.Sprintf("setting sql dialect: %v", err))
	}

	if err := goose.UpContext(ctx, db, "."); err != nil {
		panic(fmt.Sprintf("running migrations: %v", err))
	}
}
