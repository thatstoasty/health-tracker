-- name: SubmitWorkoutPerformed :one
INSERT INTO tracker.workout_performed (
  SUBMITTED_ON, WORKOUT_NAME
) VALUES (
  $1, $2
)
RETURNING *;

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
ON CONFLICT (NAME) 
DO UPDATE SET 
  PROGRAM_NAME = $2,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;

-- name: SubmitWorkoutDetails :one
INSERT INTO tracker.workout_details (
  WORKOUT_NAME, GROUP_ID, EXERCISE_NAME, SETS, REPS, WEIGHT
) VALUES (
  $1, $2, $3, $4, $5, $6
)
ON CONFLICT (WORKOUT_NAME, GROUP_ID, EXERCISE_NAME) 
DO UPDATE SET 
  SETS = $4,
  REPS = $5,
  WEIGHT = $6,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;