// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: exercise.sql

package queries

import (
	"context"
)

const getExerciseDetails = `-- name: GetExerciseDetails :one
SELECT exercise, rating, cret_ts, updt_ts FROM tracker.exercise
WHERE Exercise = $1 LIMIT 1
`

func (q *Queries) GetExerciseDetails(ctx context.Context, exercise string) (TrackerExercise, error) {
	row := q.db.QueryRowContext(ctx, getExerciseDetails, exercise)
	var i TrackerExercise
	err := row.Scan(
		&i.Exercise,
		&i.Rating,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
