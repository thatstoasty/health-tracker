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

-- name: SubmitWorkout :one
INSERT INTO tracker.workout (
  NAME, PROGRAM_NAME
) VALUES (
  $1, $2
)
RETURNING *;

-- name: SubmitWorkoutDetails :one
INSERT INTO tracker.workout_details (
  WORKOUT_NAME, GROUP_ID, EXERCISE_NAME, SETS, REPS, WEIGHT
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;