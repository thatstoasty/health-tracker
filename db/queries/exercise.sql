-- name: GetExerciseDetails :one
SELECT * FROM tracker.exercise
WHERE Exercise = $1 LIMIT 1;