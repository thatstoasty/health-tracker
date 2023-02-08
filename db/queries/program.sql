-- name: GetProgramNames :many
SELECT name FROM tracker.program
LIMIT $1;

-- name: GetProgram :many
SELECT a.name, b.name, c.group_id, c.exercise_name, c.weight, c.sets, c.reps FROM tracker.program a
JOIN tracker.workout b
ON a.name = b.program_name
JOIN tracker.workout_details c
ON b.name = c.workout_name
WHERE a.name = $1
LIMIT 1;

-- name: DeleteProgram :exec
DELETE FROM tracker.program
WHERE NAME = $1;

-- name: SubmitProgram :one
INSERT INTO tracker.program (
  NAME
) VALUES (
  $1
)
RETURNING *;

-- name: SubmitProgramDetails :one
INSERT INTO tracker.program_details (
  PROGRAM_NAME, WORKOUT_NAME
) VALUES (
  $1, $2
)
RETURNING *;