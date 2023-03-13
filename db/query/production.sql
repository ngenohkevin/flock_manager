-- name: CreateProduction :one

INSERT INTO production (
                 production_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs
) VALUES (
           $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProduction :one
SELECT * FROM production
WHERE production_id = $1 LIMIT 1;

-- name: ListProduction :many
SELECT * FROM production
WHERE production_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProduction :one
UPDATE production
SET eggs = $2,
    dirty = $3,
    wrong_shape = $4,
    weak_shell = $5,
    damaged = $6,
    hatching_eggs = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProduction :exec
DELETE FROM production
WHERE id = $1;



