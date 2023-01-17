// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: nutrition.sql

package queries

import (
	"context"
	"database/sql"
)

const deleteNutrition = `-- name: DeleteNutrition :exec
DELETE FROM tracker.nutrition
WHERE SUBMITTED_ON = $1
`

func (q *Queries) DeleteNutrition(ctx context.Context, submittedOn string) error {
	_, err := q.db.ExecContext(ctx, deleteNutrition, submittedOn)
	return err
}

const getNutritionDetails = `-- name: GetNutritionDetails :one
SELECT submitted_on, calories, protein, carbohydrate, fat, micronutrients, cret_ts, updt_ts FROM tracker.nutrition
WHERE SUBMITTED_ON = $1 LIMIT 1
`

func (q *Queries) GetNutritionDetails(ctx context.Context, submittedOn string) (TrackerNutrition, error) {
	row := q.db.QueryRowContext(ctx, getNutritionDetails, submittedOn)
	var i TrackerNutrition
	err := row.Scan(
		&i.SubmittedOn,
		&i.Calories,
		&i.Protein,
		&i.Carbohydrate,
		&i.Fat,
		&i.Micronutrients,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const submitNutrition = `-- name: SubmitNutrition :one
INSERT INTO tracker.nutrition (
  CALORIES, PROTEIN, CARBOHYDRATE, FAT
) VALUES (
  $1, $2, $3, $4
)
RETURNING submitted_on, calories, protein, carbohydrate, fat, micronutrients, cret_ts, updt_ts
`

type SubmitNutritionParams struct {
	Calories     int16
	Protein      sql.NullInt16
	Carbohydrate sql.NullInt16
	Fat          sql.NullInt16
}

func (q *Queries) SubmitNutrition(ctx context.Context, arg SubmitNutritionParams) (TrackerNutrition, error) {
	row := q.db.QueryRowContext(ctx, submitNutrition,
		arg.Calories,
		arg.Protein,
		arg.Carbohydrate,
		arg.Fat,
	)
	var i TrackerNutrition
	err := row.Scan(
		&i.SubmittedOn,
		&i.Calories,
		&i.Protein,
		&i.Carbohydrate,
		&i.Fat,
		&i.Micronutrients,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
