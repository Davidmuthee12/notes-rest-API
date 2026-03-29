-- name: ListNotes :many
SELECT * FROM notes;

-- name: CreateNote :one
INSERT INTO notes (content) VALUES ($1) RETURNING *;

-- name: GetNoteByID :one
SELECT * FROM notes WHERE id = $1;

-- name: DeleteNoteByID :exec
DELETE FROM notes WHERE id = $1;

-- name: UpdateNoteByID :one
UPDATE notes SET content = $2 WHERE id = $1 RETURNING *;
