// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: breed.sql

package db

import (
	"context"
)

const createBreed = `-- name: CreateBreed :one

INSERT INTO breed (
                   breed_name
) VALUES (
           $1
) RETURNING breed_id, breed_name, created_at
`

func (q *Queries) CreateBreed(ctx context.Context, breedName string) (Breed, error) {
	row := q.db.QueryRowContext(ctx, createBreed, breedName)
	var i Breed
	err := row.Scan(&i.BreedID, &i.BreedName, &i.CreatedAt)
	return i, err
}

const deleteBreed = `-- name: DeleteBreed :exec
DELETE FROM breed
WHERE breed_id = $1
`

func (q *Queries) DeleteBreed(ctx context.Context, breedID int64) error {
	_, err := q.db.ExecContext(ctx, deleteBreed, breedID)
	return err
}

const getBreed = `-- name: GetBreed :one
SELECT breed_id, breed_name, created_at FROM breed
WHERE breed_id = $1 LIMIT 1
`

func (q *Queries) GetBreed(ctx context.Context, breedID int64) (Breed, error) {
	row := q.db.QueryRowContext(ctx, getBreed, breedID)
	var i Breed
	err := row.Scan(&i.BreedID, &i.BreedName, &i.CreatedAt)
	return i, err
}

const listBreeds = `-- name: ListBreeds :many
SELECT breed_id, breed_name, created_at FROM breed
ORDER BY breed_id
    LIMIT $1
OFFSET $2
`

type ListBreedsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBreeds(ctx context.Context, arg ListBreedsParams) ([]Breed, error) {
	rows, err := q.db.QueryContext(ctx, listBreeds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Breed
	for rows.Next() {
		var i Breed
		if err := rows.Scan(&i.BreedID, &i.BreedName, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBreed = `-- name: UpdateBreed :one
UPDATE breed
set breed_name = $2
WHERE breed_id = $1
    RETURNING breed_id, breed_name, created_at
`

type UpdateBreedParams struct {
	BreedID   int64  `json:"breed_id"`
	BreedName string `json:"breed_name"`
}

func (q *Queries) UpdateBreed(ctx context.Context, arg UpdateBreedParams) (Breed, error) {
	row := q.db.QueryRowContext(ctx, updateBreed, arg.BreedID, arg.BreedName)
	var i Breed
	err := row.Scan(&i.BreedID, &i.BreedName, &i.CreatedAt)
	return i, err
}
