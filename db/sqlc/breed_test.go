package db

import (
	"context"
	"database/sql"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

func TestGetBreed(t *testing.T) {
	breed1 := createdRandomBreed(t)
	breed2, err := testQueries.GetBreed(context.Background(), breed1.BreedID)
	require.NoError(t, err)

	require.Equal(t, breed1.BreedID, breed2.BreedID)
	require.Equal(t, breed1.BreedName, breed2.BreedName)

	require.WithinDuration(t, breed1.CreatedAt, breed2.CreatedAt, time.Second)
}

func TestUpdateBreed(t *testing.T) {
	breed1 := createdRandomBreed(t)

	arg := UpdateBreedParams{
		BreedID:   breed1.BreedID,
		BreedName: util.RandomBreed(),
	}

	breed2, err := testQueries.UpdateBreed(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, breed2)

	require.Equal(t, breed1.BreedID, breed2.BreedID)
	require.Equal(t, arg.BreedID, breed2.BreedID)

	require.WithinDuration(t, breed1.CreatedAt, breed2.CreatedAt, time.Second)

}

func TestListBreeds(t *testing.T) {
	for i := 0; i < 10; i++ {
		createdRandomBreed(t)
	}
	arg := ListBreedsParams{
		Limit:  5,
		Offset: 5,
	}

	breeds, err := testQueries.ListBreeds(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, breeds, 5)

	for _, breed := range breeds {
		require.NotEmpty(t, breed)
	}
}

func TestDeleteBreed(t *testing.T) {
	breed1 := createdRandomBreed(t)
	err := testQueries.DeleteBreed(context.Background(), breed1.BreedID)
	require.NoError(t, err)

	breed2, err := testQueries.GetBreed(context.Background(), breed1.BreedID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, breed2)
}
