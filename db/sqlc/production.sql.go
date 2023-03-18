// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: production.sql

package db

import (
	"context"
)

const createProduction = `-- name: CreateProduction :one

INSERT INTO production (
                 breed_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs
) VALUES (
           $1, $2, $3, $4, $5, $6, $7
) RETURNING id, breed_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs, created_at
`

type CreateProductionParams struct {
	BreedID      int64 `json:"breed_id"`
	Eggs         int64 `json:"eggs"`
	Dirty        int64 `json:"dirty"`
	WrongShape   int64 `json:"wrong_shape"`
	WeakShell    int64 `json:"weak_shell"`
	Damaged      int64 `json:"damaged"`
	HatchingEggs int64 `json:"hatching_eggs"`
}

func (q *Queries) CreateProduction(ctx context.Context, arg CreateProductionParams) (Production, error) {
	row := q.db.QueryRowContext(ctx, createProduction,
		arg.BreedID,
		arg.Eggs,
		arg.Dirty,
		arg.WrongShape,
		arg.WeakShell,
		arg.Damaged,
		arg.HatchingEggs,
	)
	var i Production
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Eggs,
		&i.Dirty,
		&i.WrongShape,
		&i.WeakShell,
		&i.Damaged,
		&i.HatchingEggs,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProduction = `-- name: DeleteProduction :exec
DELETE FROM production
WHERE id = $1
`

func (q *Queries) DeleteProduction(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduction, id)
	return err
}

const getProduction = `-- name: GetProduction :one
SELECT id, breed_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs, created_at FROM production
WHERE breed_id = $1 LIMIT 1
`

func (q *Queries) GetProduction(ctx context.Context, breedID int64) (Production, error) {
	row := q.db.QueryRowContext(ctx, getProduction, breedID)
	var i Production
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Eggs,
		&i.Dirty,
		&i.WrongShape,
		&i.WeakShell,
		&i.Damaged,
		&i.HatchingEggs,
		&i.CreatedAt,
	)
	return i, err
}

const listProduction = `-- name: ListProduction :many
SELECT id, breed_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs, created_at FROM production
WHERE breed_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListProductionParams struct {
	BreedID int64 `json:"breed_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListProduction(ctx context.Context, arg ListProductionParams) ([]Production, error) {
	rows, err := q.db.QueryContext(ctx, listProduction, arg.BreedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Production
	for rows.Next() {
		var i Production
		if err := rows.Scan(
			&i.ID,
			&i.BreedID,
			&i.Eggs,
			&i.Dirty,
			&i.WrongShape,
			&i.WeakShell,
			&i.Damaged,
			&i.HatchingEggs,
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

const updateProduction = `-- name: UpdateProduction :one
UPDATE production
SET eggs = $2,
    dirty = $3,
    wrong_shape = $4,
    weak_shell = $5,
    damaged = $6,
    hatching_eggs = $7
WHERE id = $1
RETURNING id, breed_id, eggs, dirty, wrong_shape, weak_shell, damaged, hatching_eggs, created_at
`

type UpdateProductionParams struct {
	ID           int64 `json:"id"`
	Eggs         int64 `json:"eggs"`
	Dirty        int64 `json:"dirty"`
	WrongShape   int64 `json:"wrong_shape"`
	WeakShell    int64 `json:"weak_shell"`
	Damaged      int64 `json:"damaged"`
	HatchingEggs int64 `json:"hatching_eggs"`
}

func (q *Queries) UpdateProduction(ctx context.Context, arg UpdateProductionParams) (Production, error) {
	row := q.db.QueryRowContext(ctx, updateProduction,
		arg.ID,
		arg.Eggs,
		arg.Dirty,
		arg.WrongShape,
		arg.WeakShell,
		arg.Damaged,
		arg.HatchingEggs,
	)
	var i Production
	err := row.Scan(
		&i.ID,
		&i.BreedID,
		&i.Eggs,
		&i.Dirty,
		&i.WrongShape,
		&i.WeakShell,
		&i.Damaged,
		&i.HatchingEggs,
		&i.CreatedAt,
	)
	return i, err
}
