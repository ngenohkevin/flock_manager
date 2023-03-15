package db

import (
	"context"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createdRandomBreed(t *testing.T) Breed {
	arg := util.RandomBreed()

	breed, err := testQueries.CreateBreed(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, breed)

	require.Equal(t, arg, breed.BreedName)

	require.NotZero(t, breed.BreedID)
	require.NotZero(t, breed.CreatedAt)

	return breed
}

func TestCreateAccount(t *testing.T) {
	createdRandomBreed(t)
}
