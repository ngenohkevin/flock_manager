- name: CreateProduction :one

INSERT INTO roduction (
                   breed_name
) VALUES (
           $1
) RETURNING *;