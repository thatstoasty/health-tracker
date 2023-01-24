-- name: GetExerciseDetails :many
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

-- name: DeleteExerciseDetails :exec
DELETE FROM tracker.exercise_details
WHERE EXERCISE_NAME = $1;