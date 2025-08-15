-- name: CreateText :exec
INSERT INTO text (id, note_id, content, created_at, updated_at) VALUES
(sqlc.arg(id), sqlc.arg(note_id), sqlc.arg(content), sqlc.arg(created_at), sqlc.arg(updated_at));

-- name: GetTextByID :one
SELECT id, note_id, content, created_at, updated_at 
FROM text 
WHERE id = sqlc.arg(id) LIMIT 1;