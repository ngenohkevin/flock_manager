package db

import (
	"context"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
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
