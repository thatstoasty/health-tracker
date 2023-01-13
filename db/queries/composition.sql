-- name: CreateComposition :one
INSERT INTO tracker.composition (
  WEIGHT, BODYFAT
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetComposition :one
SELECT * FROM tracker.composition
WHERE DATE = $1 LIMIT 1;

-- name: DeleteComposition :exec
DELETE FROM tracker.composition
WHERE DATE = $1;