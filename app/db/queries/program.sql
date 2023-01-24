-- name: GetProgramNames :many
SELECT name FROM tracker.program
LIMIT $1;

-- name: GetProgramDetails :many
SELECT a.name, b.name, c.group_id, c.exercise_name, c.weight, c.sets, c.reps FROM tracker.program a
JOIN tracker.workout b
ON a.name = b.program_name
JOIN tracker.workout_details c
ON b.name = c.workout_name
LIMIT $1;

-- name: DeleteProgram :exec
DELETE FROM tracker.program CASCADE
WHERE NAME = $1;

-- name: DeleteProgramDetails :exec
DELETE FROM tracker.program_details
WHERE PROGRAM_NAME = $1;