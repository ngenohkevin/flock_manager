// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateBreed(ctx context.Context, breedName string) (Breed, error)
	CreateHatchery(ctx context.Context, arg CreateHatcheryParams) (Hatchery, error)
	CreatePremises(ctx context.Context, arg CreatePremisesParams) (Premise, error)
	CreateProduction(ctx context.Context, arg CreateProductionParams) (Production, error)
	DeleteBreed(ctx context.Context, breedID int64) error
	DeleteHatchery(ctx context.Context, id int64) error
	DeletePremises(ctx context.Context, id int64) error
	DeleteProduction(ctx context.Context, id int64) error
	GetBreed(ctx context.Context, breedID int64) (Breed, error)
	GetHatchery(ctx context.Context, productionID int64) (Hatchery, error)
	GetPremises(ctx context.Context, breedID int64) (Premise, error)
	GetProduction(ctx context.Context, id int64) (Production, error)
	ListBreeds(ctx context.Context, arg ListBreedsParams) ([]Breed, error)
	ListHatchery(ctx context.Context, arg ListHatcheryParams) ([]Hatchery, error)
	ListPremises(ctx context.Context, arg ListPremisesParams) ([]Premise, error)
	ListProduction(ctx context.Context, arg ListProductionParams) ([]Production, error)
	UpdateBreed(ctx context.Context, arg UpdateBreedParams) (Breed, error)
	UpdateHatchery(ctx context.Context, arg UpdateHatcheryParams) (Hatchery, error)
	UpdatePremises(ctx context.Context, arg UpdatePremisesParams) (Premise, error)
	UpdateProduction(ctx context.Context, arg UpdateProductionParams) (Production, error)
}

var _ Querier = (*Queries)(nil)
