// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: composition.sql

package models

import (
	"context"
)

const deleteComposition = `-- name: DeleteComposition :exec
DELETE FROM tracker.composition
WHERE SUBMITTED_ON = $1
`

func (q *Queries) DeleteComposition(ctx context.Context, submittedOn string) error {
	_, err := q.db.ExecContext(ctx, deleteComposition, submittedOn)
	return err
}

const getComposition = `-- name: GetComposition :one
SELECT submitted_on, weight, bodyfat, neck, shoulders, left_bicep, right_bicep, left_tricep, right_tricep, left_forearm, right_forearm, chest, waist, left_quad, right_quad, left_calf, right_calf, cret_ts, updt_ts FROM tracker.composition
WHERE SUBMITTED_ON = $1 LIMIT 1
`

func (q *Queries) GetComposition(ctx context.Context, submittedOn string) (TrackerComposition, error) {
	row := q.db.QueryRowContext(ctx, getComposition, submittedOn)
	var i TrackerComposition
	err := row.Scan(
		&i.SubmittedOn,
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

const getCompositionDates = `-- name: GetCompositionDates :many
SELECT SUBMITTED_ON FROM tracker.composition
LIMIT $1
`

func (q *Queries) GetCompositionDates(ctx context.Context, limit int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getCompositionDates, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var submitted_on string
		if err := rows.Scan(&submitted_on); err != nil {
			return nil, err
		}
		items = append(items, submitted_on)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const submitComposition = `-- name: SubmitComposition :one
INSERT INTO tracker.composition (
  SUBMITTED_ON, WEIGHT, BODYFAT
) VALUES (
  $1, $2, $3
)
RETURNING submitted_on, weight, bodyfat, neck, shoulders, left_bicep, right_bicep, left_tricep, right_tricep, left_forearm, right_forearm, chest, waist, left_quad, right_quad, left_calf, right_calf, cret_ts, updt_ts
`

type SubmitCompositionParams struct {
	SubmittedOn string `json:"submittedOn"`
	Weight      string `json:"weight"`
	Bodyfat     string `json:"bodyfat"`
}

func (q *Queries) SubmitComposition(ctx context.Context, arg SubmitCompositionParams) (TrackerComposition, error) {
	row := q.db.QueryRowContext(ctx, submitComposition, arg.SubmittedOn, arg.Weight, arg.Bodyfat)
	var i TrackerComposition
	err := row.Scan(
		&i.SubmittedOn,
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
