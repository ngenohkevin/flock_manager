package db

import (
	"context"
	"database/sql"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomPremises(t *testing.T, breed Breed) Premise {
	arg := CreatePremisesParams{
		BreedID: breed.BreedID,
		Farm:    util.RandomPremises(),
		House:   util.RandomPremises(),
	}
	premise, err := testQueries.CreatePremises(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, premise)

	require.Equal(t, arg.BreedID, premise.BreedID)
	require.Equal(t, arg.Farm, premise.Farm)
	require.Equal(t, arg.House, premise.House)

	require.NotZero(t, premise.ID)
	require.NotZero(t, premise.CreatedAt)

	return premise
}

func TestCreatePremise(t *testing.T) {
	breed := createdRandomBreed(t)
	createRandomPremises(t, breed)
}

func TestGetPremise(t *testing.T) {
	breed := createdRandomBreed(t)

	premise1 := createRandomPremises(t, breed)

	premise2, err := testQueries.GetPremises(context.Background(), premise1.BreedID)
	require.NoError(t, err)
	require.NotEmpty(t, premise2)

	require.Equal(t, premise1.BreedID, premise2.BreedID)
	require.Equal(t, premise1.Farm, premise2.Farm)
	require.Equal(t, premise1.House, premise2.House)

	require.WithinDuration(t, premise1.CreatedAt, premise2.CreatedAt, time.Second)
}

func TestUpdatePremise(t *testing.T) {
	breed := createdRandomBreed(t)

	premise1 := createRandomPremises(t, breed)

	arg := UpdatePremisesParams{
		ID:    premise1.ID,
		Farm:  util.RandomPremises(),
		House: util.RandomPremises(),
	}

	premise2, err := testQueries.UpdatePremises(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, premise2)

	require.Equal(t, arg.ID, premise2.ID)
	require.Equal(t, arg.Farm, premise2.Farm)
	require.Equal(t, arg.House, premise2.House)

	require.WithinDuration(t, premise1.CreatedAt, premise2.CreatedAt, time.Second)

}

func TestListPremise(t *testing.T) {
	breed := createdRandomBreed(t)

	for i := 0; i < 10; i++ {
		createRandomPremises(t, breed)
	}
	arg := ListPremisesParams{
		BreedID: breed.BreedID,
		Limit:   5,
		Offset:  5,
	}

	premise, err := testQueries.ListPremises(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, premise, 5)

	for _, premises := range premise {
		require.NotEmpty(t, premises)
		require.Equal(t, arg.BreedID, premises.BreedID)
	}

}

func TestDeletePremise(t *testing.T) {
	breed := createdRandomBreed(t)
	premise1 := createRandomPremises(t, breed)

	err := testQueries.DeletePremises(context.Background(), premise1.BreedID)

	require.NoError(t, err)

	premise2, err := testQueries.GetPremises(context.Background(), premise1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, premise2)
}