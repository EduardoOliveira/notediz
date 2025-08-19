-- name: CreateNote :exec
INSERT INTO note (id, kind, created_at, updated_at) VALUES
(sqlc.arg(id), sqlc.arg(kind), sqlc.arg(created_at), sqlc.arg(updated_at))
