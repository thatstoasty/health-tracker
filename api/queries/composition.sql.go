// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: composition.sql

package api

import (
	"context"
	"database/sql"
	"time"
)

const createComposition = `-- name: CreateComposition :one
INSERT INTO tracker.composition (
  WEIGHT, BODYFAT
) VALUES (
  $1, $2
)
RETURNING date, weight, bodyfat, neck, shoulders, left_bicep, right_bicep, left_tricep, right_tricep, left_forearm, right_forearm, chest, waist, left_quad, right_quad, left_calf, right_calf, cret_ts, updt_ts
`

type CreateCompositionParams struct {
	Weight  string
	Bodyfat sql.NullInt16
}

func (q *Queries) CreateComposition(ctx context.Context, arg CreateCompositionParams) (TrackerComposition, error) {
	row := q.db.QueryRowContext(ctx, createComposition, arg.Weight, arg.Bodyfat)
	var i TrackerComposition
	err := row.Scan(
		&i.Date,
		&i.Weight,
		&i.Bodyfat,
		&i.Neck,
		&i.Shoulders,
		&i.LeftBicep,
		&i.RightBicep,
		&i.LeftTricep,
		&i.RightTricep,
		&i.LeftForearm,
		&i.RightForearm,
		&i.Chest,
		&i.Waist,
		&i.LeftQuad,
		&i.RightQuad,
		&i.LeftCalf,
		&i.RightCalf,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const deleteComposition = `-- name: DeleteComposition :exec
DELETE FROM tracker.composition
WHERE DATE = $1
`

func (q *Queries) DeleteComposition(ctx context.Context, date time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteComposition, date)
	return err
}

const getComposition = `-- name: GetComposition :one
SELECT date, weight, bodyfat, neck, shoulders, left_bicep, right_bicep, left_tricep, right_tricep, left_forearm, right_forearm, chest, waist, left_quad, right_quad, left_calf, right_calf, cret_ts, updt_ts FROM tracker.composition
WHERE DATE = $1 LIMIT 1
`

func (q *Queries) GetComposition(ctx context.Context, date time.Time) (TrackerComposition, error) {
	row := q.db.QueryRowContext(ctx, getComposition, date)
	var i TrackerComposition
	err := row.Scan(
		&i.Date,
		&i.Weight,
		&i.Bodyfat,
		&i.Neck,
		&i.Shoulders,
		&i.LeftBicep,
		&i.RightBicep,
		&i.LeftTricep,
		&i.RightTricep,
		&i.LeftForearm,
		&i.RightForearm,
		&i.Chest,
		&i.Waist,
		&i.LeftQuad,
		&i.RightQuad,
		&i.LeftCalf,
		&i.RightCalf,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
