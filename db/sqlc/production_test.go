package db

import (
	"context"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

func TestGetProduction(t *testing.T) {
	breed := createdRandomBreed(t)

	production1 := createRandomProduction(t, breed)

	production2, err := testQueries.GetProduction(context.Background(), production1.ProductionID)
	require.NoError(t, err)
	require.NotEmpty(t, production2)

	require.Equal(t, production1.ProductionID, production2.ProductionID)
	require.Equal(t, production1.ID, production2.ID)
	require.Equal(t, production1.Eggs, production2.Eggs)
	require.Equal(t, production1.Dirty, production2.Dirty)
	require.Equal(t, production1.WrongShape, production2.WrongShape)
	require.Equal(t, production1.WeakShell, production2.WeakShell)
	require.Equal(t, production1.Damaged, production2.Damaged)
	require.Equal(t, production1.HatchingEggs, production2.HatchingEggs)

	require.WithinDuration(t, production1.CreatedAt, production2.CreatedAt, time.Second)
}
