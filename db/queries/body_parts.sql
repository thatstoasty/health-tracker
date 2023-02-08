-- name: SubmitBodyPart :one
INSERT INTO tracker.body_parts (
  NAME, REGION, UPPER_OR_LOWER
) VALUES (
  $1, $2, $3
)
RETURNING *;