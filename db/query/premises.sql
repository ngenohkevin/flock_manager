-- name: CreatePremises :one

INSERT INTO premises (breed_id, farm, house)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPremises :one
SELECT * FROM premises
WHERE breed_id = $1 LIMIT 1;

-- name: ListPremises :many
SELECT * FROM premises
WHERE breed_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePremises :one
UPDATE premises
SET farm = $2,
    house = $3
WHERE id = $1
RETURNING *;

-- name: DeletePremises :exec
DELETE FROM premises
WHERE id = $1;
