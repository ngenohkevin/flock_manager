-- name: CreateBreed :one

INSERT INTO breed (
                   breed_name
) VALUES (
           $1
) RETURNING *;
