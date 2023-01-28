-- name: CreateItem :one
INSERT INTO items (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: ListItems :many
SELECT *
  FROM items
 WHERE (id < sqlc.narg('id') OR sqlc.narg('id') IS NULL)
 ORDER BY id DESC
 FETCH FIRST $1 ROWS ONLY;