-- name: CreateBreed :one

INSERT INTO breed (
                   breed_name,
                   username
) VALUES (
           $1, $2
) RETURNING *;

-- name: GetBreed :one
SELECT * FROM breed
WHERE breed_id = $1 LIMIT 1;

-- name: ListBreeds :many
SELECT * FROM breed
WHERE username = $1
ORDER BY breed_id
    LIMIT $2
OFFSET $3;

-- name: UpdateBreed :one
UPDATE breed
set breed_name = $2
WHERE breed_id = $1
    RETURNING *;

-- name: DeleteBreed :exec
DELETE FROM breed
WHERE breed_id = $1;
