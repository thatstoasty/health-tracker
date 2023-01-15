-- name: SubmitWorkout :one
INSERT INTO tracker.workout (
  EXERCISE, SETS, REPS, WEIGHT, REPS_IN_RESERVE
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetWorkoutDetails :one
SELECT * FROM tracker.workout
WHERE DATE = $1 LIMIT 1;

-- name: DeleteWorkout :exec
DELETE FROM tracker.workout
WHERE DATE = $1;