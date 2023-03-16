package db

import (
	"context"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomProduction(t *testing.T, breed Breed) Production {
	arg := CreateProductionParams{
		ProductionID: breed.BreedID,
		Eggs:         util.RandomProduction(),
		Dirty:        util.RandomProduction(),
		WrongShape:   util.RandomProduction(),
		WeakShell:    util.RandomProduction(),
		Damaged:      util.RandomProduction(),
		HatchingEggs: util.RandomProduction(),
	}
	production, err := testQueries.CreateProduction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, production)

	require.Equal(t, arg.ProductionID, production.ProductionID)
	require.Equal(t, arg.Eggs, production.Eggs)
	require.Equal(t, arg.Dirty, production.Dirty)
	require.Equal(t, arg.WrongShape, production.WrongShape)
	require.Equal(t, arg.WeakShell, production.WeakShell)
	require.Equal(t, arg.Damaged, production.Damaged)
	require.Equal(t, arg.HatchingEggs, production.HatchingEggs)

	require.NotZero(t, production.ID)
	require.NotZero(t, production.CreatedAt)

	return production
}

func TestCreateProduction(t *testing.T) {
	breed := createdRandomBreed(t)
	createRandomProduction(t, breed)
}
