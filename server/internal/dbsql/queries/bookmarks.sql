-- name: CreateBookmark :exec
INSERT INTO bookmark (id, note_id, url, title, description, created_at, updated_at) VALUES
(sqlc.arg(id), sqlc.arg(note_id), sqlc.arg(url), sqlc.arg(title), sqlc.arg(description), sqlc.arg(created_at), sqlc.arg(updated_at));

-- name: GetBookmarkByID :one
SELECT id, note_id, url, title, description, created_at, updated_at
FROM bookmark
WHERE id = sqlc.arg(id);
