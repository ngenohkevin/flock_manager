package db

import (
	"context"
	"database/sql"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomProduction(t *testing.T, breed Breed) Production {
	arg := CreateProductionParams{
		BreedID:      breed.BreedID,
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

	require.Equal(t, arg.BreedID, production.BreedID)
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

	production2, err := testQueries.GetProduction(context.Background(), production1.BreedID)
	require.NoError(t, err)
	require.NotEmpty(t, production2)

	require.Equal(t, production1.BreedID, production2.BreedID)
	require.Equal(t, production1.ID, production2.ID)
	require.Equal(t, production1.Eggs, production2.Eggs)
	require.Equal(t, production1.Dirty, production2.Dirty)
	require.Equal(t, production1.WrongShape, production2.WrongShape)
	require.Equal(t, production1.WeakShell, production2.WeakShell)
	require.Equal(t, production1.Damaged, production2.Damaged)
	require.Equal(t, production1.HatchingEggs, production2.HatchingEggs)

	require.WithinDuration(t, production1.CreatedAt, production2.CreatedAt, time.Second)
}

func TestUpdateProduction(t *testing.T) {
	breed := createdRandomBreed(t)

	production1 := createRandomProduction(t, breed)

	arg := UpdateProductionParams{
		ID:           production1.ID,
		Eggs:         util.RandomProduction(),
		Dirty:        util.RandomProduction(),
		WrongShape:   util.RandomProduction(),
		Damaged:      util.RandomProduction(),
		HatchingEggs: util.RandomProduction(),
	}

	production2, err := testQueries.UpdateProduction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, production2)

	require.Equal(t, production1.ID, production2.ID)
	require.Equal(t, arg.Eggs, production2.Eggs)
	require.Equal(t, arg.Dirty, production2.Dirty)
	require.Equal(t, arg.WrongShape, production2.WrongShape)
	require.Equal(t, arg.WeakShell, production2.WeakShell)
	require.Equal(t, arg.Damaged, production2.Damaged)
	require.Equal(t, arg.HatchingEggs, production2.HatchingEggs)

	require.WithinDuration(t, production1.CreatedAt, production2.CreatedAt, time.Second)

}

func TestListProduction(t *testing.T) {
	breed := createdRandomBreed(t)

	for i := 0; i < 10; i++ {
		createRandomProduction(t, breed)
	}

	arg := ListProductionParams{
		BreedID: breed.BreedID,
		Limit:   5,
		Offset:  5,
	}

	production, err := testQueries.ListProduction(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, production, 5)

	for _, prod := range production {
		require.NotEmpty(t, prod)
		require.Equal(t, arg.BreedID, prod.BreedID)
	}

}

func TestDeleteProduction(t *testing.T) {
	breed := createdRandomBreed(t)

	production1 := createRandomProduction(t, breed)
	err := testQueries.DeleteProduction(context.Background(), production1.BreedID)

	require.NoError(t, err)

	production2, err := testQueries.GetProduction(context.Background(), production1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, production2)
}
