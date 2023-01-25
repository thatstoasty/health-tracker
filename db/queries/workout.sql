-- name: GetWorkoutPerformed :one
SELECT * FROM tracker.workout_performed
WHERE SUBMITTED_ON = $1 LIMIT 1;

-- name: DeleteWorkoutPerformed :exec
DELETE FROM tracker.workout_performed
WHERE SUBMITTED_ON = $1;

-- name: GetWorkoutNames :many
SELECT NAME FROM tracker.workout
LIMIT $1;

-- name: GetWorkout :one
SELECT * FROM tracker.workout
WHERE NAME = $1 LIMIT 1;

-- name: DeleteWorkout :exec
DELETE FROM tracker.workout
WHERE NAME = $1;