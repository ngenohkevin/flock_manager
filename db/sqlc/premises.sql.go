// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: premises.sql

package db

import (
	"context"
)

const createPremises = `-- name: CreatePremises :one

INSERT INTO premises (breed_id, farm, house)
VALUES ($1, $2, $3)
RETURNING id, breed_id, farm, house, created_at
`

type CreatePremisesParams struct {
	BreedID int64  `json:"breed_id"`
	Farm    string `json:"farm"`
	House   string `json:"house"`
}

func (q *Queries) CreatePremises(ctx context.Context, arg CreatePremisesParams) (Premise, error) {
	row := q.db.QueryRowContext(ctx, createPremises, arg.BreedID, arg.Farm, arg.House)
	var i Premise
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Farm,
		&i.House,
		&i.CreatedAt,
	)
	return i, err
}

const deletePremises = `-- name: DeletePremises :exec
DELETE FROM premises
WHERE id = $1
`

func (q *Queries) DeletePremises(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePremises, id)
	return err
}

const getPremises = `-- name: GetPremises :one
SELECT id, breed_id, farm, house, created_at FROM premises
WHERE breed_id = $1 LIMIT 1
`

func (q *Queries) GetPremises(ctx context.Context, breedID int64) (Premise, error) {
	row := q.db.QueryRowContext(ctx, getPremises, breedID)
	var i Premise
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Farm,
		&i.House,
		&i.CreatedAt,
	)
	return i, err
}

const listPremises = `-- name: ListPremises :many
SELECT id, breed_id, farm, house, created_at FROM premises
WHERE breed_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListPremisesParams struct {
	BreedID int64 `json:"breed_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListPremises(ctx context.Context, arg ListPremisesParams) ([]Premise, error) {
	rows, err := q.db.QueryContext(ctx, listPremises, arg.BreedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Premise
	for rows.Next() {
		var i Premise
		if err := rows.Scan(
			&i.ID,
			&i.BreedID,
			&i.Farm,
			&i.House,
			&i.CreatedAt,
		); err != nil {
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

const updatePremises = `-- name: UpdatePremises :one
UPDATE premises
SET farm = $2,
    house = $3
WHERE id = $1
RETURNING id, breed_id, farm, house, created_at
`

type UpdatePremisesParams struct {
	ID    int64  `json:"id"`
	Farm  string `json:"farm"`
	House string `json:"house"`
}

func (q *Queries) UpdatePremises(ctx context.Context, arg UpdatePremisesParams) (Premise, error) {
	row := q.db.QueryRowContext(ctx, updatePremises, arg.ID, arg.Farm, arg.House)
	var i Premise
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Farm,
		&i.House,
		&i.CreatedAt,
	)
	return i, err
}
