-- name: CreateHatchery :one
INSERT INTO hatchery(
                     hatchery_id, infertile, early, middle, late, dead_chicks, alive_chicks
) VALUES (
          $1, $2, $3, $4, $5, $6, $7
         )
RETURNING *;

-- name: GetHatchery :one
SELECT * FROM hatchery
WHERE hatchery_id = $1 LIMIT 1;

-- name: ListHatchery :many
SELECT * FROM hatchery
WHERE hatchery_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateHatchery :one
UPDATE hatchery
SET infertile = $2,
    early = $3,
    middle = $4,
    late = $5,
    dead_chicks = $6,
    alive_chicks = $7
WHERE id = $1
RETURNING *;

-- name: DeleteHatchery :exec
DELETE FROM hatchery
WHERE id = $1;