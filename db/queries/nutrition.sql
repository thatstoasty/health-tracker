-- name: SubmitNutrition :one
INSERT INTO tracker.nutrition (
  SUBMITTED_ON, CALORIES, PROTEIN, CARBOHYDRATE, FAT
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetNutrition :one
SELECT * FROM tracker.nutrition
WHERE SUBMITTED_ON = $1 LIMIT 1;

-- name: DeleteNutrition :exec
DELETE FROM tracker.nutrition
WHERE SUBMITTED_ON = $1;

-- name: GetNutritionDates :many
SELECT SUBMITTED_ON FROM tracker.nutrition
LIMIT $1;