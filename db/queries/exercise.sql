-- name: GetExerciseDetails :one
SELECT * FROM tracker.exercise
WHERE EXERCISE = $1 LIMIT 1;