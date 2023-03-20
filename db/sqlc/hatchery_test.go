package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomHatchery(t *testing.T, production Production) Hatchery {
	arg := CreateHatcheryParams{
		ProductionID: production.ID,
		Infertile:    util.RandomHatchery(),
		Early:        util.RandomHatchery(),
		Middle:       util.RandomHatchery(),
		Late:         util.RandomHatchery(),
		DeadChicks:   util.RandomHatchery(),
		AliveChicks:  util.RandomHatchery(),
	}
	hatchery, err := testQueries.CreateHatchery(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery)

	require.Equal(t, arg.ProductionID, hatchery.ProductionID)
	require.Equal(t, arg.Infertile, hatchery.Infertile)
	require.Equal(t, arg.Early, hatchery.Early)
	require.Equal(t, arg.Middle, hatchery.Middle)
	require.Equal(t, arg.Late, hatchery.Late)
	require.Equal(t, arg.DeadChicks, hatchery.DeadChicks)
	require.Equal(t, arg.AliveChicks, hatchery.AliveChicks)

	require.NotZero(t, hatchery.ID)
	require.NotZero(t, hatchery.CreatedAt)

	return hatchery
}

func TestCreateHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	createRandomHatchery(t, production)

}

func TestGetHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	hatchery1 := createRandomHatchery(t, production)

	hatchery2, err := testQueries.GetHatchery(context.Background(), hatchery1.ProductionID)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery2)

	require.Equal(t, hatchery1.ProductionID, hatchery2.ProductionID)
	require.Equal(t, hatchery1.Infertile, hatchery2.Infertile)
	require.Equal(t, hatchery1.Early, hatchery2.Early)
	require.Equal(t, hatchery1.Middle, hatchery2.Middle)
	require.Equal(t, hatchery1.Late, hatchery2.Late)
	require.Equal(t, hatchery1.DeadChicks, hatchery2.DeadChicks)
	require.Equal(t, hatchery1.AliveChicks, hatchery2.AliveChicks)

	require.WithinDuration(t, hatchery1.CreatedAt, hatchery2.CreatedAt, time.Second)
}

func TestUpdateHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	hatchery1 := createRandomHatchery(t, production)

	arg := UpdateHatcheryParams{
		ID:          hatchery1.ID,
		Infertile:   util.RandomHatchery(),
		Early:       util.RandomHatchery(),
		Middle:      util.RandomHatchery(),
		Late:        util.RandomHatchery(),
		DeadChicks:  util.RandomHatchery(),
		AliveChicks: util.RandomHatchery(),
	}

	hatchery2, err := testQueries.UpdateHatchery(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hatchery2)

	require.Equal(t, hatchery1.ID, hatchery2.ID)
	require.Equal(t, arg.Infertile, hatchery2.Infertile)
	require.Equal(t, arg.Early, hatchery2.Early)
	require.Equal(t, arg.Middle, hatchery2.Middle)
	require.Equal(t, arg.Late, hatchery2.Late)
	require.Equal(t, arg.DeadChicks, hatchery2.DeadChicks)
	require.Equal(t, arg.AliveChicks, hatchery2.AliveChicks)

	require.WithinDuration(t, hatchery1.CreatedAt, hatchery2.CreatedAt, time.Second)

}
func TestListHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	for i := 0; i < 10; i++ {
		createRandomHatchery(t, production)
	}
	arg := ListHatcheryParams{
		ProductionID: production.ID,
		Limit:        5,
		Offset:       5,
	}
	hatchery, err := testQueries.ListHatchery(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, hatchery, 5)

	for _, hatcheries := range hatchery {
		require.NotEmpty(t, hatcheries)
		require.Equal(t, arg.ProductionID, hatcheries.ProductionID)
	}
}
func TestDeleteHatchery(t *testing.T) {
	breed := createdRandomBreed(t)
	production := createRandomProduction(t, breed)

	hatchery1 := createRandomHatchery(t, production)
	err := testQueries.DeleteHatchery(context.Background(), hatchery1.ID)
	fmt.Println("Delete hatchery", err)
	require.NoError(t, err)

	hatchery2, err := testQueries.GetHatchery(context.Background(), hatchery1.ProductionID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, hatchery2)
	fmt.Println("get hatchery", err)
}
