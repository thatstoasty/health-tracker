-- name: SubmitComposition :one
INSERT INTO tracker.composition (
  WEIGHT, BODYFAT
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetComposition :one
SELECT * FROM tracker.composition
WHERE SUBMITTED_ON = $1 LIMIT 1;

-- name: DeleteComposition :exec
DELETE FROM tracker.composition
WHERE SUBMITTED_ON = $1;

-- name: GetCompositionDates :many
SELECT SUBMITTED_ON FROM tracker.composition
LIMIT $1;
