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
