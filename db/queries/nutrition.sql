-- name: SubmitNutrition :one
INSERT INTO tracker.nutrition (
  CALORIES, PROTEIN, CARBOHYDRATE, FAT
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetNutritionDetails :one
SELECT * FROM tracker.nutrition
WHERE DATE = $1 LIMIT 1;

-- name: DeleteNutrition :exec
DELETE FROM tracker.nutrition
WHERE DATE = $1;