-- name: GetExercise :many
SELECT a.name, b.body_part, b.level FROM tracker.exercise a
JOIN tracker.exercise_details b
ON a.name = b.exercise_name
WHERE NAME = $1 LIMIT 1;

-- name: GetExercises :many
SELECT name FROM tracker.exercise
LIMIT $1;

-- name: DeleteExercise :exec
DELETE FROM tracker.exercise
WHERE NAME = $1;

-- name: SubmitExercise :one
INSERT INTO tracker.exercise (
  NAME, TYPE, VARIATION
) VALUES (
  $1, $2, $3
)
ON CONFLICT (NAME) 
DO UPDATE SET 
  TYPE = $2,
  VARIATION = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;

-- name: SubmitExerciseDetails :one
INSERT INTO tracker.exercise_details (
  EXERCISE_NAME, BODY_PART, LEVEL
) VALUES (
  $1, $2, $3
)
ON CONFLICT (EXERCISE_NAME, BODY_PART) 
DO UPDATE SET 
  LEVEL = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;