-- name: CreateBreed :one

INSERT INTO breed (
                   breed_name
) VALUES (
           $1
) RETURNING *;

-- name: GetBreed :one
SELECT * FROM breed
WHERE breed_id = $1 LIMIT 1;

-- name: ListBreeds :many
SELECT * FROM breed
ORDER BY breed_id
    LIMIT $1
OFFSET $2;

-- name: UpdateBreed  :one
UPDATE breed
set breed_name = $2
WHERE breed_id = $1
    RETURNING *;

-- name: DeleteBreed :exec
DELETE FROM breed
WHERE breed_id = $1;
