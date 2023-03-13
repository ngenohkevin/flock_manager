-- name: CreateProduction :one

INSERT INTO production (
                 production_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs
) VALUES (
           $1, $2, $3, $4, $5, $6, $7
) RETURNING *;
